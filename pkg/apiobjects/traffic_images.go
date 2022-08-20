package apiobjects

import (
	"encoding/json"
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/datetime"
)

type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ImageMetadata struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Md5    string `json:"md5"`
}

type Image struct {
	Timestamp     time.Time     `json:"timestamp"`
	ImageURL      string        `json:"image"`
	Location      Location      `json:"location"`
	CameraID      string        `json:"camera_id"`
	ImageMetadata ImageMetadata `json:"image_metadata"`
}

type TrafficImages struct {
	Timestamp time.Time `json:"timestamp"`
	Images    []Image   `json:"cameras"`
}

func (i *Image) UnmarshalJSON(data []byte) error {
	type intermediateImage struct {
		Timestamp     string        `json:"timestamp"`
		ImageURL      string        `json:"image"`
		Location      Location      `json:"location"`
		CameraID      string        `json:"camera_id"`
		ImageMetadata ImageMetadata `json:"image_metadata"`
	}

	var image intermediateImage
	if err := json.Unmarshal(data, &image); err != nil {
		return err
	}

	i.Timestamp = datetime.ConvertTimestampToTime(image.Timestamp)
	i.ImageURL = image.ImageURL
	i.Location = image.Location
	i.CameraID = image.CameraID
	i.ImageMetadata = image.ImageMetadata
	return nil
}

func (t *TrafficImages) UnmarshalJSON(data []byte) error {
	type intermediateTrafficImage struct {
		Timestamp string  `json:"timestamp"`
		Images    []Image `json:"cameras"`
	}

	var obj intermediateTrafficImage
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	t.Timestamp = datetime.ConvertTimestampToTime(obj.Timestamp)
	t.Images = obj.Images
	return nil
}
