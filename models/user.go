package models

import "time"

// User is user models property
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"size(128)" json:"email"`
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
