package openapi

// Response represents a response in OpenAPI
type Response struct {
	Description string               `json:"description"`
	Headers     map[string]Header    `json:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
	Links       map[string]Link      `json:"links,omitempty"`
}

// NewResponse creates a new response
func NewResponse(description string) Response {
	return Response{
		Description: description,
		Headers:     make(map[string]Header),
		Content:     make(map[string]MediaType),
		Links:       make(map[string]Link),
	}
}

// NewJSONResponse creates a response with JSON content
func NewJSONResponse(description string, schema *Schema) Response {
	return Response{
		Description: description,
		Content: map[string]MediaType{
			"application/json": {
				Schema: schema,
			},
		},
	}
}

// WithHeader adds a header to the response
func (r Response) WithHeader(name string, header Header) Response {
	if r.Headers == nil {
		r.Headers = make(map[string]Header)
	}
	r.Headers[name] = header
	return r
}

// WithContent adds content to the response
func (r Response) WithContent(mediaType string, content MediaType) Response {
	if r.Content == nil {
		r.Content = make(map[string]MediaType)
	}
	r.Content[mediaType] = content
	return r
}

// WithJSONContent adds JSON content to the response
func (r Response) WithJSONContent(schema *Schema) Response {
	return r.WithContent("application/json", MediaType{
		Schema: schema,
	})
}

// WithLink adds a link to the response
func (r Response) WithLink(name string, link Link) Response {
	if r.Links == nil {
		r.Links = make(map[string]Link)
	}
	r.Links[name] = link
	return r
}
