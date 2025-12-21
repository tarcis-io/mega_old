// Package config loads and provides the application configuration.
package config

type (
	// Config defines the immutable application configuration.
	Config interface {
		// Log returns the application configuration for [Log].
		Log() Log

		// Server returns the application configuration for [Server].
		Server() Server
	}
)
