package engine

// Option is configurable Engine behaver.
type Option func(*opts)

// IgnoreEmpty ignores empty value from external store.
func IgnoreEmpty() Option {
	return func(o *opts) {
		o.IgnoreEmpty = true
	}
}

type opts struct {
	IgnoreEmpty bool
}

var defaultOpts = &opts{
	IgnoreEmpty: false,
}
