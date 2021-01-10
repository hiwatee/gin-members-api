package models

import "time"

// Token ...
type Token struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `json:"token"`
	User      User      `json:"-" binding:"required"`
	UserID    uint      `gorm:"not null"  json:"userId"`
	CreatedAt time.Time `gorm:"autoCreateTime"  json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
	ExpiredAt time.Time `json:"expiredAt"`
}

// AccessToken ...
type AccessToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `json:"token"`
	User      User      `json:"-" binding:"required"`
	UserID    uint      `gorm:"not null"  json:"userId"`
	CreatedAt time.Time `gorm:"autoCreateTime"  json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
	ExpiredAt time.Time `json:"expiredAt"`
}
