.PHONY: build mcp-list-tools mcp-get-current-weather mcp-get-latlon

build:
	go build -o openmeteo-mcp ./cmd

mcp-list-tools: build
	echo '{"jsonrpc": "2.0", "id": 1, "method": "tools/list", "params": {}}' | ./openmeteo-mcp | jq -r '.result.tools[] | "🔧 \(.name)\n   \(.description)\n   Parameters: \([.inputSchema.required[] as $$key | "\($$key) (\(.inputSchema.properties[$$key].description))"] | join(", "))\n"'

mcp-tool-get-current-weather: build
	echo '{"jsonrpc": "2.0", "id": 1, "method": "tools/call", "params": {"name": "get_current_weather", "arguments": {"lat": 47.36667, "lon": 8.55}}}' | ./openmeteo-mcp | jq

mcp-tool-get-latlon: build
	echo '{"jsonrpc": "2.0", "id": 1, "method": "tools/call", "params": {"name": "get_latlon", "arguments": {"location": "Zurich"}}}' | ./openmeteo-mcp | jq