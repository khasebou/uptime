package structures

// Metric represents a single data point
type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}