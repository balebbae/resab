package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/balebbae/resa-crud/models"
	"github.com/balebbae/resa-crud/models/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	
	server.GET("/available", getAvailables)
	server.GET("/available/:id", getAvailable)
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

func getAvailable(c *gin.Context) {
	availableId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse available id int"})
	}
	
	available, err := models.GetAvailableByID(availableId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch available by id"})
	}

	c.JSON(http.StatusOK, available)
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