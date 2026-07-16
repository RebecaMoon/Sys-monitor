package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) { // This function handles requests sent to /. w is the response sent by the server, and r contains the data from the request.
	response := map[string]string{ // Create the response.
		"status":  "ok",
		"message": "Go API is running",
	}

	w.Header().Set("Content-Type", "application/json") // Set the response content type.
	json.NewEncoder(w).Encode(response)                // Encode the response as JSON.
}

func metricsHandler(w http.ResponseWriter, r *http.Request) { // This is the most important part. It handles requests sent to /metrics.
	if r.Method != http.MethodPost { // First, it validates the HTTP method. It only allows POST.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var metric Metric // Create the metric variable. It is empty at this point.

	err := json.NewDecoder(r.Body).Decode(&metric) // This line gets the JSON body and saves the data into the metric variable.
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Metric received:") // Print the data just to check if it is working.
	fmt.Println("CPU:", metric.CPUPercent)
	fmt.Println("Memory:", metric.MemoryPercent)
	fmt.Println("Disk:", metric.DiskPercent)
	fmt.Println("Timestamp:", metric.Timestamp)
	fmt.Println("------------------------------")

	err = SaveMetric(r.Context(), metric)
	if err != nil {
		log.Println("Failed to save metric:", err)
		http.Error(w, "Failed to save metric", http.StatusInternalServerError)
		return
	}

	response := map[string]string{ // Create the success response.
		"status":  "ok",
		"message": "Metric saved successfully",
	}

	w.Header().Set("Content-Type", "application/json") // Set the response content type.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response) // Encode the response as JSON.
}

func latestMetricHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	metric, err := GetLatestMetric(r.Context())
	if err != nil {
		log.Println("Failed to retrieve latest metric:", err)

		response := map[string]string{
			"detail": "No metrics available",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metric)
}

func metricsHistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	metrics, err := GetMetricsHistory(r.Context())
	if err != nil {
		log.Println("Failed to retrieve metrics history:", err)
		http.Error(w, "Failed to retrieve metrics history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func main() { // This is the main function.
	err := ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer DB.Close()

	err = CreateTables(context.Background())
	if err != nil {
		log.Fatal("Table creation error:", err)
	}

	http.HandleFunc("/", healthHandler) // These two lines define the endpoints.
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/metrics/latest", latestMetricHandler)
	http.HandleFunc("/metrics/history", metricsHistoryHandler)

	port := "8000" // This is the server port.

	fmt.Println("Go API running on port", port)

	err = http.ListenAndServe(":"+port, nil) // This line starts the server.
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
