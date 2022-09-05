package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherData struct {
	Temperature string     `json:"temperature"`
	Wind        string     `json:"wind"`
	Description string     `json:"description"`
	Forecasts   []Forecast `json:"forecast"`
}

type Forecast struct {
	Day         string `json:"day"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}

func main() {

	var weatherData WeatherData

	response, err := http.Get("https://goweather.herokuapp.com/weather/Berlin")
	if err != nil {
		fmt.Println(err)
	}

	jsonFile, err := ioutil.ReadAll(response.Body)

	json.Unmarshal(jsonFile, &weatherData)

	fmt.Printf("Today it's %s, the wind is at %s and overall the weather can be descibed as %s\n", weatherData.Temperature, weatherData.Wind, weatherData.Description)
	for i, weather := range weatherData.Forecasts {
		if i == 0 {
			fmt.Printf("In 1 day, it will be %s and the wind will be at %s\n", weather.Temperature, weather.Wind)
		} else {
			fmt.Printf("In %d days, it will be %s and the wind will be at %s\n", i+1, weather.Temperature, weather.Wind)
		}
	}
}
