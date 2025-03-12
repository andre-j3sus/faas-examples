package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type WeatherRequest struct {
	Location string `json:"location"`
}

type WeatherResponse struct {
	Location    string        `json:"location"`
	Temperature float64       `json:"temperature"`
	Unit        string        `json:"unit"`
	Conditions  string        `json:"conditions"`
	Forecast    []DayForecast `json:"forecast"`
}

type DayForecast struct {
	Day         string  `json:"day"`
	Temperature float64 `json:"temperature"`
	Conditions  string  `json:"conditions"`
}

func handleWeatherForecast(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req WeatherRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if req.Location == "" {
		http.Error(w, "Location is required", http.StatusBadRequest)
		return
	}

	// Note: In a real-world scenario, this would call an actual weather API.
	// This is a mock response for demonstration purposes.
	mockForecast := WeatherResponse{
		Location:    req.Location,
		Temperature: 25.5,
		Unit:        "Celsius",
		Conditions:  "Partly Cloudy",
		Forecast: []DayForecast{
			{
				Day:         "Today",
				Temperature: 25.5,
				Conditions:  "Partly Cloudy",
			},
			{
				Day:         "Tomorrow",
				Temperature: 27.2,
				Conditions:  "Sunny",
			},
			{
				Day:         "Day After",
				Temperature: 22.8,
				Conditions:  "Rain Showers",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockForecast)
}

func main() {
	http.HandleFunc("/forecast", handleWeatherForecast)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Weather Forecast function listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}