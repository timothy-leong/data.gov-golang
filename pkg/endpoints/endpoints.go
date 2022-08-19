package endpoints

// Real-time APIs
const (
	RealTimeDomain          = "https://api.data.gov.sg/v1"
	CarparkAvailability     = RealTimeDomain + "/transport/carpark-availability"
	TaxiAvailability        = RealTimeDomain + "/transport/taxi-availability"
	IposApplications        = RealTimeDomain + "/technology/ipos/designs"
	PollutionStandardsIndex = RealTimeDomain + "/environment/psi"
	RealtimeWeatherReadings = RealTimeDomain + "/environment/air-temperature"
	UltravioletIndex        = RealTimeDomain + "/environment/uv-index"
	TrafficImages           = RealTimeDomain + "/dataset/traffic-images"
	Pm25                    = RealTimeDomain + "/environment/pm25"
	WeatherForecast         = RealTimeDomain + "/environment/2-hour-weather-forecast"
)
