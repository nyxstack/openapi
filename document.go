package openapi

import (
	"encoding/json"
)

// ExternalDocs represents external documentation
type ExternalDocs struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
}

// SecurityRequirement represents a security requirement
type SecurityRequirement map[string][]string

// Document represents the root OpenAPI v3 document
type Document struct {
	OpenAPI      string                `json:"openapi"`
	Info         Info                  `json:"info"`
	Servers      []Server              `json:"servers,omitempty"`
	Paths        map[string]PathItem   `json:"paths"`
	Components   *Components           `json:"components,omitempty"`
	Security     []SecurityRequirement `json:"security,omitempty"`
	Tags         []Tag                 `json:"tags,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty"`
}

// NewDocument creates a new OpenAPI document with basic info
func NewDocument(title, version string) *Document {
	return &Document{
		OpenAPI: "3.0.3",
		Info: Info{
			Title:   title,
			Version: version,
		},
		Paths: make(map[string]PathItem),
		Tags:  []Tag{},
	}
}

// WithInfo sets additional info for the OpenAPI document
func (d *Document) WithInfo(description, termsOfService string) *Document {
	d.Info.Description = description
	d.Info.TermsOfService = termsOfService
	return d
}

// WithContact adds contact information to the OpenAPI document
func (d *Document) WithContact(name, url, email string) *Document {
	d.Info.Contact = &Contact{
		Name:  name,
		URL:   url,
		Email: email,
	}
	return d
}

// WithLicense adds license information to the OpenAPI document
func (d *Document) WithLicense(name, url string) *Document {
	d.Info.License = &License{
		Name: name,
		URL:  url,
	}
	return d
}

// AddServer adds a server to the document
func (d *Document) AddServer(url, description string) *Document {
	d.Servers = append(d.Servers, Server{
		URL:         url,
		Description: description,
		Variables:   make(map[string]ServerVariable),
	})
	return d
}

// AddTag adds a tag to the document
func (d *Document) AddTag(name, description string) *Document {
	d.Tags = append(d.Tags, Tag{
		Name:        name,
		Description: description,
	})
	return d
}

// AddTagWithDocs adds a tag with external documentation
func (d *Document) AddTagWithDocs(name, description, docsURL, docsDescription string) *Document {
	d.Tags = append(d.Tags, Tag{
		Name:        name,
		Description: description,
		ExternalDocs: &ExternalDocs{
			URL:         docsURL,
			Description: docsDescription,
		},
	})
	return d
}

// SetExternalDocs sets external documentation for the entire API
func (d *Document) SetExternalDocs(url, description string) *Document {
	d.ExternalDocs = &ExternalDocs{
		URL:         url,
		Description: description,
	}
	return d
}

// AddPath adds a path to the document with an empty PathItem
func (d *Document) AddPath(path string) *PathItem {
	if d.Paths == nil {
		d.Paths = make(map[string]PathItem)
	}
	pathItem := PathItem{}
	d.Paths[path] = pathItem
	return &pathItem
}

// GetPath gets a path item or creates it if it doesn't exist
func (d *Document) GetPath(path string) *PathItem {
	if d.Paths == nil {
		d.Paths = make(map[string]PathItem)
	}
	if pathItem, exists := d.Paths[path]; exists {
		return &pathItem
	}
	return d.AddPath(path)
}

// SetPath sets a complete path item
func (d *Document) SetPath(path string, pathItem PathItem) *Document {
	if d.Paths == nil {
		d.Paths = make(map[string]PathItem)
	}
	d.Paths[path] = pathItem
	return d
}

// AddOperation adds an operation to a specific path and method
func (d *Document) AddOperation(path, method string, operation Operation) *Document {
	pathItem := d.GetPath(path)

	switch method {
	case "GET":
		pathItem.Get = &operation
	case "POST":
		pathItem.Post = &operation
	case "PUT":
		pathItem.Put = &operation
	case "DELETE":
		pathItem.Delete = &operation
	case "PATCH":
		pathItem.Patch = &operation
	case "HEAD":
		pathItem.Head = &operation
	case "OPTIONS":
		pathItem.Options = &operation
	case "TRACE":
		pathItem.Trace = &operation
	}

	d.Paths[path] = *pathItem
	return d
}

// AddComponents adds or updates components section
func (d *Document) AddComponents() *Components {
	if d.Components == nil {
		d.Components = &Components{
			Schemas:         make(map[string]*Schema),
			Responses:       make(map[string]Response),
			Parameters:      make(map[string]Parameter),
			Examples:        make(map[string]Example),
			RequestBodies:   make(map[string]RequestBody),
			Headers:         make(map[string]Header),
			SecuritySchemes: make(map[string]SecurityScheme),
			Links:           make(map[string]Link),
			Callbacks:       make(map[string]Callback),
		}
	}
	return d.Components
}

// AddSchema adds a schema to components
func (d *Document) AddSchema(name string, schema Schema) *Document {
	components := d.AddComponents()
	components.Schemas[name] = &schema
	return d
}

// AddSecurityScheme adds a security scheme to components
func (d *Document) AddSecurityScheme(name string, scheme SecurityScheme) *Document {
	components := d.AddComponents()
	components.SecuritySchemes[name] = scheme
	return d
}

// AddSecurityRequirement adds a security requirement at document level
func (d *Document) AddSecurityRequirement(requirement SecurityRequirement) *Document {
	d.Security = append(d.Security, requirement)
	return d
}

// ToJSON converts the OpenAPI document to JSON
func (d *Document) ToJSON() ([]byte, error) {
	return json.MarshalIndent(d, "", "  ")
}

// ToJSONString converts the OpenAPI document to JSON string
func (d *Document) ToJSONString() (string, error) {
	data, err := d.ToJSON()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
