package openapi

// Link represents a link in OpenAPI
type Link struct {
	OperationRef string                 `json:"operationRef,omitempty"`
	OperationId  string                 `json:"operationId,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Server       *Server                `json:"server,omitempty"`
}

// NewLink creates a new link
func NewLink() Link {
	return Link{
		Parameters: make(map[string]interface{}),
	}
}

// WithOperationRef sets the operation reference
func (l Link) WithOperationRef(ref string) Link {
	l.OperationRef = ref
	return l
}

// WithOperationId sets the operation ID
func (l Link) WithOperationId(operationId string) Link {
	l.OperationId = operationId
	return l
}

// WithParameter adds a parameter to the link
func (l Link) WithParameter(name string, value interface{}) Link {
	if l.Parameters == nil {
		l.Parameters = make(map[string]interface{})
	}
	l.Parameters[name] = value
	return l
}

// WithRequestBody sets the request body for the link
func (l Link) WithRequestBody(body interface{}) Link {
	l.RequestBody = body
	return l
}

// WithDescription sets the description
func (l Link) WithDescription(description string) Link {
	l.Description = description
	return l
}

// WithServer sets the server for the link
func (l Link) WithServer(server *Server) Link {
	l.Server = server
	return l
}
