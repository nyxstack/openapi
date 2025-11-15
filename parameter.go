package openapi

// Parameter represents a parameter in OpenAPI
type Parameter struct {
	Name            string               `json:"name"`
	In              string               `json:"in"`
	Description     string               `json:"description,omitempty"`
	Required        bool                 `json:"required,omitempty"`
	Deprecated      bool                 `json:"deprecated,omitempty"`
	AllowEmptyValue bool                 `json:"allowEmptyValue,omitempty"`
	Style           string               `json:"style,omitempty"`
	Explode         *bool                `json:"explode,omitempty"`
	AllowReserved   bool                 `json:"allowReserved,omitempty"`
	Schema          *Schema              `json:"schema,omitempty"`
	Example         interface{}          `json:"example,omitempty"`
	Examples        map[string]Example   `json:"examples,omitempty"`
	Content         map[string]MediaType `json:"content,omitempty"`
}

// NewParameter creates a new parameter
func NewParameter(name, in, description string) Parameter {
	return Parameter{
		Name:        name,
		In:          in,
		Description: description,
	}
}

// NewPathParameter creates a new path parameter
func NewPathParameter(name, description string, schema *Schema) Parameter {
	return Parameter{
		Name:        name,
		In:          "path",
		Description: description,
		Required:    true,
		Schema:      schema,
	}
}

// NewQueryParameter creates a new query parameter
func NewQueryParameter(name, description string, required bool, schema *Schema) Parameter {
	return Parameter{
		Name:        name,
		In:          "query",
		Description: description,
		Required:    required,
		Schema:      schema,
	}
}

// NewHeaderParameter creates a new header parameter
func NewHeaderParameter(name, description string, required bool, schema *Schema) Parameter {
	return Parameter{
		Name:        name,
		In:          "header",
		Description: description,
		Required:    required,
		Schema:      schema,
	}
}

// NewCookieParameter creates a new cookie parameter
func NewCookieParameter(name, description string, required bool, schema *Schema) Parameter {
	return Parameter{
		Name:        name,
		In:          "cookie",
		Description: description,
		Required:    required,
		Schema:      schema,
	}
}

// WithRequired sets whether the parameter is required
func (p Parameter) WithRequired(required bool) Parameter {
	p.Required = required
	return p
}

// WithDeprecated marks the parameter as deprecated
func (p Parameter) WithDeprecated(deprecated bool) Parameter {
	p.Deprecated = deprecated
	return p
}

// WithAllowEmptyValue sets whether empty values are allowed
func (p Parameter) WithAllowEmptyValue(allow bool) Parameter {
	p.AllowEmptyValue = allow
	return p
}

// WithStyle sets the parameter style
func (p Parameter) WithStyle(style string) Parameter {
	p.Style = style
	return p
}

// WithExplode sets the explode option
func (p Parameter) WithExplode(explode bool) Parameter {
	p.Explode = &explode
	return p
}

// WithAllowReserved sets whether reserved characters are allowed
func (p Parameter) WithAllowReserved(allow bool) Parameter {
	p.AllowReserved = allow
	return p
}

// WithExample sets an example for the parameter
func (p Parameter) WithExample(example interface{}) Parameter {
	p.Example = example
	return p
}
