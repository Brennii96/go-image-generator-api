package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"openaigo/api"
)

func main() {
	router := gin.Default()
	api.Init(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
