package openapi

// Components represents the components object in OpenAPI
type Components struct {
	Schemas         map[string]*Schema        `json:"schemas,omitempty"`
	Responses       map[string]Response       `json:"responses,omitempty"`
	Parameters      map[string]Parameter      `json:"parameters,omitempty"`
	Examples        map[string]Example        `json:"examples,omitempty"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies,omitempty"`
	Headers         map[string]Header         `json:"headers,omitempty"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
	Links           map[string]Link           `json:"links,omitempty"`
	Callbacks       map[string]Callback       `json:"callbacks,omitempty"`
}

// NewComponents creates a new components object
func NewComponents() *Components {
	return &Components{
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
