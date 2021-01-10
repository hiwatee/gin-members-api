package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"members/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserController ...
type UserController struct{}

// UserResponse ...
type UserResponse struct {
	ID        uint      `gorm:"primaryKey" json:"id" example:"1"`
	Email     string    `gorm:"unique;size(128)" json:"email" example:"info@example.com"`
	Password  string    `gorm:"size(128)" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt" example:"2021-01-10T15:42:36+09:00"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt" example:"2021-01-10T15:42:36+09:00"`
}

// UserCreateRequest ...
type UserCreateRequest struct {
	Email    string `gorm:"unique;size(128)" json:"email" example:"info@example.com"`
	Password string `gorm:"size(128)" json:"password" example:"password2021"`
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
		users := make([]UserResponse, len(p))
		for n := range p {
			users[n] = UserResponse(p[n])
		}
		c.JSON(200, users)
	}
}

// Show action: Get /users/:id
// @description ユーザー詳細情報API
// @Success 200 {object} UserResponse
// @Param	id		path 	string	true		"The key for user"
// @router /users/{id} [get]
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
