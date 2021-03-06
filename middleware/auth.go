package middleware

import (
	"members/db"
	"members/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// UserAuth ...
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := func() string {
			// NOTE: release以外ではTokenをCookieからも取得 | swagger用の対応です
			if os.Getenv("GIN_MODE") == "debug" {
				t, _ := c.Cookie("access_token")
				return t
			}
			return ""
		}

		db := db.GetDB()
		var t models.AccessToken

		if err := db.Where("token = ? AND expired_at > ?", token(), time.Now()).First(&t).Error; err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "login_failure"})
			return
		}

		c.Next()
	}
}
