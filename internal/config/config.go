package config

type (
	Config interface {
		Log() Log
		Server() Server
	}
)
