package middleware

import (
	"github.com/gin-gonic/gin"
)

// UserAuth ...
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(401, gin.H{"error": "login_failure"})
		return
		c.Next()
		// log.Println("after logic")
	}
}
