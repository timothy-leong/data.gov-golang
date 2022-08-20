package apiobjects

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/datetime"
)

func TestUnmarshalCarparkInfo(t *testing.T) {
	data := `
		  {
			"total_lots": "105",
			"lot_type": "C",
			"lots_available": "60"
		  }`

	var actualCarparkInfo CarparkInfo
	if err := json.Unmarshal([]byte(data), &actualCarparkInfo); err != nil {
		t.Errorf("Cannot unmarshal carpark info: data = %v, err = %v", data, err)
	}

	// Check that the values are correct
	if actualCarparkInfo.TotalLots != 105 {
		t.Errorf("Expected total lots to be 105, got: %v", actualCarparkInfo.TotalLots)
	}

	if actualCarparkInfo.LotType != "C" {
		t.Errorf("Expected lot type to be C, got: %v", actualCarparkInfo.LotType)
	}

	if actualCarparkInfo.LotsAvailable != 60 {
		t.Errorf("Expected lots available to be 60, got: %v", actualCarparkInfo.LotsAvailable)
	}
}

func TestUnmarshalCarpark(t *testing.T) {
	data := `{
		"carpark_info": [
		  {
			"total_lots": "105",
			"lot_type": "C",
			"lots_available": "0"
		  }
		],
		"carpark_number": "HE12",
		"update_datetime": "2022-08-20T08:39:43"
	  }`

	var actualCarpark Carpark
	if err := json.Unmarshal([]byte(data), &actualCarpark); err != nil {
		t.Errorf("Cannot unmarshal carpark: data = %v, err = %v", data, err)
	}

	// Make sure the carpark info was unmarshalled
	if actualCarpark.Info.TotalLots != 105 {
		t.Errorf("Expected total lots to be 105, got: %v", actualCarpark.Info.TotalLots)
	}

	if actualCarpark.CarparkNumber != "HE12" {
		t.Errorf("Expected carpark number to be HE12, got: %v", actualCarpark.CarparkNumber)
	}

	expectedTime := time.Date(2022, 8, 20, 8, 39, 43, 0, datetime.Singapore)
	if expectedTime != actualCarpark.UpdateDatetime {
		t.Errorf("Expected update datetime to be %v, got: %v", expectedTime, actualCarpark.UpdateDatetime)
	}
}

func TestUnmarshalCarparkAvailability(t *testing.T) {
	data := `{
		"items":[
			{
				"timestamp":"2022-08-20T08:52:27+08:00",
				"carpark_data":[
					{
						"carpark_info":[
							{
								"total_lots":"105",
								"lot_type":"C",
								"lots_available":"0"
							}
						],
						"carpark_number":"HE12",
						"update_datetime":"2022-08-20T08:39:43"
					}
				]
			}
		]
	}`

	var actualCarparkAvailability CarparkAvailability
	if err := json.Unmarshal([]byte(data), &actualCarparkAvailability); err != nil {
		t.Errorf("Cannot unmarshal data: %v, err: %v", data, err)
	}

	expectedTime := time.Date(2022, 8, 20, 8, 52, 27, 0, datetime.Singapore)
	if actualTimestamp := actualCarparkAvailability.Timestamp; expectedTime != actualTimestamp {
		t.Errorf("Expected timestamp to be %v, got %v", expectedTime, actualTimestamp)
	}

	// Do a brief check for the carpark data
	if carparkNumber := actualCarparkAvailability.CarparkData[0].CarparkNumber; carparkNumber != "HE12" {
		t.Errorf("Expected carpark number of first carpark to be %v, got %v instead", "HE12", carparkNumber)
	}
}
