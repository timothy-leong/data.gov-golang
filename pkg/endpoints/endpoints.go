package endpoints

type RealAPIEndpoint string

// Real-time APIs
const (
	RealTimeDomain                RealAPIEndpoint = "https://api.data.gov.sg/v1"
	CarparkAvailability           RealAPIEndpoint = RealTimeDomain + "/transport/carpark-availability"
	TaxiAvailability              RealAPIEndpoint = RealTimeDomain + "/transport/taxi-availability"
	IposApplications              RealAPIEndpoint = RealTimeDomain + "/technology/ipos/designs"
	PollutionStandardsIndex       RealAPIEndpoint = RealTimeDomain + "/environment/psi"
	RealtimeWeatherReadings       RealAPIEndpoint = RealTimeDomain + "/environment/air-temperature"
	UltravioletIndex              RealAPIEndpoint = RealTimeDomain + "/environment/uv-index"
	TrafficImages                 RealAPIEndpoint = RealTimeDomain + "/dataset/traffic-images"
	Pm25                          RealAPIEndpoint = RealTimeDomain + "/environment/pm25"
	TwoHourWeatherForecast        RealAPIEndpoint = RealTimeDomain + "/environment/2-hour-weather-forecast"
	TwentyFourHourWeatherForecast RealAPIEndpoint = RealTimeDomain + "/environment/24-hour-weather-forecast"
	FourDayWeatherForecast        RealAPIEndpoint = RealTimeDomain + "/environment/4-day-weather-forecast"
)
