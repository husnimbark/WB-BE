package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"name" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
}
type UserResponseDelete struct {
	ID int `json:"id"`
}
