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

func roundTimeDuration(date time.Time) time.Time {
	yy, mm, dd := date.Date()
	hh, MM, _ := date.Clock()
	num, _ := strconv.Atoi(config.GetFromEnv("MEETING_DURATION"))
	dif := num - MM
	if dif < 0 {
		dif = (num + dif) + num
	}
	//***Округлять время до ближайшего следующего времени!!!
	roundTime := time.Date(yy, mm, dd, hh, MM+dif, 0, 0, date.Location())
	return roundTime
}

func InitDay(date time.Time) *Day {
	var day Day
	day.Date = date.Format("02 January 2006")

	num, _ := strconv.Atoi(config.GetFromEnv("MEETING_DURATION"))
	duration := time.Duration(int64(num)) * time.Minute
	startWork, endWork := config.StartStop(date)
	if date.After(startWork) && date.Before(endWork) {
		date = roundTimeDuration(date)
		startWork = date
	}
	for startWork.Compare(endWork) <= 0 {
		day.TimesArr = append(day.TimesArr, startWork.Format("15:04"))
		startWork = startWork.Add(duration)
	}
	return &day
}
