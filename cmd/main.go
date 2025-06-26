package main

import (
	"fmt"

	"github.com/liouk/openmeteo-mcp/pkg/mcp/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("openmeteo-mcp", "1.0.0", server.WithToolCapabilities(false))

	toolGetCurrentWeather := mcp.NewTool("get_current_weather",
		mcp.WithDescription("Get current weather forecast from OpenMeteo for the specified location"),
		mcp.WithNumber("lat", mcp.Required(), mcp.Description("Latitude of location")),
		mcp.WithNumber("lon", mcp.Required(), mcp.Description("Longitude of location")),
	)

	toolGetLatLon := mcp.NewTool("get_latlon",
		mcp.WithDescription("Convert a location name to latitude and longitude coordinates using OpenMeteo geocoding API"),
		mcp.WithString("location", mcp.Required(), mcp.Description("Location name (e.g., 'New York', 'London, UK')")),
	)

	s.AddTool(toolGetCurrentWeather, tools.GetCurrentWeather)
	s.AddTool(toolGetLatLon, tools.GetLatLon)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("server error: %v\n", err)
	}
}
