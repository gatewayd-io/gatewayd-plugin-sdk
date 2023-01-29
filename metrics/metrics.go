package metrics

import (
	"net"
	"net/http"
	"os"

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

func ExposeMetrics(metricsConfig MetricsConfig, logger hclog.Logger) {
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
