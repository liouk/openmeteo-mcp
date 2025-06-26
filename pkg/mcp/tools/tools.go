package tools

import (
	"context"
	"encoding/json"

	"github.com/liouk/openmeteo-mcp/pkg/openmeteo"
	mcplib "github.com/mark3labs/mcp-go/mcp"
)

func GetCurrentWeather(ctx context.Context, request mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	var lat, lon float64
	var err error

	lat, latErr := request.RequireFloat("lat")
	lon, lonErr := request.RequireFloat("lon")

	if latErr != nil || lonErr != nil {
		location, locErr := request.RequireString("location")
		if locErr != nil {
			return mcplib.NewToolResultError("Either 'lat' and 'lon' parameters or 'location' parameter must be provided"), nil
		}

		geocodeData, err := openmeteo.Geocoding(location)
		if err != nil {
			return mcplib.NewToolResultError(err.Error()), nil
		}

		var geocodeResult struct {
			Results []struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"results"`
		}

		if err := json.Unmarshal(geocodeData, &geocodeResult); err != nil {
			return mcplib.NewToolResultError("Failed to parse geocoding response"), nil
		}

		if len(geocodeResult.Results) == 0 {
			return mcplib.NewToolResultError("No location found for the given query"), nil
		}

		lat = geocodeResult.Results[0].Latitude
		lon = geocodeResult.Results[0].Longitude
	}

	data, err := openmeteo.Forecast(lat, lon)
	if err != nil {
		return mcplib.NewToolResultError(err.Error()), nil
	}

	// return response data unprocessed, let the model do the work
	return mcplib.NewToolResultText(string(data)), nil
}

func GetLatLon(ctx context.Context, request mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	location, err := request.RequireString("location")
	if err != nil {
		return mcplib.NewToolResultError(err.Error()), nil
	}

	data, err := openmeteo.Geocoding(location)
	if err != nil {
		return mcplib.NewToolResultError(err.Error()), nil
	}

	return mcplib.NewToolResultText(string(data)), nil
}
