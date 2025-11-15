package openapi

// Header represents a header in OpenAPI
type Header struct {
	Description     string               `json:"description,omitempty"`
	Required        bool                 `json:"required,omitempty"`
	Deprecated      bool                 `json:"deprecated,omitempty"`
	AllowEmptyValue bool                 `json:"allowEmptyValue,omitempty"`
	Style           string               `json:"style,omitempty"`
	Explode         bool                 `json:"explode,omitempty"`
	AllowReserved   bool                 `json:"allowReserved,omitempty"`
	Schema          *Schema              `json:"schema,omitempty"`
	Example         interface{}          `json:"example,omitempty"`
	Examples        map[string]Example   `json:"examples,omitempty"`
	Content         map[string]MediaType `json:"content,omitempty"`
}

// NewHeader creates a new header
func NewHeader() Header {
	return Header{
		Examples: make(map[string]Example),
		Content:  make(map[string]MediaType),
	}
}

// WithDescription sets the description
func (h Header) WithDescription(description string) Header {
	h.Description = description
	return h
}

// WithRequired sets whether the header is required
func (h Header) WithRequired(required bool) Header {
	h.Required = required
	return h
}

// WithDeprecated marks the header as deprecated
func (h Header) WithDeprecated(deprecated bool) Header {
	h.Deprecated = deprecated
	return h
}

// WithAllowEmptyValue sets whether empty values are allowed
func (h Header) WithAllowEmptyValue(allow bool) Header {
	h.AllowEmptyValue = allow
	return h
}

// WithStyle sets the header style
func (h Header) WithStyle(style string) Header {
	h.Style = style
	return h
}

// WithExplode sets the explode option
func (h Header) WithExplode(explode bool) Header {
	h.Explode = explode
	return h
}

// WithAllowReserved sets whether reserved characters are allowed
func (h Header) WithAllowReserved(allow bool) Header {
	h.AllowReserved = allow
	return h
}

// WithSchema sets the schema for the header
func (h Header) WithSchema(schema *Schema) Header {
	h.Schema = schema
	return h
}

// WithExample sets an example for the header
func (h Header) WithExample(example interface{}) Header {
	h.Example = example
	return h
}

// WithExamples adds named examples
func (h Header) WithExamples(name string, example Example) Header {
	if h.Examples == nil {
		h.Examples = make(map[string]Example)
	}
	h.Examples[name] = example
	return h
}
