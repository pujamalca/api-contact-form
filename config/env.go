// Package config provides utilities for managing configuration settings.
//
// It includes functions to retrieve environment variables with default fallback values,
// ensuring that the application can gracefully handle missing or unset environment variables.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package config

import "os"

// GetEnv retrieves the value of the environment variable named by the key.
// If the environment variable is not set or is empty, it returns the provided default value.
//
// Parameters:
//   - key: The name of the environment variable to retrieve.
//   - defaultVal: The default value to return if the environment variable is not set.
//
// Returns:
//   - A string containing the value of the environment variable or the default value.
func GetEnv(key, defaultVal string) string {
	// Look up the environment variable by key.
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultVal
}