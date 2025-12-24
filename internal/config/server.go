package config

import (
	"time"
)

const (
	// EnvServerAddress specifies the environment variable name for configuring the
	// server address.
	//
	// Expected format: "<host>:port" (e.g., "localhost:8080", ":3030").
	//
	// Default: [DefaultServerAddress].
	EnvServerAddress = "SERVER_ADDRESS"

	// DefaultServerAddress specifies the default server address, used as the fallback
	// when [EnvServerAddress] is unset or contains an invalid value.
	DefaultServerAddress = "localhost:8080"
)

const (
	// EnvServerReadTimeout specifies the environment variable name for configuring the
	// server read timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m").
	//
	// Default: [DefaultServerReadTimeout].
	EnvServerReadTimeout = "SERVER_READ_TIMEOUT"

	// DefaultServerReadTimeout specifies the default server read timeout, used as the
	// fallback when [EnvServerReadTimeout] is unset or contains an invalid value.
	DefaultServerReadTimeout = 5 * time.Second
)

const (
	// EnvServerReadHeaderTimeout specifies the environment variable name for
	// configuring the server read header timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m").
	//
	// Default: [DefaultServerReadHeaderTimeout].
	EnvServerReadHeaderTimeout = "SERVER_READ_HEADER_TIMEOUT"

	// DefaultServerReadHeaderTimeout specifies the default server read header timeout,
	// used as the fallback when [EnvServerReadHeaderTimeout] is unset or contains an
	// invalid value.
	DefaultServerReadHeaderTimeout = 2 * time.Second
)

const (
	// EnvServerWriteTimeout specifies the environment variable name for configuring
	// the server write timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m").
	//
	// Default: [DefaultServerWriteTimeout].
	EnvServerWriteTimeout = "SERVER_WRITE_TIMEOUT"

	// DefaultServerWriteTimeout specifies the default server write timeout, used as
	// the fallback when [EnvServerWriteTimeout] is unset or contains an invalid value.
	DefaultServerWriteTimeout = 10 * time.Second
)

const (
	// EnvServerIdleTimeout specifies the environment variable name for configuring the
	// server idle timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m").
	//
	// Default: [DefaultServerIdleTimeout].
	EnvServerIdleTimeout = "SERVER_IDLE_TIMEOUT"

	// DefaultServerIdleTimeout specifies the default server idle timeout, used as the
	// fallback when [EnvServerIdleTimeout] is unset or contains an invalid value.
	DefaultServerIdleTimeout = 60 * time.Second
)

const (
	// EnvServerShutdownTimeout specifies the environment variable name for configuring
	// the server shutdown timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m").
	//
	// Default: [DefaultServerShutdownTimeout].
	EnvServerShutdownTimeout = "SERVER_SHUTDOWN_TIMEOUT"

	// DefaultServerShutdownTimeout specifies the default server shutdown timeout, used
	// as the fallback when [EnvServerShutdownTimeout] is unset or contains an invalid
	// value.
	DefaultServerShutdownTimeout = 15 * time.Second
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

type (
	server struct {
		address           string
		readTimeout       time.Duration
		readHeaderTimeout time.Duration
		writeTimeout      time.Duration
		idleTimeout       time.Duration
		shutdownTimeout   time.Duration
	}
)

func (s *server) Address() string {
	return s.address
}

func (s *server) ReadTimeout() time.Duration {
	return s.readTimeout
}

func (s *server) ReadHeaderTimeout() time.Duration {
	return s.readHeaderTimeout
}

func (s *server) WriteTimeout() time.Duration {
	return s.writeTimeout
}

func (s *server) IdleTimeout() time.Duration {
	return s.idleTimeout
}

func (s *server) ShutdownTimeout() time.Duration {
	return s.shutdownTimeout
}
