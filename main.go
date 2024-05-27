package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"openaigo/pkg/ai"
)

var router = gin.Default()

type AIPrompt struct {
	Prompt string `json:"prompt" binding:"required"`
}

func main() {
	router.POST("/generate/image/", func(c *gin.Context) {
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

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
