package common_prometheus_exporter

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metrics = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "cms_gap",
		Help: "Node Count",
	})
)

func recordMetrics() {
	for {
		metrics.Set(14.0)
		time.Sleep(1 * time.Minute)
	}
}
func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
