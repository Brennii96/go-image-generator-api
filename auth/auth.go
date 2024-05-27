package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authToken = "test"

func validateToken(token string) bool {
	return token == fmt.Sprintf("Bearer %s", authToken)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if !validateToken(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		c.Next()
	}
}
