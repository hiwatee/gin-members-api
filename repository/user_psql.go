package repository

import (
	"members/db"
	"members/models"

	"github.com/gin-gonic/gin"
)

// UserRepository ...
type UserRepository struct{}

// User is alias of entity.User struct
type User models.User

// UserProfile ...
type UserProfile struct {
	ID    uint   `json:"id"`
	Email string `gorm:"size(128)" json:"email"`
}

type UserCreateRequest struct {
	Email    string `gorm:"size(128)" json:"email"`
	Password string `json:"password"`
}

// GetAll is get all User
func (user UserRepository) GetAll() ([]UserProfile, error) {
	db := db.GetDB()
	var u []UserProfile
	if err := db.Table("users").Select("id, email").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// CreateModel is create User model
func (user UserRepository) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// GetByID is get a User by ID
func (user UserRepository) GetByID(id int) (models.User, error) {
	db := db.GetDB()
	var u models.User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// UpdateByID is update a User
func (user UserRepository) UpdateByID(id int, c *gin.Context) (models.User, error) {
	db := db.GetDB()
	var u models.User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	u.ID = uint(id)
	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a User by ID
func (user UserRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
