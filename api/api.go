package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"openaigo/ai"
	"openaigo/auth"
)

type AIPrompt struct {
	Prompt string `json:"prompt" binding:"required"`
}

func Init(route *gin.Engine) {
	apiRoutes := route.Group("/api")
	{
		apiRoutes.Use(auth.AuthMiddleware())

		apiRoutes.POST("/generate/image", generateImageHandler)
	}
}

func generateImageHandler(c *gin.Context) {
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
}
