package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}
