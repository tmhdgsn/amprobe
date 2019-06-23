package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	AlertsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "alerts_processed_total",
		Help: "The total number of processed alerts",
	})
)
