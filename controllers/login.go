package controllers

import (
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
// @Param   body        body    LoginRequest   true        "User Create Request"
// @router /login [post]
func (pc LoginController) Create(c *gin.Context) {

	var requestBody LoginRequest
	c.BindJSON(&requestBody)

	// requestBody.Password
	// requestBody.Email

	c.JSON(201, gin.H{"message": "login_success", "token": "this is acccess_token"})
}
