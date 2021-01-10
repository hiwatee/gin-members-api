package controllers

import (
	"members/db"
	"members/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegistrationController ....
type RegistrationController struct{}

// Create action: Post /registration
// @description ユーザー新規登録API
// @Success 200 {object} UserResponse
// @Param   body        body    UserCreateRequest   true        "User Create Request"
// @router /registration [post]
func (pc RegistrationController) Create(c *gin.Context) {

	var requestBody UserCreateRequest
	c.BindJSON(&requestBody)
	token := hashAndSalt(requestBody.Password)
	requestBody.Password = token

	var user models.User
	user.Email = requestBody.Email
	user.Password = token

	db := db.GetDB()
	if result := db.Create(&user); result.Error != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(201, UserResponse(user))
	}
}
