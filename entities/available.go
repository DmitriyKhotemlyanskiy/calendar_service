package entities

import (
	"calendar/config"
	"strconv"
)

type AvailableTime struct {
	DaysInterval    int64  //DAYS
	MeetingDuration int64  //MINUTES
	WorkTimeFrom    string //EXAMPLE -> "15:04"
	WorkTimeTo      string //EXAMPLE -> "18:00"
	Days            []Day
}

func getInt(str string) int64 {
	num, _ := strconv.Atoi(str)
	return int64(num)
}

func InitAvailableTime() *AvailableTime {
	return &AvailableTime{
		DaysInterval:    getInt(config.GetFromEnv("DAYS_INTERVAL")),
		MeetingDuration: getInt(config.GetFromEnv("MEETING_DURATION")),
		WorkTimeFrom:    config.GetFromEnv("WORK_TIME_FROM"),
		WorkTimeTo:      config.GetFromEnv("WORK_TIME_TO"),
	}
}

//func FindAvailableTimes(events []*calendar.Event) []string {
//	//availableTime := InitAvailableTime()
//
//}
