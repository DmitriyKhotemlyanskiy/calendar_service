package entities

import (
	"calendar/config"
	"strconv"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"
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
	for startWork.Compare(endWork) < 0 {
		day.TimesArr = append(day.TimesArr, startWork.Format("15:04"))
		startWork = startWork.Add(duration)
	}
	return &day
}

func (d Day) CompareDate(date string) int {
	newDate, _ := time.Parse(time.RFC3339, date)
	if strings.Compare(d.Date, newDate.Format("02 January 2006")) < 0 {
		return -1
	} else if strings.Compare(d.Date, newDate.Format("02 January 2006")) > 0 {
		return 1
	}
	return 0
}

func (d Day) CompareTime(date string, index int) int {
	newTime, _ := time.Parse(time.RFC3339, date)
	if strings.Compare(d.TimesArr[index], newTime.Format("15:04")) < 0 {
		return -1
	} else if strings.Compare(d.TimesArr[index], newTime.Format("15:04")) > 0 {
		return 1
	}
	return 0
}
func (d Day) FindAvailableTimeInDay(events []*calendar.Event) {
	var startEvent string
	var endEvent string
	for _, event := range events {
		startEvent = event.Start.Date
		endEvent = event.End.Date
		if d.CompareDate(startEvent) == 0 {
			for i, j := 0, 0; i < len(d.TimesArr); i++ {
				//***найти все свободные времена между встречами!!!
			}
		}
	}
}
