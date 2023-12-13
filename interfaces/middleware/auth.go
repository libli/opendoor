package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth 鉴权中间件
func Auth(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.GetHeader("Authorization")
		expectedToken := "Bearer " + token
		if userToken != expectedToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
