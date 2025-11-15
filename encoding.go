package openapi

// Encoding represents encoding in OpenAPI
type Encoding struct {
	ContentType   string            `json:"contentType,omitempty"`
	Headers       map[string]Header `json:"headers,omitempty"`
	Style         string            `json:"style,omitempty"`
	Explode       bool              `json:"explode,omitempty"`
	AllowReserved bool              `json:"allowReserved,omitempty"`
}

// NewEncoding creates a new encoding
func NewEncoding() Encoding {
	return Encoding{
		Headers: make(map[string]Header),
	}
}

// WithContentType sets the content type
func (e Encoding) WithContentType(contentType string) Encoding {
	e.ContentType = contentType
	return e
}

// WithHeader adds a header to the encoding
func (e Encoding) WithHeader(name string, header Header) Encoding {
	if e.Headers == nil {
		e.Headers = make(map[string]Header)
	}
	e.Headers[name] = header
	return e
}

// WithStyle sets the style
func (e Encoding) WithStyle(style string) Encoding {
	e.Style = style
	return e
}

// WithExplode sets the explode option
func (e Encoding) WithExplode(explode bool) Encoding {
	e.Explode = explode
	return e
}

// WithAllowReserved sets whether reserved characters are allowed
func (e Encoding) WithAllowReserved(allow bool) Encoding {
	e.AllowReserved = allow
	return e
}
