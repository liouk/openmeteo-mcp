package tools

import (
	"context"

	"github.com/liouk/openmeteo-mcp/pkg/openmeteo"
	mcplib "github.com/mark3labs/mcp-go/mcp"
)

func GetCurrentWeather(ctx context.Context, request mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	lat, err := request.RequireFloat("lat")
	if err != nil {
		return mcplib.NewToolResultError(err.Error()), nil
	}

	lon, err := request.RequireFloat("lon")
	if err != nil {
		return mcplib.NewToolResultError(err.Error()), nil
	}

	data, err := openmeteo.Forecast(lat, lon)
	if err != nil {
		return mcplib.NewToolResultError(err.Error()), nil
	}

	// return response data unprocessed, let the model do the work
	return mcplib.NewToolResultText(string(data)), nil
}
