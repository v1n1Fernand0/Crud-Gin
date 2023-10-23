package main

import (
	"crud-gin/api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupUserRoutes(router)

	port := 8080
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
