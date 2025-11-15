package openapi

// MediaType represents a media type in OpenAPI
type MediaType struct {
	Schema   *Schema             `json:"schema,omitempty"`
	Example  interface{}         `json:"example,omitempty"`
	Examples map[string]Example  `json:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty"`
}

// NewMediaType creates a new media type
func NewMediaType() MediaType {
	return MediaType{
		Examples: make(map[string]Example),
		Encoding: make(map[string]Encoding),
	}
}

// NewJSONMediaType creates a JSON media type with a schema
func NewJSONMediaType(schema *Schema) MediaType {
	return MediaType{
		Schema: schema,
	}
}

// WithSchema sets the schema for the media type
func (m MediaType) WithSchema(schema *Schema) MediaType {
	m.Schema = schema
	return m
}

// WithExample sets an example for the media type
func (m MediaType) WithExample(example interface{}) MediaType {
	m.Example = example
	return m
}

// WithExamples adds named examples to the media type
func (m MediaType) WithExamples(name string, example Example) MediaType {
	if m.Examples == nil {
		m.Examples = make(map[string]Example)
	}
	m.Examples[name] = example
	return m
}

// WithEncoding adds encoding information
func (m MediaType) WithEncoding(property string, encoding Encoding) MediaType {
	if m.Encoding == nil {
		m.Encoding = make(map[string]Encoding)
	}
	m.Encoding[property] = encoding
	return m
}
