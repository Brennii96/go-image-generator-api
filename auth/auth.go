package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func loadAuthToken() string {
	token := os.Getenv("AUTH_TOKEN")
	if token == "" {
		panic("Authentication token not found. Please set the AUTH_TOKEN environment variable.")
	}
	return token
}

func validateToken(token string) bool {
	authToken := loadAuthToken()
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
