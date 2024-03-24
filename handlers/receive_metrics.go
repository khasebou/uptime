package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/khasebou/uptime/structures"
	"github.com/khasebou/uptime/logging"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	logger = logging.DefaultLogger()
}


// MetricsStore (in-memory for this example)
var metrics []structures.Metric

// receiveMetrics handler for POST requests
func ReceiveMetrics(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		logger.Info().
			Str("method", r.Method).
			Str("uri", r.RequestURI).
			Float64("duration", time.Since(start).Seconds()).
			Msg("Received metrics request")
	}()

	// Decode the request body into a Metric struct
	var metric structures.Metric
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
