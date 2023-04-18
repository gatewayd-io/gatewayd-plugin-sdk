package metrics

import (
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/hashicorp/go-hclog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	Namespace = "gatewayd"
)

type MetricsConfig struct {
	Enabled          bool
	UnixDomainSocket string
	Endpoint         string
}

// NewMetricsConfig returns a new MetricsConfig struct.
func NewMetricsConfig(config map[string]interface{}) *MetricsConfig {
	metricsConfig := MetricsConfig{}
	if enabled, ok := config["metricsEnabled"].(string); ok {
		if enabled, err := strconv.ParseBool(enabled); err == nil {
			metricsConfig.Enabled = enabled
		} else {
			return nil
		}
	}
	if uds, ok := config["metricsUnixDomainSocket"].(string); ok {
		metricsConfig.UnixDomainSocket = uds
	}
	if endpoint, ok := config["metricsEndpoint"].(string); ok {
		metricsConfig.Endpoint = endpoint
	}

	return &metricsConfig
}

// ExposeMetrics exposes the Prometheus metrics via HTTP over Unix domain socket.
func ExposeMetrics(metricsConfig *MetricsConfig, logger hclog.Logger) {
	logger.Info(
		"Starting metrics server via HTTP over Unix domain socket",
		"unixDomainSocket", metricsConfig.UnixDomainSocket,
		"endpoint", metricsConfig.Endpoint)

	if file, err := os.Stat(metricsConfig.UnixDomainSocket); err == nil && !file.IsDir() && file.Mode().Type() == os.ModeSocket {
		if err := os.Remove(metricsConfig.UnixDomainSocket); err != nil {
			logger.Error("Failed to remove unix domain socket")
		}
	}

	listener, err := net.Listen("unix", metricsConfig.UnixDomainSocket)
	if err != nil {
		logger.Error("Failed to start metrics server")
	}

	if err := http.Serve(listener, promhttp.Handler()); err != nil {
		logger.Error("Failed to start metrics server")
	}
}
