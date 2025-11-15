package openapi

// Example represents an example in OpenAPI
type Example struct {
	Summary       string                 `json:"summary,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Value         interface{}            `json:"value,omitempty"`
	ExternalValue string                 `json:"externalValue,omitempty"`
	Extensions    map[string]interface{} `json:"-"`
}

// NewExample creates a new example
func NewExample() Example {
	return Example{}
}

// WithSummary sets the summary of the example
func (e Example) WithSummary(summary string) Example {
	e.Summary = summary
	return e
}

// WithDescription sets the description of the example
func (e Example) WithDescription(description string) Example {
	e.Description = description
	return e
}

// WithValue sets the value of the example
func (e Example) WithValue(value interface{}) Example {
	e.Value = value
	return e
}

// WithExternalValue sets the external value reference
func (e Example) WithExternalValue(url string) Example {
	e.ExternalValue = url
	return e
}
