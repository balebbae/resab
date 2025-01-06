package routes

import (
	"net/http"

	"github.com/balebbae/resa-crud/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user) // Binding request body the model of the availabiltiy

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse bad request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}