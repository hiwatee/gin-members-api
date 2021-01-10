package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"members/db"
	"members/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserController ...
type UserController struct{}

// UserResponse ...
type UserResponse struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;size(128)" json:"email"`
	Password  string    `gorm:"size(128)" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

// UserCreateRequest ...
type UserCreateRequest struct {
	Email    string `gorm:"unique;size(128)" json:"email"`
	Password string `gorm:"size(128)" json:"password"`
}

// Index action: GET /users
// @description ユーザー一覧取得API
// @Success 200 {object} UserResponse
// @router /users [get]
func (pc UserController) Index(c *gin.Context) {
	var u models.User
	p, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Create action: Post /users
// @description ユーザー一覧取得API
// @Success 200 {object} UserResponse
// @router /users [post]
func (pc UserController) Create(c *gin.Context) {

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

// Show action: Get /users/:id
// @description ユーザー詳細情報API
// @Success 200 {object} UserResponse
// @router /users/:id [get]
func (pc UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var u models.User
	idInt, _ := strconv.Atoi(id)
	user, err := u.GetByID(idInt)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, UserResponse(user))
	}
}

// Update action: Put /users/:id
func (pc UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var u models.User
	idInt, _ := strconv.Atoi(id)
	p, err := u.UpdateByID(idInt, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (pc UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var u models.User
	idInt, _ := strconv.Atoi(id)
	if err := u.DeleteByID(idInt); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "ID" + id + "のユーザーを削除しました"})
	return
}

func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
