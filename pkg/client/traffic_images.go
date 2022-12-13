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
