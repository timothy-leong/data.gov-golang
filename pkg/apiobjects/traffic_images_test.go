package apiobjects

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalImage(t *testing.T) {
	data := `{
		"timestamp": "2022-08-20T14:11:46+08:00",
		"image": "https://images.data.gov.sg/api/traffic-images/2022/08/daddc13e-2374-4859-b6ab-ee2abbe352cc.jpg",
		"location": {
		  "latitude": 1.319541067,
		  "longitude": 103.8785627
		},
		"camera_id": "1002",
		"image_metadata": {
		  "height": 240,
		  "width": 320,
		  "md5": "f2ba086e702063b1396a989e3698421d"
		}
	  }`

	var image Image
	if err := json.Unmarshal([]byte(data), &image); err != nil {
		t.Error(err)
	}

	// Check some details
	if image.CameraID != "1002" {
		t.Errorf("Camera ID is not 1002, but %v", image.CameraID)
	}
}

func TestUnmarshalImages(t *testing.T) {
	data := `{
		"items":[
			{
				"timestamp":"2022-08-20T14:12:06+08:00",
				"cameras":[
					{
						"timestamp":"2022-08-20T14:11:46+08:00",
						"image":"https://images.data.gov.sg/api/traffic-images/2022/08/daddc13e-2374-4859-b6ab-ee2abbe352cc.jpg",
						"location":{
							"latitude":1.319541067,
							"longitude":103.8785627
						},
						"camera_id":"1002",
						"image_metadata":{
							"height":240,
							"width":320,
							"md5":"f2ba086e702063b1396a989e3698421d"
						}
					}
				]
			}
		]
	}`

	var images TrafficImages
	if err := json.Unmarshal([]byte(data), &images); err != nil {
		t.Error(err)
	}

	// Check some details
	if images.Images[0].CameraID != "1002" {
		t.Errorf("Expected the camera ID to be 1002, got %v instead", images.Images[0].CameraID)
	}
}
