package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Metric struct {
	CPUPercent    float64 `json:"cpu_percent"`
	MemoryPercent float64 `json:"memory_percent"`
	DiskPercent   float64 `json:"disk_percent"`
	Timestamp     float64 `json:"timestamp"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "ok",
		"message": "Go API is running",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var metric Metric

	err := json.NewDecoder(r.Body).Decode(&metric)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Metric received:")
	fmt.Println("CPU:", metric.CPUPercent)
	fmt.Println("Memory:", metric.MemoryPercent)
	fmt.Println("Disk:", metric.DiskPercent)
	fmt.Println("Timestamp:", metric.Timestamp)
	fmt.Println("------------------------------")

	response := map[string]string{
		"status":  "ok",
		"message": "Metric received successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/metrics", metricsHandler)

	port := "8000"

	fmt.Println("Go API running on port", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
