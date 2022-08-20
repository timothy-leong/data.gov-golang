package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/apiobjects"
	"github.com/timothy-leong/data.gov-golang/pkg/datetime"
	"github.com/timothy-leong/data.gov-golang/pkg/endpoints"
	"github.com/valyala/fasthttp"
)

/*
	Client maintains a cache, so that if requests are sent
	faster than the refresh rate of the endpoint, it returns the last fetched data.

	The endpoints tell you the update timestamp of the information you got, and the refresh
	rate of that endpoint, i.e. 1 minute, so any requests within the refresh rate of the
	update timestamp can simply get back
*/
type DataGovClient struct {
	// The key is the API endpoint, and the value is the struct
	// that was last retrieved from this endpoint
	cache map[endpoints.RealAPIEndpoint]any
}

func NewDataGovClient() *DataGovClient {
	return &DataGovClient{
		cache: make(map[endpoints.RealAPIEndpoint]any),
	}
}

func (d *DataGovClient) refreshRate(key endpoints.RealAPIEndpoint) (rate time.Duration) {
	switch key {
	case endpoints.CarparkAvailability:
		return time.Minute
	}
	return
}

func (d *DataGovClient) checkPreviousValue(key endpoints.RealAPIEndpoint, t time.Time) (prevValue any, exists bool) {
	if value, ok := d.cache[key]; ok {
		switch key {
		case endpoints.CarparkAvailability:
			lastFetchedValue := value.(apiobjects.CarparkAvailability)
			refreshRate := d.refreshRate(endpoints.CarparkAvailability)
			if lastFetchedValue.Timestamp.Add(refreshRate).After(t) {
				return value, true
			}
		}
	}
	return
}

func (d *DataGovClient) CarparkAvailability(t time.Time) (*apiobjects.CarparkAvailability, error) {
	// Check cache for a previously fetched value
	if value, ok := d.checkPreviousValue(endpoints.CarparkAvailability, t); ok {
		return value.(*apiobjects.CarparkAvailability), nil
	}

	params := url.Values{}
	params.Add("date_time", datetime.MakeQueryDateTime(t))
	url := string(endpoints.CarparkAvailability) + "?" + params.Encode()

	statusCode, body, err := fasthttp.Get([]byte{}, url)

	if err != nil {
		fmt.Println("Cannot fetch carpark availability:", err)
		return nil, err
	}

	if statusCode != fasthttp.StatusOK {
		fmt.Println("Status was not OK:", statusCode)
		return nil, err
	}

	var result apiobjects.CarparkAvailability
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("Could not unmarshal carpark availability response:", err)
		return nil, err
	}

	d.cache[endpoints.CarparkAvailability] = result
	return &result, nil
}
