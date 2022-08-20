package apiobjects

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
}
