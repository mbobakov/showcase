// Package metric is resposible for the all metric related queries
package metric

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mbobakov/showcase/restapi/operations/metrics"
	"github.com/mbobakov/showcase/storage"
)

type provider interface {
	Summary() (map[string]int64, error)
	Store(*storage.MetricDTO) error
}

// Service for the metric operations
type Service struct {
	prov provider
}

// New metric Service
func New(p provider) *Service {
	// TODO: Add check for p != nil
	return &Service{prov: p}
}

// FindMetrics implements api.MetricsFindMetricsHandler
func (s *Service) FindMetrics(params metrics.FindMetricsParams) middleware.Responder {
	sum, err := s.prov.Summary()
	if err != nil {
		return metrics.NewFindMetricsBadRequest()
	}

	res := make(map[string]metrics.FindMetricsOKBodyAnon, len(sum))
	for k, v := range sum {
		res[k] = metrics.FindMetricsOKBodyAnon{Datapoints: v}
	}
	return metrics.NewFindMetricsOK().WithPayload(res)
}

// PostDatapoint implements api.MetricsPostDatapointHandler
func (s *Service) PostDatapoint(params metrics.PostDatapointParams) middleware.Responder {

	err := s.prov.Store(
		&storage.MetricDTO{
			Name:       params.MetricName,
			DataPoints: map[int64]int64{*params.Body.TimestampUtc: *params.Body.Value},
		},
	)
	if err != nil {
		return metrics.NewPostDatapointInternalServerError()
	}

	return metrics.NewPostDatapointCreated()
}
