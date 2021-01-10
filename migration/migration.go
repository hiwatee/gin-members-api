package migration

import (
	"members/db"
	"members/models"
)

func AutoMigration() {
	database := db.GetDB()
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Profile{})
}
