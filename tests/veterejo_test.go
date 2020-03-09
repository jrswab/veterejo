package veterejo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	v "github.com/jrswab/veterejo"
)

func TestMakeURL(t *testing.T) {
	type args struct {
		cityID string
		units  string
		apiID  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test with all args",
			args: args{
				cityID: "0000",
				units:  "metric",
				apiID:  "0x0p0q",
			},
			want:    "https://api.openweathermap.org/data/2.5/weather/?id=0000&units=metric&appid=0x0p0q",
			wantErr: false,
		},
		{
			name: "Test without CityID",
			args: args{
				cityID: "",
				units:  "metric",
				apiID:  "0x0p0q",
			},
			want:    "https://api.openweathermap.org/data/2.5/weather/?units=metric&appid=0x0p0q",
			wantErr: false,
		},
		{
			name: "Test without units",
			args: args{
				cityID: "0000",
				units:  "",
				apiID:  "0x0p0q",
			},
			want:    "https://api.openweathermap.org/data/2.5/weather/?id=0000&appid=0x0p0q",
			wantErr: false,
		},
		{
			name: "Test without API Key",
			args: args{
				cityID: "0000",
				units:  "imperial",
				apiID:  "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := v.MakeURL(tt.args.cityID, tt.args.units, tt.args.apiID)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MakeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
	}))
	defer ts.Close()
	fakeURL := ""

	tests := []struct {
		name    string
		url     string
		data    *v.WeatherData
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test with valid url",
			url:     ts.URL,
			wantErr: false,
		},
		{
			name:    "Test with an empty string as the url",
			url:     fakeURL,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.GetData(tt.url); (err != nil) != tt.wantErr {
				t.Errorf("WeatherData.GetData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetMaxTemp(t *testing.T) {
	tests := []struct {
		name string
		data *v.WeatherData
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Temperature returns integer on a correct call to the api (imperial)",
			data: RunTestServer("imperial"),
			want: 36,
		},
		{
			name: "Temperature returns integer on a correct call to the api (metric)",
			data: RunTestServer("metric"),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.GetMaxTemp()
			if got != tt.want {
				t.Errorf("GetMaxTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMinTemp(t *testing.T) {
	tests := []struct {
		name string
		data *v.WeatherData
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "Temperature returns float on a correct call to the api (imperial)",
			data: RunTestServer("imperial"),
			want: 26.01,
		},
		{
			name: "Temperature returns float on a correct call to the api (metric)",
			data: RunTestServer("metric"),
			want: -3.33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.GetMinTemp()
			if got != tt.want {
				t.Errorf("GetMinTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFeelsLike(t *testing.T) {
	tests := []struct {
		name string
		data *v.WeatherData
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "Returns a temperature float on a correct call to the api (imperial)",
			data: RunTestServer("imperial"),
			want: 24.13,
		},
		{
			name: "Returns a temperature float on a correct call to the api (metric)",
			data: RunTestServer("metric"),
			want: -4.73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.GetFeelsLike()
			if got != tt.want {
				t.Errorf("GetMinTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPressure(t *testing.T) {
	tests := []struct {
		name string
		data *v.WeatherData
		want int
	}{
		// TODO: Add test cases.
		{
			name: "GetPressure returns integer on a correct call to the api (imperial)",
			data: RunTestServer("imperial"),
			want: 1031,
		},
		{
			name: "GetPressure returns integer on a correct call to the api (metric)",
			data: RunTestServer("metric"),
			want: 1031,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.GetPressure()
			if got != tt.want {
				t.Errorf("GetMaxTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHumidity(t *testing.T) {
	tests := []struct {
		name string
		data *v.WeatherData
		want int
	}{
		// TODO: Add test cases.
		{
			name: "GetHumidity returns integer on a correct call to the api (imperial)",
			data: RunTestServer("imperial"),
			want: 68,
		},
		{
			name: "GetHumidity returns integer on a correct call to the api (metric)",
			data: RunTestServer("metric"),
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.GetHumidity()
			if got != tt.want {
				t.Errorf("GetMaxTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCoords(t *testing.T) {
	tests := []struct {
		name string
		data *v.WeatherData
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Return coords as string on a correct call to the api (imperial)",
			data: RunTestServer("imperial"),
			want: "50.56, -79.96",
		},
		{
			name: "Return coords as string on a correct call to the api (metric)",
			data: RunTestServer("metric"),
			want: "50.56, -79.96",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.GetCoords()
			if got != tt.want {
				t.Errorf("GetCoords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWindSpeed(t *testing.T) {
	tests := []struct {
		name string
		data *v.WeatherData
		want float32
	}{
		// TODO: Add test cases.
		{
			name: "return float on a correct call to the api (imperial)",
			data: RunTestServer("imperial"),
			want: 3.42,
		},
		{
			name: "return float on a correct call to the api (metric)",
			data: RunTestServer("metric"),
			want: 1.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.GetWindSpeed()
			if got != tt.want {
				t.Errorf("GetMinTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetCloudCoverage(t *testing.T) {
	tests := []struct {
		name string
		data *v.WeatherData
		want int
	}{
		// TODO: Add test cases.
		{
			name: "GetCloudCoverage returns integer on a correct call to the api (imperial)",
			data: RunTestServer("imperial"),
			want: 1,
		},
		{
			name: "GetCloudCoverage returns integer on a correct call to the api (metric)",
			data: RunTestServer("metric"),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.GetCloudCoverage()
			if got != tt.want {
				t.Errorf("GetMaxTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}
