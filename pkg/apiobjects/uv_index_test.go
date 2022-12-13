package apiobjects

import (
	"encoding/json"
	"testing"
	"time"
)

func TestUnmarshalReading(t *testing.T) {
	data := `{
		"value": 1,
		"timestamp": "2022-12-13T09:00:00+08:00"
	  }`

	var reading UVReading
	if err := json.Unmarshal([]byte(data), &reading); err != nil {
		t.Error(err)
	}

	// Briefly check the time
	if reading.Timestamp.Month() != time.December {
		t.Errorf("Month is not December, but %v", reading.Timestamp.Month())
	}
}

func TestUnmarshalListOfReadings(t *testing.T) {
	data := `{
		"items": [
		  {
			"timestamp": "2022-12-13T07:00:00+08:00",
			"update_timestamp": "2022-12-13T07:46:07+08:00",
			"index": [
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T08:00:00+08:00",
			"update_timestamp": "2022-12-13T08:07:07+08:00",
			"index": [
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T09:00:00+08:00",
			"update_timestamp": "2022-12-13T09:08:07+08:00",
			"index": [
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T10:00:00+08:00",
			"update_timestamp": "2022-12-13T10:06:07+08:00",
			"index": [
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T11:00:00+08:00",
			"update_timestamp": "2022-12-13T11:12:07+08:00",
			"index": [
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T12:00:00+08:00",
			"update_timestamp": "2022-12-13T12:47:07+08:00",
			"index": [
			  {
				"value": 5,
				"timestamp": "2022-12-13T12:00:00+08:00"
			  },
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T13:00:00+08:00",
			"update_timestamp": "2022-12-13T13:05:07+08:00",
			"index": [
			  {
				"value": 5,
				"timestamp": "2022-12-13T13:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T12:00:00+08:00"
			  },
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T14:00:00+08:00",
			"update_timestamp": "2022-12-13T14:06:07+08:00",
			"index": [
			  {
				"value": 2,
				"timestamp": "2022-12-13T14:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T13:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T12:00:00+08:00"
			  },
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T15:00:00+08:00",
			"update_timestamp": "2022-12-13T15:06:07+08:00",
			"index": [
			  {
				"value": 0,
				"timestamp": "2022-12-13T15:00:00+08:00"
			  },
			  {
				"value": 2,
				"timestamp": "2022-12-13T14:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T13:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T12:00:00+08:00"
			  },
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T16:00:00+08:00",
			"update_timestamp": "2022-12-13T16:08:07+08:00",
			"index": [
			  {
				"value": 0,
				"timestamp": "2022-12-13T16:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T15:00:00+08:00"
			  },
			  {
				"value": 2,
				"timestamp": "2022-12-13T14:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T13:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T12:00:00+08:00"
			  },
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T17:00:00+08:00",
			"update_timestamp": "2022-12-13T17:09:07+08:00",
			"index": [
			  {
				"value": 0,
				"timestamp": "2022-12-13T17:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T16:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T15:00:00+08:00"
			  },
			  {
				"value": 2,
				"timestamp": "2022-12-13T14:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T13:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T12:00:00+08:00"
			  },
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T18:00:00+08:00",
			"update_timestamp": "2022-12-13T18:07:07+08:00",
			"index": [
			  {
				"value": 0,
				"timestamp": "2022-12-13T18:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T17:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T16:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T15:00:00+08:00"
			  },
			  {
				"value": 2,
				"timestamp": "2022-12-13T14:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T13:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T12:00:00+08:00"
			  },
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  },
		  {
			"timestamp": "2022-12-13T19:00:00+08:00",
			"update_timestamp": "2022-12-13T19:06:07+08:00",
			"index": [
			  {
				"value": 0,
				"timestamp": "2022-12-13T19:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T18:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T17:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T16:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T15:00:00+08:00"
			  },
			  {
				"value": 2,
				"timestamp": "2022-12-13T14:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T13:00:00+08:00"
			  },
			  {
				"value": 5,
				"timestamp": "2022-12-13T12:00:00+08:00"
			  },
			  {
				"value": 4,
				"timestamp": "2022-12-13T11:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T10:00:00+08:00"
			  },
			  {
				"value": 1,
				"timestamp": "2022-12-13T09:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T08:00:00+08:00"
			  },
			  {
				"value": 0,
				"timestamp": "2022-12-13T07:00:00+08:00"
			  }
			]
		  }
		],
		"api_info": {
		  "status": "healthy"
		}
	  }`

	var response UVReadingAPIResponse
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		t.Error(err)
	}

	// Briefly check the length of items and some values
	if len(response.Items) != 13 {
		t.Errorf("Expected 13 items, got %v", len(response.Items))
	}

	if response.Items[0].Index[0].Value != 0 {
		t.Errorf("Expected first value reading to be 0, got %v", response.Items[0].Index[0].Value)
	}
}
