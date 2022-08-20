package apiobjects

import (
	"time"
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
