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
	TODO implement a cache in this client, so that if requests are sent
	faster than the refresh rate of the endpoint, return the last fetched data.
*/
type DataGovClient struct{}

func NewDataGovClient() *DataGovClient {
	return &DataGovClient{}
}

func (d *DataGovClient) CarparkAvailability(t time.Time) (*apiobjects.CarparkAvailability, error) {
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
