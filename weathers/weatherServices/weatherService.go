package weatherServices

import "WeatherInfo/configManagers"

type WeatherInfo struct {
	Temperature float64
	Weather     string
	Time        string
}

type GetWeatherRequest struct {
	Location string
}

type GetWeatherResponse struct {
	Location              string
	WeatherInfoCollection []WeatherInfo
}

type WeatherService interface {
	GetWeatherInfo(req GetWeatherRequest) (GetWeatherResponse, error)
}

func NewWeatherService(configManager configManagers.ConfigManager) WeatherService {
	return newOpenWeatherService(configManager)
}
