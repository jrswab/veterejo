package veterejo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	v "github.com/jrswab/veterejo"
)

// RunTestServer creates a temporary webserver to serve the test data defined below.
func RunTestServer(unit string) *v.WeatherData {
	w := new(v.WeatherData)
	imperialJSON := getImperialTestData()
	metricJSON := getMetricTestData()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if unit == "imperial" {
			fmt.Fprintln(w, imperialJSON)
		}
		if unit == "metric" {
			fmt.Fprintln(w, metricJSON)
		}
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&w)
	return w
}

func getImperialTestData() string {
	return `{
		"coord": {
			"lon": -79.96,
			"lat": 40.56
		},
		"weather": [
			{
				"id": 800,
				"main": "Clear",
				"description": "clear sky",
				"icon": "01n"
			}
		],
		"base": "stations",
		"main": {
			"temp": 30.9,
			"feels_like": 24.13,
			"temp_min": 26.01,
			"temp_max": 36,
			"pressure": 1031,
			"humidity": 68
		},
		"visibility": 16093,
		"wind": {
			"speed": 3.42,
			"deg": 239
		},
		"clouds": {
			"all": 1
		},
		"dt": 1583637756,
		"sys": {
			"type": 1,
			"id": 3511,
			"country": "US",
			"sunrise": 1583581421,
			"sunset": 1583623083
		},
		"timezone": -18000,
		"id": 5178165,
		"name": "Allison Park",
		"cod": 200
	}`
}

func getMetricTestData() string {
	return `{
		"coord": {
			"lon": -79.96,
			"lat": 40.56
		},
		"weather": [
			{
				"id": 800,
				"main": "Clear",
				"description": "clear sky",
				"icon": "01n"
			}
		],
		"base": "stations",
		"main": {
			"temp": -0.73,
			"feels_like": -4.73,
			"temp_min": -3.33,
			"temp_max": 2,
			"pressure": 1031,
			"humidity": 55
		},
		"visibility": 16093,
		"wind": {
			"speed": 1.5,
			"deg": 200
		},
		"clouds": {
			"all": 1
		},
		"dt": 1583638861,
		"sys": {
			"type": 1,
			"id": 3510,
			"country": "US",
			"sunrise": 1583581421,
			"sunset": 1583623083
		},
		"timezone": -18000,
		"id": 5178165,
		"name": "Allison Park",
		"cod": 200
	}`
}
