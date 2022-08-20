package test

import (
	"testing"
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/datetime"
)

func TestMakeQueryDateTime(t *testing.T) {
	time := time.Date(2022, 1, 2, 3, 4, 5, 0, time.UTC)
	queryString := datetime.MakeQueryDateTime(time)

	if queryString != "2022-01-02T03:04:05" {
		t.Errorf("Wrong conversion: time = %v, queryString = %v", time, queryString)
	}
}

func TestConvertTimestampToTime(t *testing.T) {
	timestamp := "2022-01-02T03:03:27+08:00"
	timeValue := datetime.ConvertTimestampToTime(timestamp)
	expectedTime := time.Date(2022, 1, 2, 3, 3, 27, 0, time.UTC)

	if timeValue != expectedTime {
		t.Errorf("Wrong conversion: timestamp = %v, converted time = %v", timestamp, timeValue)
	}
}
