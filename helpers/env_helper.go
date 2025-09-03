// helpers/helpers.go
// Package helpers provides utility functions for the API Contact Form application.
//
// It includes functions for parsing environment variables and handling configuration-related tasks.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package helpers

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// ParseEnvList parses a comma-separated environment variable into a slice of strings.
// It trims any whitespace around the elements.
//
// Parameters:
//   - key: The name of the environment variable to parse.
//
// Returns:
//   - A slice of strings containing the parsed values, or an empty slice if the variable is not set or empty.
func ParseEnvList(key string) []string {
	val, exists := os.LookupEnv(key)
	if !exists || val == "" {
		return []string{}
	}
	parts := strings.Split(val, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// GetEnvBool retrieves a boolean environment variable.
// It returns the defaultValue if the environment variable is not set or cannot be parsed.
//
// Parameters:
//   - key: The name of the environment variable to retrieve.
//   - defaultValue: The default boolean value to return if the variable is not set or invalid.
//
// Returns:
//   - A boolean value representing the environment variable's value or the default value.
func GetEnvBool(key string, defaultValue bool) bool {
	val, exists := os.LookupEnv(key)
	if !exists || val == "" {
		return defaultValue
	}
	parsedVal, err := strconv.ParseBool(strings.ToLower(val))
	if err != nil {
		log.Printf("Warning: Could not parse boolean value for %s: %v. Using default: %v", key, err, defaultValue)
		return defaultValue
	}
	return parsedVal
}