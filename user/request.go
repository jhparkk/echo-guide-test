package user

type userRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
	Token    string `json:"token"`
}
