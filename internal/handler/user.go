package handler

import (
	v1 "cheemshappy_pay/api/v1"
	"cheemshappy_pay/internal/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req v1.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	// 获取客户端IP
	clientIP := getClientIP(ctx.Request)

	token, err := h.userService.Login(ctx, &req, clientIP, "admin")
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrUnauthorizedAP, nil)
		return
	}
	v1.HandleSuccess(ctx, v1.LoginResponseData{
		AccessToken: token,
	})
}

func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, user)
}

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req v1.UpdateProfileRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	if req.Email == "" {
		v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("请输入邮箱"), nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func (h *UserHandler) IsLogin(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleSuccess(ctx, false)
		return
	}
	v1.HandleSuccess(ctx, true)
}

func (h *UserHandler) UpdateUserInfo(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	var req v1.RegisterRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err = h.userService.UpdateUserInfo(ctx, userId, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func getClientIP(r *http.Request) string {
	// Try Forwarded header (RFC 7239 standard)
	if forwarded := r.Header.Get("Forwarded"); forwarded != "" {
		if parts := strings.Split(forwarded, ";"); len(parts) > 0 {
			for _, part := range parts {
				part = strings.TrimSpace(parts[0])
				if strings.HasPrefix(part, "for=") {
					ip := strings.TrimPrefix(part, "for=")
					ip = strings.Trim(ip, `"`) // Remove quotes if present
					return strings.Split(ip, ",")[0]
				}
			}
		}
	}

	// Fallback to X-Forwarded-For (common de facto standard)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.Split(xff, ",")[0]
	}

	// Final fallback to remote address
	return strings.Split(r.RemoteAddr, ":")[0] // Remove port number
}
