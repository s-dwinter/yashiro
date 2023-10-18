package engine

type EngineOption func(*engineOpts)

// Ignore Empty ignores empty value from external store.
func IgnoreEmpty() EngineOption {
	return func(co *engineOpts) {
		co.IgnoreEmpty = true
	}
}

type engineOpts struct {
	IgnoreEmpty bool
}

var defaultEngineOpts = &engineOpts{
	IgnoreEmpty: false,
}
