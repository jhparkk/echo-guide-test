package user

type userResponse struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
}
