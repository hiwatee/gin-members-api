package main

import (
	"members/db"
	"members/models"
	"members/server"
)

// @title メンバーズAPI
// @version 1.0
// @description メンバーズAPI
// @termsOfService https://github.com/hiwatee/gin-members-api

// @contact.name hiwatee
// @contact.url https://github.com/hiwatee
// @contact.email tochika.biz@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
func main() {
	db.Init()
	migration()
	server.Init()
}

func migration() {
	database := db.GetDB()
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Profile{})
}
