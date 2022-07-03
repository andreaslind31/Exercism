package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)

	return parsedTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	start := time.Now()
	layout := "January 2, 2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)

	return parsedTime.Before(start)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)

	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	schedule := Schedule(date)

	return fmt.Sprintf("You have an appointment on %v, %v %d, %d, at %d:%d.", schedule.Weekday(), schedule.Month(), schedule.Day(), schedule.Year(), schedule.Hour(), schedule.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	yearNow := time.Now().Year()
	dateWithAddedYear := fmt.Sprintf("%d-09-15", yearNow)
	layout := "2006-01-02"
	anniversary, _ := time.Parse(layout, dateWithAddedYear)

	return anniversary
}
