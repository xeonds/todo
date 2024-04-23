package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != token {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Next()
	}
}
