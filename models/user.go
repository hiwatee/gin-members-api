package models

import (
	"members/db"
	"time"

	"github.com/gin-gonic/gin"
)

// User is user models property
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;size(128)" json:"email"`
	Password  string    `gorm:"size(128)" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Profile ...
type Profile struct {
	ID        uint      `json:"id" binding:"required"`
	Name      string    `json:"content" binding:"required"`
	User      User      `json:"-" binding:"required"`
	UserID    uint      `gorm:"not null"  json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// GetAll is get all User
func (user User) GetAll() ([]User, error) {
	db := db.GetDB()
	var u []User
	if err := db.Table("users").Select("id, email").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// CreateModel is create User model
func (user User) CreateModel(c *gin.Context) (User, error) {
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
func (user User) GetByID(id int) (User, error) {
	db := db.GetDB()
	var u User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// UpdateByID is update a User
func (user User) UpdateByID(id int, c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User
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
func (user User) DeleteByID(id int) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
