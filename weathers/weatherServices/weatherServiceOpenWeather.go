package weatherServices

import (
	"WeatherInfo/configManagers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type openWeatherHourlyWeatherResponse struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility int     `json:"visibility"`
		Pop        float64 `json:"pop"`
		Rain       struct {
			H float64 `json:"1h"`
		} `json:"rain"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
		Sunrise    int    `json:"sunrise"`
		Sunset     int    `json:"sunset"`
	} `json:"city"`
}

type openWeatherService struct {
	configManager configManagers.ConfigManager
}

func newOpenWeatherService(configManager configManagers.ConfigManager) openWeatherService {
	return openWeatherService{configManager: configManager}
}

func (o openWeatherService) GetWeatherInfo(req GetWeatherRequest) (GetWeatherResponse, error) {
	res := GetWeatherResponse{}

	url := o.prepareUrl(req)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return res, err
	}

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return res, err
	}

	if httpRes.StatusCode == http.StatusNotFound {
		return res, ErrorLocationNotFound
	} else if httpRes.StatusCode == http.StatusUnauthorized {
		return res, ErrorUnauthorizedRequest
	} else if httpRes.StatusCode != http.StatusOK {
		return res, ErrorWeatherServiceGatewayError
	}

	httpResBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return res, err
	}

	hourlyWeatherResponse := openWeatherHourlyWeatherResponse{}
	err = json.Unmarshal(httpResBody, &hourlyWeatherResponse)
	if err != nil {
		return res, err
	}

	res.Location = fmt.Sprintf("%s / %s", hourlyWeatherResponse.City.Name, hourlyWeatherResponse.City.Country)
	res.WeatherInfoCollection = mapOpenWeatherResponseToWeatherInfoCollection(hourlyWeatherResponse)

	return res, nil
}

func (o openWeatherService) prepareUrl(req GetWeatherRequest) string {
	apiKey := o.configManager.GetStr("OpenWeather.ApiKey")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s&units=metric", req.Location, apiKey)
	return url
}

func mapOpenWeatherResponseToWeatherInfoCollection(owResponse openWeatherHourlyWeatherResponse) []WeatherInfo {
	var infos []WeatherInfo
	for _, item := range owResponse.List {
		info := WeatherInfo{
			Temperature: item.Main.FeelsLike,
			Weather:     item.Weather[0].Main,
			Time:        item.DtTxt,
		}
		infos = append(infos, info)
	}
	return infos
}
