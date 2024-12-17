package main

import (
	"fmt"
	"net/http"

	"github.com/balebbae/resa-crud/models"
	"github.com/balebbae/resa-crud/models/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	
	server.GET("/available", getAvailables)
	server.POST("/available", createAvailable)

	server.Run(":8080") // localhost:8080
}

func getAvailables(c *gin.Context) {
	availabilties, err := models.GetAllAvailables()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch availables"})
	}

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

	err = available.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create available"})
	}

	fmt.Println(available)
	c.JSON(http.StatusCreated, gin.H{"message": "Available created", "available": available})

}