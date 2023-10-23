package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users, err := service.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H("Erro":"There was an error trying to get users"))
		return
	}

	c.JSON(http.StatusOK, users)
}
