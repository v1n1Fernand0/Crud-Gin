package handlers

import (
	"crud-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := service.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error trying to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
