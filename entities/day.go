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

// Compare time d.Day.TimesArr[index] with date. If TimesArr[index] < date -> retrun -1, if TimesArr[index] == date -> return 0, if TimesArr[index] > date -> return 1
func (d Day) CompareTime(date string, index int) int {
	newTime, _ := time.Parse(time.RFC3339, date)
	if strings.Compare(d.TimesArr[index], newTime.Format("15:04")) < 0 {
		return -1
	} else if strings.Compare(d.TimesArr[index], newTime.Format("15:04")) > 0 {
		return 1
	}
	return 0
}

func (d *Day) FindAvailableTimeInDay(events []*calendar.Event) {
	var startEvent string
	var endEvent string
	var newTimesArr []string
	start := 0
	for _, event := range events {
		startEvent = event.Start.DateTime
		endEvent = event.End.DateTime
		fmt.Println("Start time event: ", startEvent, "End time event: ", endEvent)
		if d.CompareDate(startEvent) == 0 {
			for start < len(d.TimesArr) {
				if d.CompareTime(startEvent, start) < 0 && d.CompareTime(endEvent, start) < 0 {
					newTimesArr = append(newTimesArr, d.TimesArr[start])
					start++
				} else if d.CompareTime(startEvent, start) == 0 && d.CompareTime(endEvent, start) < 0 {
					start++
					continue
				} else if d.CompareTime(startEvent, start) > 0 && d.CompareTime(endEvent, start) == 0 {
					newTimesArr = append(newTimesArr, d.TimesArr[start])
					start++
					break
				}
			}
		}
	}
	fmt.Println("New Day ARR: ", newTimesArr)
	d.TimesArr = newTimesArr
}
