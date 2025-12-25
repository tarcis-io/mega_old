package config

import (
	"time"
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
