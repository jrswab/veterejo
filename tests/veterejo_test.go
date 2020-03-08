package veterejo

import (
	"testing"

	v "github.com/jrswab/veterejo"
)

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
