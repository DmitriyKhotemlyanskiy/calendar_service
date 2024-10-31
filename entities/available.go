package entities

import (
	"calendar/config"
	"strconv"
	"time"

	"google.golang.org/api/calendar/v3"
)

type AvailableTime struct {
	DaysInterval    int64  //DAYS
	MeetingDuration int64  //MINUTES
	WorkTimeFrom    string //EXAMPLE -> "15:04"
	WorkTimeTo      string //EXAMPLE -> "18:00"
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

func FindAvailableTimes(events []*calendar.Event) []string {
	var availableTime *AvailableTime
	availableTime = InitAvailableTime()
	var availableTimes []string
	var slotDuration = time.Duration(availableTime.MeetingDuration) * time.Minute

	// Устанавливаем начальное и конечное время для анализа
	workingStart, _ := time.Parse("15:04", availableTime.WorkTimeFrom) // Начало рабочего дня
	workingEnd, _ := time.Parse("15:04", availableTime.WorkTimeTo)     // Конец рабочего дня

	for i := 0; i <= len(events); i++ {
		var currentEnd time.Time
		// Определяем начало и конец текущего слота
		if i == len(events) {
			currentEnd = workingEnd
		} else {
			currentEnd, _ = time.Parse(time.RFC3339, events[i].Start.DateTime)
		}

		// Если есть временной интервал до следующего события
		for workingStart.Before(currentEnd) && workingStart.Add(slotDuration).Before(currentEnd) {

			availableTimes = append(availableTimes, workingStart.Format("15:04"))
			workingStart = workingStart.Add(slotDuration)
			if workingStart.Format("15:04") == availableTime.WorkTimeTo || workingStart.Format("15:04") == "00:00" {
				break
			}
		}

		// Обновляем рабочее время после события
		if i < len(events) {
			workingStart, _ = time.Parse(time.RFC3339, events[i].End.DateTime)
		}
	}

	return availableTimes
}
