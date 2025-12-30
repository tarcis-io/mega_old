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

//

package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type loader struct {
	lookup func(string) (string, bool)
	errs   []error
}

// fail records an error. 
// Defensive: It works even if 'l' or 'l.errs' is nil/uninitialized.
func (l *loader) fail(format string, args ...any) {
	if l == nil {
		return // Or panic, depending on how defensive you want to be about developer error
	}
	l.errs = append(l.errs, fmt.Errorf(format, args...))
}

// get is the defensive foundation.
// 1. Checks for nil receiver.
// 2. Checks for nil lookup function.
// 3. Checks for empty keys.
func (l *loader) get(key string) (string, bool) {
	if l == nil || l.lookup == nil {
		return "", false
	}
	if key == "" {
		l.fail("coding error: attempted to load empty configuration key")
		return "", false
	}
	
	val, ok := l.lookup(key)
	if !ok {
		return "", false
	}
	return strings.TrimSpace(val), true
}

// mustStr (Required String)
func (l *loader) mustStr(key string) string {
	val, ok := l.get(key)
	if !ok {
		l.fail("missing required key: %s", key)
		return ""
	}
	return val
}

// str (Optional String)
func (l *loader) str(key, def string) string {
	val, ok := l.get(key)
	if !ok {
		return def
	}
	return val
}

// int (Optional Int, but STRICT parsing)
// Defensive Rule: If the variable IS present, it MUST be valid. 
// Do not silently fallback to 'def' if the user provided garbage.
func (l *loader) int(key string, def int) int {
	valStr, ok := l.get(key)
	if !ok {
		return def
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		// Defensive: Report the specific parsing error.
		// We return 'def' to satisfy the type, but l.err() will be non-nil.
		l.fail("env var %s='%s' is not a valid int: %v", key, valStr, err)
		return def
	}
	return val
}

// bool (Optional Bool, but STRICT parsing)
func (l *loader) bool(key string, def bool) bool {
	valStr, ok := l.get(key)
	if !ok {
		return def
	}

	b, err := strconv.ParseBool(valStr)
	if err != nil {
		l.fail("env var %s='%s' is not a valid boolean: %v", key, valStr, err)
		return def
	}
	return b
}

// err returns the accumulated errors
func (l *loader) err() error {
	if l == nil || len(l.errs) == 0 {
		return nil
	}
	return errors.Join(l.errs...)
}
