package openmeteo

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	openmeteoURL = "https://api.open-meteo.com/v1"
	apiForecast  = "/forecast"
)

func Forecast(lat, lon float64) ([]byte, error) {
	p, err := url.JoinPath(openmeteoURL, apiForecast)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(p)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("latitude", fmt.Sprintf("%.6f", lat))
	params.Add("longitude", fmt.Sprintf("%.6f", lon))
	params.Add("current", "temperature_2m,is_day,showers,cloud_cover,wind_speed_10m,wind_direction_10m,pressure_msl,snowfall,precipitation,relative_humidity_2m,apparent_temperature,rain,weather_code,surface_pressure,wind_gusts_10m")

	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "openmeteo-mcp/1.0.0")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %d %s", resp.StatusCode, resp.Status)
	}

	return io.ReadAll(resp.Body)
}
