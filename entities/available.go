package entities

import (
	"calendar/config"
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

func InitAvailableTime() *AvailableTime {
	DaysInterval := getInt(config.GetFromEnv("DAYS_INTERVAL"))
	var days []Day
	dateTimeNow := time.Now()
	for i := 0; i < int(DaysInterval); i++ {
		day := InitDay(dateTimeNow)
		days = append(days, *day)
		dateTimeNow = dateTimeNow.AddDate(0, 0, 1)
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
