package ai

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"github.com/sashabaranov/go-openai"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GenerateImageFromPrompt(prompt string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()
	reqUrl := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	resp, err := client.CreateImage(ctx, reqUrl)
	if err != nil {
		return "", fmt.Errorf("error creating image, %v", err)
	}

	if len(resp.Data) > 0 {
		return resp.Data[0].URL, nil
	}

	return "", fmt.Errorf("no image generated")
}
