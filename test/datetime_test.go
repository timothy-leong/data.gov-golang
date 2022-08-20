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
