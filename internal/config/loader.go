package config

type (
	loader struct {
		errs []error
	}
)

func newLoader() *loader {
	return &loader{}
}
