package datetime

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	timestampPattern = regexp.MustCompile("^([0-9]{4})-([0-9]{2})-([0-9]{2})T([0-9]{2}):([0-9]{2}):([0-9]{2})+08:00$")
)

func MakeQueryDateTime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func ConvertTimestampToTime(s string) time.Time {
	matches := timestampPattern.FindStringSubmatch(s)
	atoi := strconv.Atoi
	year, _ := atoi(matches[1])
	month, _ := atoi(matches[2])
	day, _ := atoi(matches[3])
	hour, _ := atoi(matches[4])
	minute, _ := atoi(matches[5])
	second, _ := atoi(matches[6])
	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
}
