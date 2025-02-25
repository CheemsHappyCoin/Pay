package service

import (
	v1 "cheemshappy_pay/api/v1"
	"cheemshappy_pay/internal/model"
	"cheemshappy_pay/internal/repository"
	"cheemshappy_pay/pkg/helper/regexps"
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/imroc/req/v3"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest, clientIP string, userType string) (string, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
	GetUserInfo(ctx context.Context, userId string) (*model.User, error)
	GetUserToken(ctx context.Context, key string) (string, error)
	UpdateUserInfo(ctx context.Context, userId string, req *v1.RegisterRequest) error
}

func NewUserService(service *Service, userRepo repository.UserRepository, client *req.Client) UserService {
	return &userService{
		userRepo: userRepo,
		client:   client,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	client   *req.Client
	*Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check username
	if user, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil && user != nil {
		return v1.ErrEmailAlreadyUse
	}

	if req.Password != req.RePassword {
		return v1.ErrRePasswd
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}
	user := &model.User{
		Nickname: req.Nickname,
		UserID:   userId,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		// TODO: other repo
		return nil
	})
	return err
}

func (s *userService) WxRegister(ctx context.Context, openid string) error {
	// check username
	if user, err := s.userRepo.GetByOpenID(ctx, openid); err == nil && user != nil && user.ID > 0 {
		return v1.ErrEmailAlreadyUse
	}

	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	user := &model.User{
		Nickname: "用户名" + fmt.Sprintf("%v", rand.Intn(10000)),
		UserID:   userId,
		Email:    "",
		Password: "",
	}
	// Create a user
	if err = s.userRepo.Create(ctx, user); err != nil {
		return err
	}
	return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest, clientIP string, userType string) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(user.UserID, userType, clientIP, time.Now().Add(2*time.Hour))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		Email:        user.Email,
		UserId:       user.UserID,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Introduction: user.Introduction,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	if req.Email != "" {
		euser, err := s.userRepo.GetByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if euser.ID != user.ID {
			return fmt.Errorf("重复的邮箱地址")
		}
		user.Email = req.Email
	}
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Introduction != "" {
		user.Introduction = req.Introduction
	}
	if req.Phone != "" && regexps.ValidatePhoneNumber(req.Phone) {
		user.Phone = req.Phone
	}
	if req.Avatar != "" && regexps.ValidURL(req.Avatar) {
		user.Avatar = req.Avatar
	}

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetUserInfo(ctx context.Context, userId string) (*model.User, error) {
	us, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	us.Password = "***"
	return us, nil
}

func (s *userService) GetUserToken(ctx context.Context, key string) (string, error) {
	us, err := s.userRepo.GetUserToken(ctx, key)
	if err != nil {
		return "", err
	}
	return us, nil
}

func (s *userService) UpdateUserInfo(ctx context.Context, userId string, req *v1.RegisterRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}
	user.Email = req.Email
	if req.Password == "" {
		return v1.ErrRePasswd
	}
	if req.Email == "" {
		return v1.ErrRePasswd
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.RePassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.UpdateUserInfo(ctx, user)
}
