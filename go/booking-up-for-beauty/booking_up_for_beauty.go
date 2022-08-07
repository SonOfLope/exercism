package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
// => 2019-07-25 13:45:00 +0000 UTC
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	start := time.Now()
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return start.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %02d:%02d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	time, _ := time.Parse("2006-01-02", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return time
}
