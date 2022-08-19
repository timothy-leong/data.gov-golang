package apiobjects

import (
	"encoding/json"
	"github.com/timothy-leong/data.gov-golang/pkg/datetime"
	"time"
)

type CarparkAvailability struct {
	Items []Item `json:"items"`
}

type Item struct {
	Timestamp   time.Time `json:"timestamp"`
	CarparkData []Carpark `json:"carpark_data"`
}

type Carpark struct {
	Info           []CarparkInfo `json:"carpark_info"`
	CarparkNumber  string        `json:"carpark_number"`
	UpdateDatetime time.Time     `json:"update_datetime"`
}

type CarparkInfo struct {
	TotalLots     int    `json:"total_lots"`
	LotType       string `json:"lot_type"`
	LotsAvailable int    `json:"lots_available"`
}

func (c *CarparkAvailability) UnmarshalJSON(data []byte) (err error) {
	// Unmarshal the CarparkData and leave the timestamp first
	type IntermediateItem struct {
		Timestamp   json.RawMessage `json:"timestamp"`
		CarparkData []Carpark       `json:"carpark_data"`
	}

	type IntermediateCarparkAvailability struct {
		Items []IntermediateItem `json:"items"`
	}

	var intermediateObject IntermediateCarparkAvailability
	if err = json.Unmarshal(data, &intermediateObject); err != nil {
		return err
	}

	// Copy the CarparkData into the actual object and
	// decode the timestamp on-the-fly
	c.Items = make([]Item, 0, len(intermediateObject.Items))
	for _, intermediateItem := range intermediateObject.Items {
		c.Items = append(c.Items, Item{
			Timestamp:   datetime.ConvertTimestampToTime(string(intermediateItem.Timestamp)),
			CarparkData: intermediateItem.CarparkData,
		})
	}
	return
}

func (c *Carpark) UnmarshalJSON(data []byte) (err error) {
	// Unmarshal the Info and CarparkNumber and leave the UpdateDatetime first
	type IntermediateCarpark struct {
		Info           []CarparkInfo   `json:"carpark_info"`
		CarparkNumber  string          `json:"carpark_number"`
		UpdateDatetime json.RawMessage `json:"update_datetime"`
	}

	var intermediateCarpark IntermediateCarpark
	if err = json.Unmarshal(data, &intermediateCarpark); err != nil {
		return err
	}

	// Copy the Info and CarparkNumber into the actual object
	// and decode the timestamp on the fly
	c.CarparkNumber = intermediateCarpark.CarparkNumber
	c.Info = intermediateCarpark.Info
	c.UpdateDatetime = datetime.ConvertTimestampToTime(string(intermediateCarpark.UpdateDatetime))
	return
}
