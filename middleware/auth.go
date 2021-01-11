package middleware

import (
	"log"
	"members/db"
	"members/models"

	"github.com/gin-gonic/gin"
)

// UserAuth ...
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("access_token")

		db := db.GetDB()
		var t models.AccessToken

		if err := db.Where("token = ?", token).First(&t).Error; err != nil {
			log.Print(err.Error())
			c.AbortWithStatusJSON(401, gin.H{"error": "login_failure"})
			return
		}

		c.Next()
		// log.Println("after logic")
	}
}
