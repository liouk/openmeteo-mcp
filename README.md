# OpenMeteo MCP Server

A Model Context Protocol (MCP) server that provides access to weather data through the OpenMeteo API. This server enables AI assistants to retrieve current weather forecasts and geocoding information for any location worldwide.

## Features

- Get current weather forecasts for any location
- Convert location names to latitude/longitude coordinates
- Built with Go using the MCP protocol
- No API key required (uses free OpenMeteo API)

## Tools

This MCP server exposes the following tools:

### `get_current_weather`
Retrieves current weather forecast data from OpenMeteo for a specified location.

**Parameters:**
- `lat` (number, required): Latitude of the location
- `lon` (number, required): Longitude of the location
- `location` (string, optional): Location name as alternative to lat/lon coordinates

**Description:** Returns comprehensive weather forecast data including temperature, humidity, wind speed, precipitation, and other meteorological information. If only a location name is provided, the tool will automatically geocode it to obtain coordinates.

### `get_latlon`
Converts a location name to latitude and longitude coordinates using OpenMeteo's geocoding API.

**Parameters:**
- `location` (string, required): Location name (e.g., "New York", "London, UK", "Tokyo, Japan")

**Description:** Returns geocoding results with latitude and longitude coordinates for the specified location, along with additional location details.

## Usage

This MCP server communicates over standard input/output and is designed to be used with MCP-compatible AI assistants and applications. The server will automatically handle requests for weather data and geocoding operations.

## API

The server uses the OpenMeteo API (https://open-meteo.com/), which provides free access to weather data without requiring an API key.