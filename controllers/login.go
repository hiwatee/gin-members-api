package controllers

import (
	"log"
	"members/db"
	"members/models"
	"time"

	"github.com/gin-gonic/gin"
)

// LoginController ....
type LoginController struct{}

// LoginRequest ...
type LoginRequest struct {
	Email    string `gorm:"unique;size(128)" json:"email" example:"info@example.com"`
	Password string `gorm:"size(128)" json:"password" example:"password2021"`
}

// LoginSuccessResponse ...
type LoginSuccessResponse struct {
	Message     string    `json:"message" example:"login_success"`
	AccessToken string    `json:"accessToken" example:"$$t$hisistoken"`
	ExpiredAt   time.Time `json:"expiedAt" example:"2021-01-10T15:42:36+09:00"`
}

// Create action: Post /registration
// @description ユーザー新規登録API
// @Success 200 {object} LoginSuccessResponse
// @Failure 403 {object} DefaultErrorResponse
// @Param   body        body    LoginRequest   true        "User Create Request"
// @router /login [post]
func (pc LoginController) Create(c *gin.Context) {
	db := db.GetDB()

	var requestBody LoginRequest
	c.BindJSON(&requestBody)

	email := requestBody.Email
	password := requestBody.Password

	var u models.User

	if err := db.Where("email= ?", email).First(&u).Error; err != nil {
		c.JSON(401, DefaultErrorResponse{Message: "login_failure"})
		return
	}

	if !u.CheckPassword(password) {
		c.JSON(401, DefaultErrorResponse{Message: "login_failure"})
		return
	}

	token := u.CreateToken()
	log.Print(token)

	c.JSON(201, gin.H{"message": "login_success", "token": "this is acccess_token"})
	return
}
