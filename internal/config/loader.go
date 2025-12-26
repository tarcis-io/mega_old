package config

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type (
	provider func(string) (string, bool)

	loader struct {
		provider provider
		errs     []error
	}
)

func (l *loader) positiveDuration(envKey string, fallback time.Duration) time.Duration {
	d := l.duration(envKey, fallback)
	if d <= 0 {
		l.addErrorf("invalid configuration: env=%q value=%q err=\"duration must be positive\"", envKey, d.String())
		return fallback
	}
	return d
}

func (l *loader) nonNegativeDuration(envKey string, fallback time.Duration) time.Duration {
	d := l.duration(envKey, fallback)
	if d < 0 {
		l.addErrorf("invalid configuration: env=%q value=%q err=\"duration must be non-negative\"", envKey, d.String())
		return fallback
	}
	return d
}

func (l *loader) duration(envKey string, fallback time.Duration) time.Duration {
	s, ok := l.get(envKey)
	if !ok || s == "" {
		return fallback
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		l.addErrorf("invalid configuration: env=%q value=%q err=\"%w\"", envKey, s, err)
		return fallback
	}
	return d
}

func (l *loader) get(key string) (string, bool) {
	if l.provider == nil {
		return "", false
	}
	val, ok := l.provider(key)
	if !ok {
		return "", false
	}
	return strings.TrimSpace(val), true
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
	s, ok := l.get(envKey)
	if !ok || s == "" {
		return fallback
	}
	if val, ok := match(s, allowed...); ok {
		return val
	}
	allowedStr := make([]string, len(allowed))
	for i, a := range allowed {
		allowedStr[i] = string(a)
	}
	l.addErrorf("invalid configuration: env=%q value=%q allowed=[%s]", envKey, s, strings.Join(allowedStr, ", "))
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
