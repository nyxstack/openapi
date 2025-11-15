package openapi

// Callback represents a callback in OpenAPI
type Callback map[string]PathItem

// NewCallback creates a new callback
func NewCallback() Callback {
	return make(Callback)
}

// WithPath adds a path to the callback
func (c Callback) WithPath(expression string, pathItem PathItem) Callback {
	c[expression] = pathItem
	return c
}
