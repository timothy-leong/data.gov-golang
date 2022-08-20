package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/apiobjects"
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
	cache map[string]any
}

func NewDataGovClient() *DataGovClient {
	return &DataGovClient{
		cache: make(map[string]any),
	}
}

func (d *DataGovClient) CarparkAvailability(t time.Time) (*apiobjects.CarparkAvailability, error) {
	// Check cache for a previously fetched value
	if value, ok := d.cache[endpoints.CarparkAvailability]; ok {
		// Check when it was updated
		lastFetchedValue := value.(apiobjects.CarparkAvailability)
		const refreshRate = time.Minute

		if lastFetchedValue.Items[0].Timestamp.Add(refreshRate).After(time.Now()) {
			return &lastFetchedValue, nil
		}
	}

	statusCode, body, err := fasthttp.Get([]byte{}, endpoints.CarparkAvailability)

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

	return &result, nil
}
