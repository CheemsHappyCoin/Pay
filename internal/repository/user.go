package repository

import (
	v1 "cheemshappy_pay/api/v1"
	"cheemshappy_pay/internal/model"
	"context"
	"errors"
	"time"

	"github.com/imroc/req/v3"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByOpenID(ctx context.Context, openID string) (*model.User, error)
	GetByIDs(ctx context.Context, id []string) ([]*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	SetUserToken(ctx context.Context, key string, token string) error
	GetUserToken(ctx context.Context, key string) (string, error)
	UpdateUserInfo(ctx context.Context, user *model.User) error
}

func NewUserRepository(r *Repository, client *req.Client) UserRepository {
	return &userRepository{
		Repository: r,
		client:     client,
	}
}

type userRepository struct {
	*Repository
	client *req.Client
}

func (r *userRepository) SetUserToken(ctx context.Context, key string, token string) error {
	nx := r.rdb.SetNX(ctx, "access_token_wx_user_"+key, token, time.Second*180)
	if nx.Err() != nil {
		return nx.Err()
	}
	return nil
}

func (r *userRepository) GetUserToken(ctx context.Context, key string) (string, error) {
	get := r.rdb.Get(ctx, "access_token_wx_user_"+key)
	if get.Err() != nil {
		return "", get.Err()
	}
	return get.Val(), nil
}

func (r *userRepository) GetByIDs(ctx context.Context, ids []string) ([]*model.User, error) {
	tx := newUser(r.DB(ctx))
	find, err := tx.Where(tx.UserID.In(ids...)).Find()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return find, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByOpenID(ctx context.Context, openid string) (*model.User, error) {
	var use model.User
	if err := r.DB(ctx).Where("openid = ?", openid).Find(&use).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &use, nil
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return &use, nil
	}
	return &use, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUserInfo(ctx context.Context, user *model.User) error {
	tx := newUser(r.DB(ctx))
	_, err := tx.Where(tx.UserID.Eq(user.UserID)).Updates(user)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
