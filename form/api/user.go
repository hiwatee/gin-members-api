package api

type User struct {
	ID       uint   `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"pasword" binding:"required"`
}
