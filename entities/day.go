package entities

import (
	"calendar/config"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"
)

type Day struct {
	Date     string            `json:"date"`
	TimesArr map[string]string `json:"timesarr"`
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
	day.TimesArr = make(map[string]string)
	num, _ := strconv.Atoi(config.GetFromEnv("MEETING_DURATION"))
	duration := time.Duration(int64(num)) * time.Minute
	startWork, endWork := config.StartStop(date)
	if date.After(startWork) && date.Before(endWork) {
		date = roundTimeDuration(date)
		startWork = date
	}
	for startWork.Compare(endWork) < 0 {
		day.TimesArr[startWork.Format("15:04")] = startWork.Format("15:04")
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

func countDuration(event *calendar.Event) int {
	count := 0
	defaultDuration, _ := strconv.Atoi(config.GetFromEnv("MEETING_DURATION"))
	duration := time.Duration(int64(defaultDuration)) * time.Minute
	start, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	end, _ := time.Parse(time.RFC3339, event.End.DateTime)
	for start.Before(end) {
		count++
		start = start.Add(duration)
	}
	fmt.Println("counter value --> ", count)
	return count
}

func getCountedArr(event *calendar.Event) *[]string {
	var arr []string
	defaultDuration, _ := strconv.Atoi(config.GetFromEnv("MEETING_DURATION"))
	duration := time.Duration(int64(defaultDuration)) * time.Minute
	start, _ := time.Parse(time.RFC3339, event.Start.DateTime)
	end, _ := time.Parse(time.RFC3339, event.End.DateTime)
	for start.Before(end) {
		arr = append(arr, start.Format("15:04"))
		start = start.Add(duration)
	}
	return &arr
}

func (d *Day) FindAvailableTimeInDay(events []*calendar.Event) {
	var eventTimes []string
	for _, event := range events {
		newEventS := event.Start.DateTime
		if d.CompareDate(newEventS) == 0 {
			if countDuration(event) <= 1 {
				t, _ := time.Parse(time.RFC3339, newEventS)
				eventTimes = append(eventTimes, t.Format("15:04"))
			} else {
				tempArr := getCountedArr(event)
				eventTimes = append(eventTimes, *tempArr...)
				fmt.Println("TempArr: --> ", *tempArr)
			}
		}
	}
	for _, event := range eventTimes {
		delete(d.TimesArr, event)
	}
}
