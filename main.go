package main

import (
	"calendar/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.GET("/availableTimes", handlers.GetAvailableTimes)
	v1.POST("/event", handlers.CreateEvent)
	v1.DELETE("/event", handlers.DeleteEvent)
	v1.GET("/event", handlers.GetMyEvents)
	router.Run("localhost:8080")
}
