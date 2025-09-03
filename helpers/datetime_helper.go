// helpers/helpers.go
// Package helpers provides utility functions for the API Contact Form application.
//
// It includes functions for time formatting and timezone management based on
// environment configurations.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package helpers

import (
	"api-contact-form/config"
	"log"
	"time"
)

var (
	// appTimezone holds the application's configured timezone.
	appTimezone *time.Location
)

// init initializes the application's timezone based on the environment variable.
// It loads the timezone location and logs a fatal error if the timezone is invalid.
func init() {
	timezoneStr := config.GetEnv("APP_TIMEZONE", "Asia/Jakarta")
	var err error
	appTimezone, err = time.LoadLocation(timezoneStr)
	if err != nil {
		log.Fatalf("Failed to load timezone '%s': %v", timezoneStr, err)
	}
}

// FormatTimeHuman converts a time.Time object to a human-readable string
// in the configured timezone.
//
// Parameters:
//   - t: The time.Time object to format.
//
// Returns:
//   - A string representing the formatted time.
func FormatTimeHuman(t time.Time) string {
	return t.In(appTimezone).Format("2006-01-02 15:04:05")
}