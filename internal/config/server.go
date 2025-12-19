package config

import (
	"time"
)

type (
	// Server defines the immutable application configuration for the server.
	Server interface {
		// Address returns the configured server address.
		Address() string

		// ReadTimeout returns the configured server read timeout.
		ReadTimeout() time.Duration

		// ReadHeaderTimeout returns the configured server read header timeout.
		ReadHeaderTimeout() time.Duration

		// WriteTimeout returns the configured server write timeout.
		WriteTimeout() time.Duration

		// IdleTimeout returns the configured server idle timeout.
		IdleTimeout() time.Duration

		// ShutdownTimeout returns the configured server shutdown timeout.
		ShutdownTimeout() time.Duration
	}
)
