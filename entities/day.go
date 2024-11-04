package entities

import (
	"calendar/config"
	"strconv"
	"time"
)

type Day struct {
	Date     string
	TimesArr []string
}

func InitDay(date time.Time) *Day {
	var day Day
	day.Date = date.Format("02 January 2006")
	num, _ := strconv.Atoi(config.GetFromEnv("MEETING_DURATION"))
	duration := time.Duration(int64(num)) * time.Minute
	startWork, _ := time.Parse("15:04", config.GetFromEnv("WORK_TIME_FROM"))
	endWork, _ := time.Parse("15:04", config.GetFromEnv("WORK_TIME_TO"))
	for startWork.Compare(endWork) != 0 {
		day.TimesArr = append(day.TimesArr, startWork.Format("15:04"))
		startWork = startWork.Add(duration)
	}
	return &day
}
