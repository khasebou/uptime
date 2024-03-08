package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

// Metric represents a single data point
type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// MetricsStore (in-memory for this example)
var metrics []Metric

// zerolog logger instance
var logger zerolog.Logger

func init() {
    // Create a new zerolog logger with desired output (in this case, console)
    logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}


// receiveMetrics handler for POST requests
func receiveMetrics(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
    defer func() {
        logger.Info().
            Str("method", r.Method).
            Str("uri", r.RequestURI).
            Float64("duration", time.Since(start).Seconds()).
            Msg("Received metrics request")
    }()

	// Decode the request body into a Metric struct
	var metric Metric
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&metric)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		logger.Error().Err(err).
			Msgf("Error decoding request body: %v", err)
		return
	}

	// Add the metric to the in-memory store (replace with real storage)
	metrics = append(metrics, metric)

	// Respond with success
	w.WriteHeader(http.StatusOK)
	
	logger.Info().
		Str("metric_name", metric.Name).
		Float64("metric_value", metric.Value).
		Msg("Metric received successfully")
}

func main() {
	port := ":8080"
	router := mux.NewRouter()
	router.HandleFunc("/metrics", receiveMetrics).Methods("POST") // Specify POST method

	fmt.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}
