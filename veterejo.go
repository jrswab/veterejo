package veterejo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WeatherData holds the information passed back from the OpenWeather API
type WeatherData struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
		TempMin   float32 `json:"temp_min"`
		TempMax   int     `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float32 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

// MakeURL formats the URL for the Open Weather Map API call
func MakeURL(cityID, units, apiID string) (string, error) {
	if apiID == "" {
		return "", fmt.Errorf("no API Key provided to create url")
	}

	if cityID == "" {
		return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather/?units=%s&appid=%s", units, apiID), nil
	}

	if units == "" {
		return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather/?id=%s&units=metrics&appid=%s", cityID, apiID), nil
	}

	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather/?id=%s&units=%s&appid=%s", cityID, units, apiID), nil
}

// GetData calls the OpenWeatherMap API and adds the date to the struct.
func (w *WeatherData) GetData(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		return err
	}
	return nil
}

// GetTemp returns the current temperature that GetData gets in the call.
func (w *WeatherData) GetTemp() float32 {
	return w.Main.Temp
}

// GetMaxTemp returns the current high temperature that GetData gets in the call.
func (w *WeatherData) GetMaxTemp() int {
	return w.Main.TempMax
}

// GetMinTemp returns the current high temperature that GetData gets in the call.
func (w *WeatherData) GetMinTemp() float32 {
	return w.Main.TempMin
}

// GetFeelsLike returns the current "feels like" temperature that GetData gets in the call.
func (w *WeatherData) GetFeelsLike() float32 {
	return w.Main.FeelsLike
}

// GetPressure returns the current pressure that GetData gets in the call.
func (w *WeatherData) GetPressure() int {
	return w.Main.Pressure
}

// GetHumidity returns the current humidity (as a percentage) that GetData gets in the call.
func (w *WeatherData) GetHumidity() int {
	return w.Main.Humidity
}

// GetCoords returns a string of "lat, long"
func (w *WeatherData) GetCoords() string {
	return fmt.Sprintf("%.2f, %.2f", w.Coord.Lat, w.Coord.Lon)
}

// GetVisibility returns a visiblity meter integer
func (w *WeatherData) GetVisibility() int {
	return w.Visibility
}

// GetWindSpeed returns the current windspeed of the city.
func (w *WeatherData) GetWindSpeed() float32 {
	return w.Wind.Speed
}

// GetCloudCoverage returns an int as a percentage of cloud coverage
func (w *WeatherData) GetCloudCoverage() int {
	return w.Clouds.All
}
