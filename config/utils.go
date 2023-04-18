package config

import "os"

// GetEnv is a helper function that returns the value of an environment variable.
// If the environment variable is not set, it returns the default value as fallback.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
