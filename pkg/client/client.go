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
	case endpoints.TrafficImages:
		return 20 * time.Second
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
		case endpoints.TrafficImages:
			lastFetchedValue := value.(apiobjects.TrafficImages)
			refreshRate := d.refreshRate(endpoints.TrafficImages)
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
		returnValue := value.(apiobjects.CarparkAvailability)
		return &returnValue, nil
	}

	params := url.Values{}
	params.Add("date_time", datetime.MakeQueryDateTime(t))
	url := string(endpoints.CarparkAvailability) + "?" + params.Encode()

	statusCode, body, err := fasthttp.Get([]byte{}, url)

	if err != nil || statusCode != fasthttp.StatusOK {
		fmt.Printf("Cannot fetch carpark availability: err = %v, status = %v\n", err, statusCode)
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

func (d *DataGovClient) TrafficImages(t time.Time) (*apiobjects.TrafficImages, error) {
	// Check cache for a previously fetched value
	if value, ok := d.checkPreviousValue(endpoints.TrafficImages, t); ok {
		returnValue := value.(apiobjects.TrafficImages)
		return &returnValue, nil
	}

	params := url.Values{}
	params.Add("date_time", datetime.MakeQueryDateTime(t))
	url := string(endpoints.TrafficImages) + "?" + params.Encode()

	statusCode, body, err := fasthttp.Get([]byte{}, url)

	if err != nil || statusCode != fasthttp.StatusOK {
		fmt.Printf("Cannot fetch traffic images: err = %v, status = %v\n", err, statusCode)
		return nil, err
	}

	var result apiobjects.TrafficImages
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("Could not unmarshal traffic image results", err)
		return nil, err
	}

	d.cache[endpoints.TrafficImages] = result
	return &result, nil
}
