package krtime_test

import (
	"testing"
	"time"

	"github.com/getogrand/krtime"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	someDate := time.Date(1994, 8, 10, 0, 0, 0, 0, time.UTC)
	someDateKr := krtime.New(someDate)

	// should be the same time
	assert.True(t, someDateKr.Equal(someDate))

	// hour should be 0 in UTC, 9 in KST
	assert.Equal(t, 0, someDate.Hour())
	assert.Equal(t, 9, someDateKr.Hour())
}
