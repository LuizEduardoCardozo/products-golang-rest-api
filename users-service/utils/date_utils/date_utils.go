package date_utils

import (
	"fmt"
	"time"
)

// Mon Jan 2 15:04:05 -0700 MST 2006

const (
	year        = "2006"
	month       = "01"
	day         = "02"
	hour        = "15"
	minutes     = "04"
	seconds     = "05"
	trueSeconds = "000"
)

var universalDateLayout string = fmt.Sprintf("%s-%s-%sT%s:%s:%s.%s", year, month, day, hour, minutes, seconds, trueSeconds)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(universalDateLayout)
}
