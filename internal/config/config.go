// Package config loads and provides the application configuration.
package config

type (
	Config interface {
		Log() Log
		Server() Server
	}
)
