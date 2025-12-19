package config

import (
	"time"
)

type (
	Server interface {
		Address() string
		ReadTimeout() time.Duration
		ReadHeaderTimeout() time.Duration
		WriteTimeout() time.Duration
		IdleTimeout() time.Duration
		ShutdownTimeout() time.Duration
	}
)
