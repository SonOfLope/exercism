// Package gigasecond adds a gigasecond to a given time.
package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

// AddGigasecond will, given a moment, determine the moment that would be after a gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	giga := 1.0 * math.Pow10(9)
	t = t.Add(time.Second * time.Duration(giga))
	return t
}
