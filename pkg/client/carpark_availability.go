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
