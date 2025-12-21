package config

import (
	"time"
)

const (
	EnvServerAddress     = "SERVER_ADDRESS"
	DefaultServerAddress = "localhost:8080"
)

const (
	EnvServerReadTimeout     = "SERVER_READ_TIMEOUT"
	DefaultServerReadTimeout = 5 * time.Second
)

const (
	EnvServerReadHeaderTimeout     = "SERVER_READ_HEADER_TIMEOUT"
	DefaultServerReadHeaderTimeout = 5 * time.Second
)

const (
	EnvServerWriteTimeout     = "SERVER_WRITE_TIMEOUT"
	DefaultServerWriteTimeout = 5 * time.Second
)

const (
	EnvServerIdleTimeout     = "SERVER_IDLE_TIMEOUT"
	DefaultServerIdleTimeout = 5 * time.Second
)

const (
	EnvServerShutdownTimeout     = "SERVER_SHUTDOWN_TIMEOUT"
	DefaultServerShutdownTimeout = 5 * time.Second
)

type (
	// Server defines the immutable application configuration for the server.
	Server interface {
		// Address returns the configured TCP address for the server to listen on, in the
		// form of "host:port".
		Address() string

		// ReadTimeout returns the configured maximum duration for reading the entire
		// request, including the body.
		ReadTimeout() time.Duration

		// ReadHeaderTimeout returns the configured amount of time allowed to read request
		// headers.
		ReadHeaderTimeout() time.Duration

		// WriteTimeout returns the configured maximum duration before timing out writes of
		// the response.
		WriteTimeout() time.Duration

		// IdleTimeout returns the configured maximum amount of time to wait for the next
		// request when keep-alives are enabled.
		IdleTimeout() time.Duration

		// ShutdownTimeout returns the configured grace period allowed for active
		// connections to close during a graceful shutdown.
		ShutdownTimeout() time.Duration
	}
)
