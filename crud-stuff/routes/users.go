package routes

import (
	"net/http"

	"github.com/balebbae/resa-crud/models"
	"github.com/balebbae/resa-crud/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user) // Binding request body the model of the availabiltiy

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse bad request"})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse bad request"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}