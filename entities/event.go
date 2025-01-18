package entities

import (
	"calendar/config"
	"fmt"
	"strings"
	"time"
)

type Event struct {
	EventId    string `json:"eventId"`
	UserId     int64  `json: "userId"`
	FirstName  string `json: "firstName"`
	LastName   string `json: "lastName"`
	UserName   string `json: "userName"`
	UserEmail  string `json: "userEmail"`
	UserTelNum string `json: "userTelNum"`
	Date       string `json: "date"`
	StartTime  string `json: "startTime"`
	EndTime    string `json: "endTime"`
	Text       string `json: "text"`
}

func (e *Event) GetDateUTCStartTime() string {
	date, err := time.Parse("02 January 2006", e.Date)
	if err != nil {
		fmt.Println("Date parsing error: ", err)
	}
	yy, mm, dd := date.Date()
	loc, _ := time.LoadLocation(config.GetFromEnv("LOCATION"))
	t, err := time.Parse("15:04", e.StartTime)
	if err != nil {
		fmt.Println("Time parsing error: ", err)
	}
	hh, MM, _ := t.Clock()
	newDate := time.Date(yy, mm, dd, hh, MM, 0, 0, loc)
	return newDate.Format(time.RFC3339)
}

func (e *Event) GetDateUTCEndTime() string {
	date, err := time.Parse("02 January 2006", e.Date)
	if err != nil {
		fmt.Println("Date parsing error: ", err)
	}
	yy, mm, dd := date.Date()
	loc, _ := time.LoadLocation(config.GetFromEnv("LOCATION"))
	t, err := time.Parse("15:04", e.EndTime)
	if err != nil {
		fmt.Println("Time parsing error: ", err)
	}
	hh, MM, _ := t.Clock()
	newDate := time.Date(yy, mm, dd, hh, MM, 0, 0, loc)
	return newDate.Format(time.RFC3339)
}
func (e *Event) GetDescription() string {
	var sb strings.Builder
	sb.WriteString(e.FirstName + " ")
	sb.WriteString(e.LastName + "\n")
	sb.WriteString("Telegram Username: " + e.UserName + "\n")
	sb.WriteString("E-Mail: " + e.UserEmail + "\n")
	sb.WriteString("Tel. number: " + e.UserTelNum + "\n")
	sb.WriteString(e.Text + "\n")
	return sb.String()
}
