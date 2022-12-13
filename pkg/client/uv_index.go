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

func (d *DataGovClient) LatestUVReadings(t time.Time) (*apiobjects.UVReadingAPIResponse, error) {
	// Check cache for a previously fetched value
	if value, ok := d.checkPreviousValue(endpoints.UvIndex, t); ok {
		returnValue := value.(apiobjects.UVReadingAPIResponse)
		return &returnValue, nil
	}

	params := url.Values{}
	params.Add("date_time", datetime.MakeQueryDateTime(t))
	url := string(endpoints.UvIndex) + "?" + params.Encode()

	statusCode, body, err := fasthttp.Get([]byte{}, url)

	if err != nil || statusCode != fasthttp.StatusOK {
		fmt.Printf("Cannot fetch latest UV readings: err = %v, status = %v\n", err, statusCode)
		return nil, err
	}

	var result apiobjects.UVReadingAPIResponse
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("Could not unmarshal latest UV readings", err)
		return nil, err
	}

	d.cache[endpoints.UvIndex] = result
	return &result, nil
}