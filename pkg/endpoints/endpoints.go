package endpoints

type RealAPIEndpoint string

// Real-time APIs
const (
	RealTimeDomain      RealAPIEndpoint = "https://api.data.gov.sg/v1"
	CarparkAvailability RealAPIEndpoint = RealTimeDomain + "/transport/carpark-availability"
	TrafficImages       RealAPIEndpoint = RealTimeDomain + "/transport/traffic-images"
	UvIndex             RealAPIEndpoint = RealTimeDomain + "/environment/uv-index"
)
