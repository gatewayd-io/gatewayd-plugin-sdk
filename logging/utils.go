package logging

import (
	"github.com/hashicorp/go-hclog"
	"github.com/rs/zerolog"
)

// GetLogLevel returns the hclog level based on the string passed in.
func GetLogLevel(logLevel string) hclog.Level {
	switch logLevel {
	case "trace":
		return hclog.Trace
	case "debug":
		return hclog.Debug
	case "info":
		return hclog.Info
	case "warn":
		return hclog.Warn
	case "error":
		return hclog.Error
	case "off":
		return hclog.Off
	default:
		return hclog.NoLevel
	}
}

// GetZeroLogLevel returns the zerolog level based on the string passed in.
func GetZeroLogLevel(logLevel string) zerolog.Level {
	switch logLevel {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	case "disabled":
		return zerolog.Disabled
	default:
		return zerolog.NoLevel
	}
}
