package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Config func to get env value from file ---
func GetFromEnv(varName string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(varName)
}

func GetService(client *http.Client) *calendar.Service {
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	return srv
}

func StartStop(timeNow time.Time) (startWork, endWork time.Time) {
	yyyy, mm, dd := timeNow.Date()
	startWork, _ = time.Parse("15:04", GetFromEnv("WORK_TIME_FROM"))
	startWork = time.Date(yyyy, mm, dd, startWork.Hour(), startWork.Minute(), 0, 0, timeNow.Location())
	endWork, _ = time.Parse("15:04", GetFromEnv("WORK_TIME_TO"))
	endWork = time.Date(yyyy, mm, dd, endWork.Hour(), endWork.Minute(), 0, 0, timeNow.Location())
	return startWork, endWork
}
