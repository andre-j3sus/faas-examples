# Weather Forecast

A function that retrieves the weather forecast for a given location.

## Build

```bash
docker build -t faas-weather-forecast .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-weather-forecast faas-weather-forecast
```

## Example Call

```bash
# Get weather forecast for a location
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "location": "New York"
  }' \
  http://localhost:8080/forecast

# Expected response:
# {
#   "location": "New York",
#   "temperature": 25.5,
#   "unit": "Celsius",
#   "conditions": "Partly Cloudy",
#   "forecast": [
#     {
#       "day": "Today",
#       "temperature": 25.5,
#       "conditions": "Partly Cloudy"
#     },
#     {
#       "day": "Tomorrow",
#       "temperature": 27.2,
#       "conditions": "Sunny"
#     },
#     {
#       "day": "Day After",
#       "temperature": 22.8,
#       "conditions": "Rain Showers"
#     }
#   ]
# }
```
