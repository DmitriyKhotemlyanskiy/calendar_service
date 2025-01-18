package handlers

import (
	"calendar/config"
	"calendar/entities"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/calendar/v3"
)

func CreateEvent(c *gin.Context) {
	conn := config.Connect()
	db := conn.GetDB()
	defer db.Close()
	client := config.Auth()
	srv, err := config.GetService(client)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Can't connect to Google Calendar.\nPlease, try later")
	}
	var newEvent entities.Event
	c.BindJSON(&newEvent)
	event := &calendar.Event{
		Summary:     newEvent.FirstName,
		Location:    "Tattoo studio ProTattoo",
		Description: newEvent.GetDescription(),
		Start: &calendar.EventDateTime{
			DateTime: newEvent.GetDateUTCStartTime(),
			TimeZone: config.GetFromEnv("LOCATION"),
		},
		End: &calendar.EventDateTime{
			DateTime: newEvent.GetDateUTCEndTime(),
			TimeZone: config.GetFromEnv("LOCATION"),
		},
	}

	calendarId := config.GetFromEnv("USER_EMAIL")
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
		c.JSON(http.StatusBadRequest, err)
	}
	newEvent.EventId = event.Id
	row := db.QueryRow(`INSERT INTO events (event_id, user_id, date, start_time, end_time, text) 
						VALUES ($1, $2, $3, $4, $5, $6)`,
		newEvent.EventId, int(newEvent.UserId), newEvent.Date, newEvent.StartTime, newEvent.EndTime, newEvent.Text)
	log.Println(row)
	row = db.QueryRow(`INSERT INTO users (id, first_name, last_name, user_name, user_email, user_tel_num)
	SELECT $1, $2, $3, $4, $5, $6
	WHERE NOT EXISTS(
	SELECT id FROM users WHERE id = $1)`, newEvent.UserId, newEvent.FirstName, newEvent.LastName, newEvent.UserName, newEvent.UserEmail, newEvent.UserTelNum)
	log.Println(row)
	c.JSON(http.StatusOK, event.Id)
}

func DeleteEvent(c *gin.Context) {
	client := config.Auth()
	conn := config.Connect()
	db := conn.GetDB()
	defer db.Close()
	srv, err := config.GetService(client)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Can't connect to Google Calendar.\nPlease, try later")
	}
	var newEvent entities.Event
	c.BindJSON(&newEvent)
	fmt.Println("Deleted event ID:   ", newEvent.EventId)
	calendarId := config.GetFromEnv("USER_EMAIL")
	resp := srv.Events.Delete(calendarId, newEvent.EventId).Do()
	row := db.QueryRow(`DELETE FROM events WHERE event_id = $1`, newEvent.EventId)
	log.Println(row)
	if resp != nil {
		c.JSON(http.StatusNotFound, "Event not found. Reason: already deleted")
		return
	}
	c.JSON(http.StatusOK, "Event deleted successfully")
}

func GetMyEvents(c *gin.Context) {
	conn := config.Connect()
	db := conn.GetDB()
	defer db.Close()
	var newEvent entities.Event
	c.BindJSON(&newEvent)
	var MyEvents []entities.Event

	rows, err := db.Query(`SELECT b.event_id, b.user_id, a.first_name, a.last_name, a.user_name, a.user_email, a.user_tel_num, b.date, b.start_time, b.end_time, b.text
							FROM users a
							LEFT JOIN events b ON a.id = b.user_id
							WHERE a.id = $1`, int(newEvent.UserId))
	if err != nil {
		log.Println("No data in DB for this query\n", err)
	}
	for rows.Next() {
		var event entities.Event
		err = rows.Scan(&event.EventId,
			&event.UserId,
			&event.FirstName,
			&event.LastName,
			&event.UserName,
			&event.UserEmail,
			&event.UserTelNum,
			&event.Date,
			&event.StartTime,
			&event.EndTime,
			&event.Text)
		if err != nil {
			log.Println("Row not readable\n", err)
		}
		MyEvents = append(MyEvents, event)
	}
	c.JSON(http.StatusOK, MyEvents)
}
