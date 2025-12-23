package config

import (
	"errors"
	"fmt"
)

type (
	loader struct {
		errs []error
	}
)

func newLoader() *loader {
	return &loader{}
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
