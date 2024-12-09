package ssomodels

type RegisterUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,password_complexity"`
	Name     string `json:"name,omitempty"`
}

type RegisterResp struct {
	Success bool `json:"success"`
}

type LogIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LogInResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LogInViaTg struct {
	Email string `json:"email" validate:"required,email"`
}

type LogInViaTgResp struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
}

type CheckOTPAndLogIn struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type CheckOTPAndLogInResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UpdateUserInfo struct {
	ID      string `json:"id" validate:"required"`
	Email   string `json:"email,omitempty"`
	Name    string `json:"name,omitempty"`
	TgLink  string `json:"tg_link,omitempty"`
	IsAdmin bool   `json:"is_admin,omitempty"`
}

type UpdateUserInfoResp struct {
	Success bool `json:"success"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenResp struct {
	AccessToken string `json:"access_token"`
}

type IsAdmin struct {
	UserID string `json:"user_id" validate:"required"`
}

type IsAdminResp struct {
	IsAdmin bool `json:"is_admin"`
}

type AuthCheck struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type AuthCheckResp struct {
	IsValid bool   `json:"is_valid"`
	UserID  string `json:"user_id"`
}
