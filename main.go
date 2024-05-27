package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"openaigo/pkg/ai"
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

var router = gin.Default()

type AIPrompt struct {
	Prompt string `json:"prompt" binding:"required"`
}

func main() {
	api := router.Group("/api")

	api.Use(AuthMiddleware())
	{
		api.POST("/generate/image/", func(c *gin.Context) {
			var json AIPrompt

			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			url, err := ai.GenerateImageFromPrompt(json.Prompt)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
				return
			}
			c.String(http.StatusOK, url)
		})
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
