package apiobjects

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/datetime"
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

func (c *CarparkAvailability) UnmarshalJSON(data []byte) error {
	// Unmarshal the CarparkData and leave the timestamp first
	type IntermediateItem struct {
		Timestamp   string    `json:"timestamp"`
		CarparkData []Carpark `json:"carpark_data"`
	}

	type IntermediateCarparkAvailability struct {
		Items []IntermediateItem `json:"items"`
	}

	var intermediateObject IntermediateCarparkAvailability
	if err := json.Unmarshal(data, &intermediateObject); err != nil {
		return err
	}

	// Copy the CarparkData into the actual object and
	// decode the timestamp on-the-fly
	c.Items = make([]Item, 0, len(intermediateObject.Items))
	for _, intermediateItem := range intermediateObject.Items {
		c.Items = append(c.Items, Item{
			Timestamp:   datetime.ConvertTimestampToTime(intermediateItem.Timestamp),
			CarparkData: intermediateItem.CarparkData,
		})
	}
	return nil
}

func (c *Carpark) UnmarshalJSON(data []byte) error {
	// Unmarshal the Info and CarparkNumber and convert the UpdateDatetime to string first
	type IntermediateCarpark struct {
		Info           []CarparkInfo `json:"carpark_info"`
		CarparkNumber  string        `json:"carpark_number"`
		UpdateDatetime string        `json:"update_datetime"`
	}

	var intermediateCarpark IntermediateCarpark
	if err := json.Unmarshal(data, &intermediateCarpark); err != nil {
		return err
	}

	// Copy the Info and CarparkNumber into the actual object
	// and decode the timestamp on the fly
	c.CarparkNumber = intermediateCarpark.CarparkNumber
	c.Info = intermediateCarpark.Info
	c.UpdateDatetime = datetime.ConvertTimestampToTime(intermediateCarpark.UpdateDatetime)
	return nil
}

func (c *CarparkInfo) UnmarshalJSON(data []byte) error {
	// Unmarshal everything into strings first
	type IntermediateCarparkInfo struct {
		TotalLots     string `json:"total_lots"`
		LotType       string `json:"lot_type"`
		LotsAvailable string `json:"lots_available"`
	}

	var intermediateCarparkInfo IntermediateCarparkInfo
	if err := json.Unmarshal(data, &intermediateCarparkInfo); err != nil {
		return err
	}

	// Copy the lot type and translate the total lots and lots available
	c.LotType = intermediateCarparkInfo.LotType

	totalLots, err := strconv.Atoi(intermediateCarparkInfo.TotalLots)
	if err != nil {
		return err
	}
	c.TotalLots = totalLots

	lotsAvailable, err := strconv.Atoi(intermediateCarparkInfo.LotsAvailable)
	if err != nil {
		return err
	}
	c.LotsAvailable = lotsAvailable
	return nil
}
