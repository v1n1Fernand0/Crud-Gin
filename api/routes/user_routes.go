// api/routes/user_routes.go
package routes

import (
	"crud-gin/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		// Definir outras rotas relacionadas ao usuário...
	}

	// Adicione uma rota básica para a raiz
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the CRUD API"})
	})
}
