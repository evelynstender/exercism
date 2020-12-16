package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond adds a gigasecond (10^9) to the time passed
func AddGigasecond(t time.Time) time.Time {

	gigasecond := math.Pow(10, 9)
	return t.Add(time.Second * (time.Duration(gigasecond)))
}
