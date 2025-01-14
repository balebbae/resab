package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/balebbae/resa-crud/models"
	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse available id to int"})
	}
	
	available, err := models.GetAvailableByID(availableId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch available by id"})
	}

	c.JSON(http.StatusOK, available)
}

func createAvailable(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "could not authorize user"})
		return
	}
	

	var available models.Available
	err := c.ShouldBindJSON(&available) // Binding request body the model of the availabiltiy

	if err != nil {
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

func updateAvailable(c *gin.Context) {

	// parse the userID from the context
	availableId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse available id to int"})
		return
	}

	// Grab the userID passed by the context from GIN
	// userID := c.GetInt64("userID") // Used later when implemenet authorization 

	// fetch available from DB
	_, err = models.GetAvailableByID(availableId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch available"})
		return
	}

	var updatedAvailable models.Available
	err = c.ShouldBindJSON(&updatedAvailable)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	updatedAvailable.ID = availableId // Don't change the ID of the available
	err = updatedAvailable.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not update the available"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "available updated successfully"})
}

func deleteAvailable(c *gin.Context) {
	availableId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse available id to int"})
		return
	}

	available, err := models.GetAvailableByID(availableId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch available"})
		return
	}

	err = available.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not delete the available"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "available successfully deleted"})
}