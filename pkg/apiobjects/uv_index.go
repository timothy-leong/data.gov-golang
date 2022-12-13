package apiobjects

import (
	"encoding/json"
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/datetime"
)

type UVReading struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type UVHourlyReadingDump struct {
	Timestamp       time.Time   `json:"timestamp"`
	UpdateTimestamp time.Time   `json:"update_timestamp"`
	Index           []UVReading `json:"index"`
}

type UVReadingAPIResponse struct {
	Items []UVHourlyReadingDump `json:"items"`
}

func (u *UVReading) UnmarshalJSON(data []byte) error {
	type intermediateReading struct {
		Value     int    `json:"value"`
		Timestamp string `json:"timestamp"`
	}

	var reading intermediateReading
	if err := json.Unmarshal(data, &reading); err != nil {
		return err
	}

	u.Value = reading.Value
	u.Timestamp = datetime.ConvertTimestampToTime(reading.Timestamp)
	return nil
}

func (u *UVHourlyReadingDump) UnmarshalJSON(data []byte) error {
	type intermediateReadingDump struct {
		Timestamp       string      `json:"timestamp"`
		UpdateTimestamp string      `json:"update_timestamp"`
		Index           []UVReading `json:"index"`
	}

	var readingDump intermediateReadingDump
	if err := json.Unmarshal(data, &readingDump); err != nil {
		return err
	}

	u.UpdateTimestamp = datetime.ConvertTimestampToTime(readingDump.UpdateTimestamp)
	u.Timestamp = datetime.ConvertTimestampToTime(readingDump.Timestamp)
	u.Index = readingDump.Index
	return nil
}
