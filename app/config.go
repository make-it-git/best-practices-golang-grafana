package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

// getEnv returns environment variable value or fallback if empty.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

// mustParseDuration parses duration or exits the app.
// Fail-fast is correct behavior for configuration errors.
func mustParseDuration(value string) time.Duration {
	d, err := time.ParseDuration(value)
	if err != nil {
		log.Fatalf("invalid duration value %q: %v", value, err)
	}
	return d
}

// mustParsePercentage parses 0-1 style value or exits the app.
// Fail-fast is correct behavior for configuration errors.
func mustParsePercentage(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("invalid percentage (0-1) value %q: %v", value, err)
	}
	if v < 0 || v > 1 {
		log.Fatalf("invalid percentage (0-1) value %q: %v", value, err)
	}
	return v
}
