package apiobjects

import (
	"encoding/json"
	"testing"
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
