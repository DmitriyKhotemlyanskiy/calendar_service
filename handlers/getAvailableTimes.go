package handlers

import (
	"calendar/config"
	"calendar/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAvailableTimes(c *gin.Context) {
	client := config.Auth()
	srv, err := config.GetService(client)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Can't connect to Google Calendar.\nPlease, try later")
	}
	items, err := entities.Calendar{}.GetUpcomingEvents(srv)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, err)
		return
	}
	availableTimes := entities.FindAvailableTimes(items)
	c.JSON(http.StatusOK, availableTimes)
}
