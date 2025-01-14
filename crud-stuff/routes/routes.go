package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	// Middleware

	// Users
	server.POST("/signup", signup)
	server.POST("/login", login)

		// Employers
		// Employees


	// Availables
	server.GET("/available", getAvailables)
	server.GET("/available/:id", getAvailable)
	server.POST("/available", createAvailable)
	server.PUT("/available/:id", updateAvailable)
	server.DELETE("/available/:id", deleteAvailable)
	
	// Schedules

	// Weekly Shift


	// Restaurants

}
