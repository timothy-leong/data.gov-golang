package client

import (
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/apiobjects"
	"github.com/timothy-leong/data.gov-golang/pkg/endpoints"
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
	case endpoints.UvIndex:
		return time.Hour
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
		case endpoints.UvIndex: // TODO: How to cache endpoints with different params?
			lastFetchedValue := value.(apiobjects.UVReadingAPIResponse)
			refreshRate := d.refreshRate(endpoints.UvIndex)
			if lastFetchedValue.Items[len(lastFetchedValue.Items)-1].Timestamp.Add(refreshRate).After(t) {
				return value, true
			}
		}
	}
	return
}
