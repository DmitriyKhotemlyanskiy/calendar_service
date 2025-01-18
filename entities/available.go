package entities

import (
	"calendar/config"
	"strconv"
	"time"

	"google.golang.org/api/calendar/v3"
)

type AvailableTime struct {
	Days []Day `json:"days"`
}

func getInt(str string) int64 {
	num, _ := strconv.Atoi(str)
	return int64(num)
}

func roundToNextDay(timeNow time.Time) time.Time {
	yyyy, mm, dd := timeNow.Date()
	newDay := time.Date(yyyy, mm, dd+1, 0, 0, 0, 1, timeNow.Location())
	return newDay
}

func initAvailableTime() *AvailableTime {
	DaysInterval := getInt(config.GetFromEnv("DAYS_INTERVAL"))
	var days []Day
	dateTimeNow := time.Now()
	startWork, endWork := config.StartStop(dateTimeNow)
	for i := 0; i < int(DaysInterval); i++ {
		day := InitDay(dateTimeNow)
		days = append(days, *day)
		if dateTimeNow.After(startWork) && dateTimeNow.Before(endWork) {
			dateTimeNow = roundToNextDay(dateTimeNow)
		} else {
			dateTimeNow = dateTimeNow.AddDate(0, 0, 1)
		}
	}
	return &AvailableTime{
		Days: days,
	}
}

func FindAvailableTimes(events []*calendar.Event) *AvailableTime {
	availableTime := initAvailableTime()
	for _, day := range availableTime.Days {
		day.FindAvailableTimeInDay(events)
	}
	return availableTime
}
