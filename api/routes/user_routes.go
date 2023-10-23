package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/v1n1Fernand0/crud-gin/api/handlers"
)

func SetupUserRoutes(routes *gin.Engine){
	userGroup := routes.Group("/users")
	{
		userGroup.Get("/", handlers.GetUsersimport)
	}
}