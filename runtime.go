package wasmrt

// Runtime ...
type Runtime struct{}

// New ...
func New() *Runtime {
	return new(Runtime)
}

func NewWithConfig(c Config) (*Runtime, error) {
	panic("not implemented")
}
