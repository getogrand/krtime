// Package krtime provides helper functions that ease to use time.Time in Korea.
//
// No more `seoul, err := time.LoadLocation("Asia/Seoul"); if err != nil { ... }; t := time.Now().In(seoul)`.
package krtime

import (
	"fmt"
	"time"
)

// Seoul represents *time.Location of Seoul in Korea.
var Seoul *time.Location

func init() {
	locationName := "Asia/Seoul"

	var err error
	Seoul, err = time.LoadLocation(locationName)
	if err != nil {
		panic(fmt.Errorf("krtime: init: load location %q: %v", locationName, err))
	}
}

// New converts time.Time with the location information set to "Asia/Seoul".
func New(t time.Time) time.Time {
	return t.In(Seoul)
}

// Now returns the current time in location "Asia/Seoul".
func Now() time.Time {
	return New(time.Now())
}
