package dto

type UserLoginDTO struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"password"`
}

type UserRegisterDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
