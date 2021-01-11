package models

import (
	"members/db"
	"members/service"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User is user models property
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;size(128)" json:"email"`
	Password  string    `gorm:"size(128)" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
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
	if err := db.Table("users").Select("*").Scan(&u).Error; err != nil {
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

// CheckPassword ...
func (user User) CheckPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

// CreateToken ...
func (user User) CreateToken() string {
	// db := db.GetDB()
	token := service.HashAndSalt(string(rune(user.ID)))
	// expired_at := time.Now().AddDate(0, 0, 30)

	// t := new(Token)
	// if err := o.QueryTable("token").Filter("User", m).One(t); err != orm.ErrNoRows {
	// 	t.Token = token
	// 	t.ExpiredAt = expired_at
	// 	o.Update(t)
	// } else {
	// 	t.Token = token
	// 	t.User = m
	// 	t.ExpiredAt = expired_at
	// 	o.Insert(t)
	// }
	return token
}
