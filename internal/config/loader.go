package config

import (
	"errors"
	"fmt"
	"strings"
)

type (
	lookupFunc func(string) string

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

func (l *loader) env(key string) string {
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
