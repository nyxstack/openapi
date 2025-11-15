package openapi

// Operation represents an operation in OpenAPI
type Operation struct {
	Tags         []string              `json:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty"`
	OperationID  string                `json:"operationId,omitempty"`
	Parameters   []Parameter           `json:"parameters,omitempty"`
	RequestBody  *RequestBody          `json:"requestBody,omitempty"`
	Responses    map[string]Response   `json:"responses"`
	Callbacks    map[string]Callback   `json:"callbacks,omitempty"`
	Deprecated   bool                  `json:"deprecated,omitempty"`
	Security     []SecurityRequirement `json:"security,omitempty"`
	Servers      []Server              `json:"servers,omitempty"`
}

// NewOperation creates a new operation with basic settings
func NewOperation(operationID, summary, description string) Operation {
	return Operation{
		OperationID: operationID,
		Summary:     summary,
		Description: description,
		Responses:   make(map[string]Response),
		Tags:        []string{},
		Parameters:  []Parameter{},
	}
}

// WithTags adds tags to an operation
func (o Operation) WithTags(tags ...string) Operation {
	o.Tags = append(o.Tags, tags...)
	return o
}

// WithTag adds a single tag to an operation
func (o Operation) WithTag(tag string) Operation {
	o.Tags = append(o.Tags, tag)
	return o
}

// WithDeprecated marks an operation as deprecated
func (o Operation) WithDeprecated() Operation {
	o.Deprecated = true
	return o
}

// WithParameter adds a parameter to an operation
func (o Operation) WithParameter(param Parameter) Operation {
	o.Parameters = append(o.Parameters, param)
	return o
}

// WithPathParameter adds a path parameter
func (o Operation) WithPathParameter(name, description string, schema *Schema) Operation {
	return o.WithParameter(Parameter{
		Name:        name,
		In:          "path",
		Description: description,
		Required:    true,
		Schema:      schema,
	})
}

// WithQueryParameter adds a query parameter
func (o Operation) WithQueryParameter(name, description string, required bool, schema *Schema) Operation {
	return o.WithParameter(Parameter{
		Name:        name,
		In:          "query",
		Description: description,
		Required:    required,
		Schema:      schema,
	})
}

// WithHeaderParameter adds a header parameter
func (o Operation) WithHeaderParameter(name, description string, required bool, schema *Schema) Operation {
	return o.WithParameter(Parameter{
		Name:        name,
		In:          "header",
		Description: description,
		Required:    required,
		Schema:      schema,
	})
}

// WithRequestBody adds a request body to an operation
func (o Operation) WithRequestBody(description string, required bool, content map[string]MediaType) Operation {
	o.RequestBody = &RequestBody{
		Description: description,
		Required:    required,
		Content:     content,
	}
	return o
}

// WithJSONRequestBody adds a JSON request body
func (o Operation) WithJSONRequestBody(description string, required bool, schema *Schema) Operation {
	content := map[string]MediaType{
		"application/json": {
			Schema: schema,
		},
	}
	return o.WithRequestBody(description, required, content)
}

// WithResponse adds a response to an operation
func (o Operation) WithResponse(code, description string, response Response) Operation {
	o.Responses[code] = response
	return o
}

// WithJSONResponse adds a JSON response
func (o Operation) WithJSONResponse(code, description string, schema *Schema) Operation {
	response := Response{
		Description: description,
		Content: map[string]MediaType{
			"application/json": {
				Schema: schema,
			},
		},
	}
	return o.WithResponse(code, description, response)
}

// WithOkResponse adds a 200 OK response
func (o Operation) WithOkResponse(description string, schema *Schema) Operation {
	return o.WithJSONResponse("200", description, schema)
}

// WithCreatedResponse adds a 201 Created response
func (o Operation) WithCreatedResponse(description string, schema *Schema) Operation {
	return o.WithJSONResponse("201", description, schema)
}

// WithNoContentResponse adds a 204 No Content response
func (o Operation) WithNoContentResponse() Operation {
	return o.WithResponse("204", "No Content", Response{
		Description: "No Content",
	})
}

// WithBadRequestResponse adds a 400 Bad Request response
func (o Operation) WithBadRequestResponse(description string) Operation {
	return o.WithResponse("400", description, Response{
		Description: description,
	})
}

// WithUnauthorizedResponse adds a 401 Unauthorized response
func (o Operation) WithUnauthorizedResponse(description string) Operation {
	return o.WithResponse("401", description, Response{
		Description: description,
	})
}

// WithForbiddenResponse adds a 403 Forbidden response
func (o Operation) WithForbiddenResponse(description string) Operation {
	return o.WithResponse("403", description, Response{
		Description: description,
	})
}

// WithNotFoundResponse adds a 404 Not Found response
func (o Operation) WithNotFoundResponse(description string) Operation {
	return o.WithResponse("404", description, Response{
		Description: description,
	})
}

// WithInternalServerErrorResponse adds a 500 Internal Server Error response
func (o Operation) WithInternalServerErrorResponse(description string) Operation {
	return o.WithResponse("500", description, Response{
		Description: description,
	})
}

// WithExternalDocs adds external documentation to an operation
func (o Operation) WithExternalDocs(url, description string) Operation {
	o.ExternalDocs = &ExternalDocs{
		URL:         url,
		Description: description,
	}
	return o
}
