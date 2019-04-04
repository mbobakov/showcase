package storage

// MetricDTO is object for communicating
type MetricDTO struct {
	Name       string
	DataPoints map[int64]int64
}
