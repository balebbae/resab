package main

import (
	"fmt"
	"net/http"

	"github.com/balebbae/resa-crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	
	server.GET("/available", getAvailables)
	server.POST("/available", createAvailable)

	server.Run(":8080") // localhost:8080
}

func getAvailables(c *gin.Context) {
	availabilties := models.GetAllAvailables()
	c.JSON(http.StatusOK, availabilties)
}

func createAvailable(c *gin.Context) {
	var available models.Available
	err := c.ShouldBindJSON(&available) // Binding request body the model of the availabiltiy

	if err != nil {
		fmt.Println(available)

		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse bad request"})
		return
	}

	available.UserID = 1
	available.ID = 1
	available.Save()
	fmt.Println(available)
	c.JSON(http.StatusCreated, gin.H{"message": "Available created", "available": available})

}