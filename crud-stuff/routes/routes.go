package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/available", getAvailables)
	server.GET("/available/:id", getAvailable)
	server.POST("/available", createAvailable)
}