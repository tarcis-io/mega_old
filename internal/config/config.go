// Package config loads and provides the application configuration.
package config

type (
	// Config defines the immutable application configuration.
	Config interface {
		// Log returns the application configuration for the logging system, including the
		// verbosity, encoding format, and destination stream.
		Log() Log

		// Server returns the application configuration for the HTTP server listener,
		// including the binding address and connection timeouts.
		Server() Server
	}
)
