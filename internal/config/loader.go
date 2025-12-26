package config

import (
	"errors"
	"fmt"
	"slices"
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

func (l *loader) positiveDuration(envKey string, fallback time.Duration) time.Duration {
	d := l.duration(envKey, fallback)
	if d <= 0 {
		l.addErrorf("invalid configuration: env=%q got=%q err=\"duration must be positive\"", envKey, d.String())
		return fallback
	}
	return d
}

func (l *loader) nonNegativeDuration(envKey string, fallback time.Duration) time.Duration {
	d := l.duration(envKey, fallback)
	if d < 0 {
		l.addErrorf("invalid configuration: env=%q got=%q err=\"duration must be non-negative\"", envKey, d.String())
		return fallback
	}
	return d
}

func (l *loader) duration(envKey string, fallback time.Duration) time.Duration {
	s := l.env(envKey)
	if s == "" {
		return fallback
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		l.addErrorf("invalid configuration: env=%q got=%q err=\"%w\"", envKey, s, err)
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

func (l *loader) Err() error {
	return errors.Join(l.errs...)
}

func oneOf[T ~string](l *loader, envKey string, fallback T, allowed ...T) T {
	s := l.env(envKey)
	if s == "" {
		return fallback
	}
	if val, ok := match(s, allowed...); ok {
		return val
	}
	allowedStr := make([]string, len(allowed))
	for i, a := range allowed {
		allowedStr[i] = string(a)
	}
	l.addErrorf("invalid configuration: env=%q got=%q allowed=[%s]", envKey, s, strings.Join(allowedStr, ", "))
	return fallback
}

func match[T ~string](s string, allowed ...T) (T, bool) {
	idx := slices.IndexFunc(allowed, func(a T) bool {
		return strings.EqualFold(s, string(a))
	})
	if idx >= 0 {
		return allowed[idx], true
	}
	var zero T
	return zero, false
}
