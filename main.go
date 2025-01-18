package main

import (
	"calendar/handlers"

	"github.com/gin-gonic/gin"
)

//http://localhost:8081/gcalendar_service - redirect URL

//	func myHandler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "Hello from Go")
//	}
//
//	func server() {
//		http.HandleFunc("/", myHandler)
//		fmt.Println("Server is running at port: 8080")
//		panic(http.ListenAndServe(":8080", nil))
//	}

func main() {
	//go server()
	//client := config.Auth()
	//srv := config.GetService(client)

	//event := &calendar.Event{
	//	Summary:     "Google I/O 2024",
	//	Location:    "800 Howard St., San Francisco, CA 94103",
	//	Description: "A chance to hear more about Google's developer products.",
	//	Start: &calendar.EventDateTime{
	//		DateTime: "2024-07-25T10:00:00",
	//		TimeZone: "Asia/Jerusalem",
	//	},
	//	End: &calendar.EventDateTime{
	//		DateTime: "2024-07-25T12:00:00",
	//		TimeZone: "Asia/Jerusalem",
	//	},
	//	//Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
	//}
	//
	//calendarId := "protattoo.bot@gmail.com"
	//event, err = srv.Events.Insert(calendarId, event).Do()
	//if err != nil {
	//	log.Fatalf("Unable to create event. %v\n", err)
	//}
	//fmt.Printf("Event created: %s\n", event.HtmlLink)

	//items := entities.Calendar{}.GetUpcomingEvents(srv)
	//for _, item := range items {
	//	date := item.Start.DateTime
	//	end := item.End.DateTime
	//	if date == "" {
	//		date = item.Start.Date
	//		end = item.End.Date
	//	}
	//
	//	fmt.Printf("%v (%v) start time: %v - end time: %v\n", item.Summary, item.Description, date, end)
	//	//	startdata := entities.NewDateTime(date)
	//	//	enddata := entities.NewDateTime(end)
	//	//date := "2024-08-04T23:00:00+03:00"
	//	//end := "2024-08-05T00:15:00+03:00"
	//	//fmt.Printf("start time: 2024-08-04T23:00:00+03:00 - end time: 2024-08-05T00:15:00+03:00\n")
	//	////startdata := entities.NewDateTime(date)
	//	//enddata := entities.NewDateTime(end)
	//	//fmt.Println(enddata.GetDate())
	//	//fmt.Println(enddata.GetDateUTC())
	//	////fmt.Println(enddata.AddTime(startdata.TimeDifference(*enddata)))
	//	//fmt.Println(enddata.GetDate())
	//	//fmt.Println(enddata.GetDayStr())
	//	//fmt.Println(enddata.GetMonthStr())
	//	//fmt.Println(enddata.GetYearStr())
	//	//}
	//
	//}
	//availableTimes := entities.FindAvailableTimes(items)
	//for _, day := range availableTimes.Days {
	//	if strings.Compare(day.Date, "20 November 2024") == 0 {
	//		fmt.Println(day.TimesArr)
	//	}
	//}
	//connect := config.Connect()
	//db := connect.GetDB()
	//rows, err := db.Query("SELECT * FROM events")
	//if err != nil {
	//	panic(err)
	//}
	//for rows.Next() {
	//	var index int
	//	var event_id string
	//	var user_id int
	//	var date string
	//	var start_time string
	//	var end_time string
	//	var text string
	//	if err := rows.Scan(&index, &event_id, &user_id, &date, &start_time, &end_time, &text); err != nil {
	//		log.Fatalln("can't read from row: ", err)
	//	}
	//	fmt.Printf("Index: %d, ID: %s, Name: %d, Lastname: %s, UserName: %s, email: %s, Tel: %s\n", index, event_id, user_id, date, start_time, end_time, text)
	//}
	//
	//defer rows.Close()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.GET("/availableTimes", handlers.GetAvailableTimes)
	v1.POST("/event", handlers.CreateEvent)
	v1.DELETE("/event", handlers.DeleteEvent)
	v1.GET("/event", handlers.GetMyEvents)
	router.Run("localhost:8080")
}
