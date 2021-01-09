package controllers

import (
	"log"
	"net/http"
	"strconv"

	"members/db"
	"members/models"
	"members/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserController ...
type UserController struct{}

// Index action: GET /users
// @description ユーザー一覧取得API
// @Success 200 {object} repository.UserProfile
// @router /users [get]
func (pc UserController) Index(c *gin.Context) {
	var u repository.UserRepository
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
// @Success 200 {object} repository.UserProfile
// @router /users [post]
func (pc UserController) Create(c *gin.Context) {

	var requestBody repository.UserCreateRequest
	c.BindJSON(&requestBody)
	token := hashAndSalt(requestBody.Password)
	requestBody.Password = token

	var user models.User
	user.Email = requestBody.Email
	user.Password = token

	// c.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))

	// var u repository.UserRepository
	db := db.GetDB()
	if result := db.Create(&user); result.Error != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(201, gin.H{"message": "user_created"})
	}
}

// Show action: Get /users/:id
func (pc UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	user, err := u.GetByID(idInt)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, user)
	}
}

// Update action: Put /users/:id
func (pc UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
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
	var u repository.UserRepository
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
