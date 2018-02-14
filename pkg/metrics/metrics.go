package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// RequestLatency is a Prometheus metrics that records the latency of each request.
	RequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "tester",
		Name:      "http_request_latency_milliseconds",
		Help:      "Latency of HTTP requests.",
	}, []string{"code", "method"})
	// RequestInFlightCount is a Prometheus metrics that counts the number of requests
	// in flight.
	RequestInFlightCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "tester",
		Name:      "http_request_inflight_count",
		Help:      "Count of all HTTP requests still in flight.",
	})
	// LoggingTotal counts the total number of logs
	LoggingTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "tester",
		Name:      "log_count",
		Help:      "Number of level logs.",
	}, []string{"level"})
)

func init() {
	prometheus.MustRegister(
		RequestLatency,
		RequestInFlightCount,
		LoggingTotal,
	)
}
