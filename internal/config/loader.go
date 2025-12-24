package config

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type (
	lookupFunc func(string) string

	// loader is a helper for loading environment variables into application
	// configuration.
	//
	// It aggregates errors internally so that all loading issues can be reported at
	// once, rather than failing on the first error.
	loader struct {
		lookup lookupFunc
		errs   []error
	}
)

func newLoader(lookup lookupFunc) *loader {
	return &loader{
		lookup: lookup,
	}
}

func (l *loader) duration(envKey string, fallback time.Duration) time.Duration {
	s := l.env(envKey)
	if s == "" {
		return fallback
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		l.addErrorf("invalid configuration (%s) got=%q: %w", envKey, s, err)
		return fallback
	}
	return d
}

func (l *loader) env(key string) string {
	if l.lookup == nil {
		return ""
	}
	return strings.TrimSpace(l.lookup(key))
}

func (l *loader) addError(err error) {
	if err == nil {
		return
	}
	l.errs = append(l.errs, err)
}

func (l *loader) addErrorf(format string, args ...any) {
	l.addError(fmt.Errorf(format, args...))
}

func (l *loader) err() error {
	return errors.Join(l.errs...)
}

func oneOf[T ~string](l *loader, envKey string, fallback T, allowed ...T) T {
	s := l.env(envKey)
	if s == "" {
		return fallback
	}
	for _, a := range allowed {
		if strings.EqualFold(s, string(a)) {
			return a
		}
	}
	l.addErrorf("invalid configuration (%s) got=%q", envKey, s)
	return fallback
}
