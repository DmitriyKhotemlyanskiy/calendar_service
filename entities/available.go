package entities

import (
	"calendar/config"
	"fmt"
	"strconv"
	"time"
)

type AvailableTime struct {
	Days []Day
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

func InitAvailableTime() *AvailableTime {
	DaysInterval := getInt(config.GetFromEnv("DAYS_INTERVAL"))

	var days []Day
	dateTimeNow := time.Now()
	startWork, endWork := config.StartStop(dateTimeNow)
	for i := 0; i < int(DaysInterval); i++ {
		day := InitDay(dateTimeNow)
		days = append(days, *day)
		if dateTimeNow.After(startWork) && dateTimeNow.Before(endWork) {
			dateTimeNow = roundToNextDay(dateTimeNow)
			fmt.Println("Round day", dateTimeNow)
		} else {
			dateTimeNow = dateTimeNow.AddDate(0, 0, 1)
			fmt.Println("Not round day", dateTimeNow)
		}

	}
	//***Добавить округление времени для следующего дня!!!!
	return &AvailableTime{
		Days: days,
	}
}

//func FindAvailableTimes(events []*calendar.Event) []string {
//	availableTime := InitAvailableTime()
//
//}
