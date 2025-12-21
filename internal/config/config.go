// Package config loads and provides the application configuration.
package config

type (
	// Config defines the immutable application configuration.
	Config interface {
		// Log returns the application configuration for logging.
		Log() Log

		// Server returns the application configuration for the server.
		Server() Server
	}
)
