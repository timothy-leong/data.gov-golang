package datetime

import (
	"testing"
	"time"
)

func TestMakeQueryDateTime(t *testing.T) {
	time := time.Date(2022, 1, 2, 3, 4, 5, 0, Singapore)
	queryString := MakeQueryDateTime(time)

	if queryString != "2022-01-02T03:04:05" {
		t.Errorf("Wrong conversion: time = %v, queryString = %v", time, queryString)
	}
}

func TestConvertTimestampToTime(t *testing.T) {
	for _, timestamp := range []string{
		"2022-01-02T03:03:27+08:00",
		"2022-01-02T03:03:27"} {
		timeValue := ConvertTimestampToTime(timestamp)
		expectedTime := time.Date(2022, 1, 2, 3, 3, 27, 0, Singapore)

		if timeValue != expectedTime {
			t.Errorf("Wrong conversion: timestamp = %v, converted time = %v", timestamp, timeValue)
		}
	}

}
