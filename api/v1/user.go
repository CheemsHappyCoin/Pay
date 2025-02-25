package v1

type RegisterRequest struct {
	Email      string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Nickname   string `json:"nickname" example:"alan"`
	Password   string `json:"password" binding:"required" example:"123456"`
	RePassword string `json:"re_password"  example:"123456"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Nickname     string `json:"nickname" example:"alan"`
	Email        string `json:"email"  example:"1234@gmail.com"`
	Phone        string `json:"phone"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	CashPic      string `json:"cash_pic"`
}
type GetProfileResponseData struct {
	UserId       string `json:"userId"`
	Email        string `json:"email" example:"1234@gmail.com"`
	Nickname     string `json:"nickname" example:"alan"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}

type HotOwner struct {
	Hot          int32  `json:"hot"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	UID          string `json:"uid"`
	Email        string `json:"email"`
}
