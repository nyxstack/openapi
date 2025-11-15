package openapi

// RequestBody represents a request body in OpenAPI
type RequestBody struct {
	Ref         string               `json:"$ref,omitempty"`
	Description string               `json:"description,omitempty"`
	Content     map[string]MediaType `json:"content"`
	Required    bool                 `json:"required,omitempty"`
}

// NewRequestBody creates a new request body
func NewRequestBody(description string, required bool) RequestBody {
	return RequestBody{
		Description: description,
		Required:    required,
		Content:     make(map[string]MediaType),
	}
}

// NewJSONRequestBody creates a JSON request body
func NewJSONRequestBody(description string, required bool, schema *Schema) RequestBody {
	return RequestBody{
		Description: description,
		Required:    required,
		Content: map[string]MediaType{
			"application/json": {
				Schema: schema,
			},
		},
	}
}

// WithContent adds content to the request body
func (r RequestBody) WithContent(mediaType string, content MediaType) RequestBody {
	if r.Content == nil {
		r.Content = make(map[string]MediaType)
	}
	r.Content[mediaType] = content
	return r
}

// WithJSONContent adds JSON content
func (r RequestBody) WithJSONContent(schema *Schema) RequestBody {
	return r.WithContent("application/json", MediaType{
		Schema: schema,
	})
}

// WithRequired sets whether the request body is required
func (r RequestBody) WithRequired(required bool) RequestBody {
	r.Required = required
	return r
}
