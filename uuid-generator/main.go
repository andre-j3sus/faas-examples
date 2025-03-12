package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type UUIDResponse struct {
	UUID      string    `json:"uuid"`
	Version   int       `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

type BatchRequest struct {
	Count int `json:"count"`
}

type BatchResponse struct {
	UUIDs     []string  `json:"uuids"`
	Count     int       `json:"count"`
	Timestamp time.Time `json:"timestamp"`
}

func handleGenerateUUID(w http.ResponseWriter, r *http.Request) {
	// Generate a new UUID v4
	id := uuid.New()
	
	response := UUIDResponse{
		UUID:      id.String(),
		Version:   4,
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleGenerateBatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req BatchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	count := req.Count
	if count <= 0 {
		count = 1 // Default to 1 if not specified or invalid
	}
	if count > 100 {
		count = 100 // Cap at 100 to prevent abuse
	}

	// Generate requested number of UUIDs
	uuids := make([]string, count)
	for i := 0; i < count; i++ {
		uuids[i] = uuid.New().String()
	}

	response := BatchResponse{
		UUIDs:     uuids,
		Count:     count,
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/uuid", handleGenerateUUID)
	http.HandleFunc("/batch", handleGenerateBatch)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("UUID Generator function listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}