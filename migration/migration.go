package migration

import (
	"members/db"
	"members/models"
)

// AutoMigration ...
func AutoMigration() {
	database := db.GetDB()
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Profile{})
	database.AutoMigrate(&models.Token{})
	database.AutoMigrate(&models.AccessToken{})
}
