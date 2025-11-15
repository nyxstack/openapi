package openapi

import (
	"encoding/json"
)

// Schema represents a schema in OpenAPI
type Schema struct {
	Title                string                `json:"title,omitempty"`
	MultipleOf           *float64              `json:"multipleOf,omitempty"`
	Maximum              *float64              `json:"maximum,omitempty"`
	ExclusiveMaximum     bool                  `json:"exclusiveMaximum,omitempty"`
	Minimum              *float64              `json:"minimum,omitempty"`
	ExclusiveMinimum     bool                  `json:"exclusiveMinimum,omitempty"`
	MaxLength            *int                  `json:"maxLength,omitempty"`
	MinLength            *int                  `json:"minLength,omitempty"`
	Pattern              string                `json:"pattern,omitempty"`
	MaxItems             *int                  `json:"maxItems,omitempty"`
	MinItems             *int                  `json:"minItems,omitempty"`
	UniqueItems          bool                  `json:"uniqueItems,omitempty"`
	MaxProperties        *int                  `json:"maxProperties,omitempty"`
	MinProperties        *int                  `json:"minProperties,omitempty"`
	Required             []string              `json:"required,omitempty"`
	Enum                 []interface{}         `json:"enum,omitempty"`
	Type                 string                `json:"type,omitempty"`
	AllOf                []*Schema             `json:"allOf,omitempty"`
	OneOf                []*Schema             `json:"oneOf,omitempty"`
	AnyOf                []*Schema             `json:"anyOf,omitempty"`
	Not                  *Schema               `json:"not,omitempty"`
	Items                *Schema               `json:"items,omitempty"`
	Properties           map[string]*Schema    `json:"properties,omitempty"`
	AdditionalProperties *AdditionalProperties `json:"additionalProperties,omitempty"`
	Description          string                `json:"description,omitempty"`
	Format               string                `json:"format,omitempty"`
	Default              interface{}           `json:"default,omitempty"`
	Nullable             bool                  `json:"nullable,omitempty"`
	Discriminator        *Discriminator        `json:"discriminator,omitempty"`
	ReadOnly             bool                  `json:"readOnly,omitempty"`
	WriteOnly            bool                  `json:"writeOnly,omitempty"`
	XML                  *XML                  `json:"xml,omitempty"`
	ExternalDocs         *ExternalDocs         `json:"externalDocs,omitempty"`
	Example              interface{}           `json:"example,omitempty"`
	Deprecated           bool                  `json:"deprecated,omitempty"`
}

// AdditionalProperties represents additional properties in a schema
type AdditionalProperties struct {
	Bool   *bool
	Schema *Schema
}

// MarshalJSON implements custom JSON marshaling for AdditionalProperties
func (ap AdditionalProperties) MarshalJSON() ([]byte, error) {
	if ap.Bool != nil {
		return json.Marshal(*ap.Bool)
	}
	if ap.Schema != nil {
		return json.Marshal(ap.Schema)
	}
	// Default to false if neither is set
	return json.Marshal(false)
}

// UnmarshalJSON implements custom JSON unmarshaling for AdditionalProperties
func (ap *AdditionalProperties) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as boolean first
	var boolVal bool
	if err := json.Unmarshal(data, &boolVal); err == nil {
		ap.Bool = &boolVal
		ap.Schema = nil
		return nil
	}

	// Try to unmarshal as schema
	var schema Schema
	if err := json.Unmarshal(data, &schema); err == nil {
		ap.Bool = nil
		ap.Schema = &schema
		return nil
	}

	return json.Unmarshal(data, &boolVal) // Return the boolean unmarshal error
}

// Discriminator represents a discriminator in OpenAPI
type Discriminator struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}

// XML represents XML metadata in OpenAPI
type XML struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}

// NewStringSchema creates a string schema
func NewStringSchema() *Schema {
	return &Schema{
		Type: "string",
	}
}

// NewIntegerSchema creates an integer schema
func NewIntegerSchema() *Schema {
	return &Schema{
		Type: "integer",
	}
}

// NewNumberSchema creates a number schema
func NewNumberSchema() *Schema {
	return &Schema{
		Type: "number",
	}
}

// NewBooleanSchema creates a boolean schema
func NewBooleanSchema() *Schema {
	return &Schema{
		Type: "boolean",
	}
}

// NewArraySchema creates an array schema
func NewArraySchema(items *Schema) *Schema {
	return &Schema{
		Type:  "array",
		Items: items,
	}
}

// NewObjectSchema creates an object schema
func NewObjectSchema() *Schema {
	return &Schema{
		Type:       "object",
		Properties: make(map[string]*Schema),
	}
}

// WithFormat sets the format of a schema
func (s Schema) WithFormat(format string) Schema {
	s.Format = format
	return s
}

// WithDescription sets the description of a schema
func (s Schema) WithDescription(description string) Schema {
	s.Description = description
	return s
}

// WithTitle sets the title of a schema
func (s Schema) WithTitle(title string) Schema {
	s.Title = title
	return s
}

// WithExample sets an example for a schema
func (s Schema) WithExample(example interface{}) Schema {
	s.Example = example
	return s
}

// WithDefault sets a default value for a schema
func (s Schema) WithDefault(defaultValue interface{}) Schema {
	s.Default = defaultValue
	return s
}

// WithEnum sets enum values for a schema
func (s Schema) WithEnum(values ...interface{}) Schema {
	s.Enum = values
	return s
}

// WithMinLength sets minimum length for string schemas
func (s Schema) WithMinLength(min int) Schema {
	s.MinLength = &min
	return s
}

// WithMaxLength sets maximum length for string schemas
func (s Schema) WithMaxLength(max int) Schema {
	s.MaxLength = &max
	return s
}

// WithPattern sets a pattern for string schemas
func (s Schema) WithPattern(pattern string) Schema {
	s.Pattern = pattern
	return s
}

// WithMinimum sets minimum value for number schemas
func (s Schema) WithMinimum(min float64) Schema {
	s.Minimum = &min
	return s
}

// WithMaximum sets maximum value for number schemas
func (s Schema) WithMaximum(max float64) Schema {
	s.Maximum = &max
	return s
}

// WithMinItems sets minimum items for array schemas
func (s Schema) WithMinItems(min int) Schema {
	s.MinItems = &min
	return s
}

// WithMaxItems sets maximum items for array schemas
func (s Schema) WithMaxItems(max int) Schema {
	s.MaxItems = &max
	return s
}

// WithUniqueItems sets unique items constraint for array schemas
func (s Schema) WithUniqueItems(unique bool) Schema {
	s.UniqueItems = unique
	return s
}

// WithProperty adds a property to object schemas
func (s Schema) WithProperty(name string, schema *Schema) Schema {
	if s.Properties == nil {
		s.Properties = make(map[string]*Schema)
	}
	s.Properties[name] = schema
	return s
}

// WithRequiredProperty adds a required property to object schemas
func (s Schema) WithRequiredProperty(name string, schema *Schema) Schema {
	s = s.WithProperty(name, schema)
	s.Required = append(s.Required, name)
	return s
}

// WithRequired sets required fields for object schemas
func (s Schema) WithRequired(fields ...string) Schema {
	s.Required = append(s.Required, fields...)
	return s
}

// WithNullable makes a schema nullable
func (s Schema) WithNullable(nullable bool) Schema {
	s.Nullable = nullable
	return s
}

// WithReadOnly makes a schema read-only
func (s Schema) WithReadOnly(readOnly bool) Schema {
	s.ReadOnly = readOnly
	return s
}

// WithWriteOnly makes a schema write-only
func (s Schema) WithWriteOnly(writeOnly bool) Schema {
	s.WriteOnly = writeOnly
	return s
}

// WithDeprecated marks a schema as deprecated
func (s Schema) WithDeprecated(deprecated bool) Schema {
	s.Deprecated = deprecated
	return s
}

// Common schema constructors for convenience

// StringSchema creates a string schema with format
func StringSchema(format string) *Schema {
	schema := NewStringSchema()
	if format != "" {
		*schema = schema.WithFormat(format)
	}
	return schema
}

// EmailSchema creates an email string schema
func EmailSchema() *Schema {
	return StringSchema("email")
}

// DateTimeSchema creates a date-time string schema
func DateTimeSchema() *Schema {
	return StringSchema("date-time")
}

// DateSchema creates a date string schema
func DateSchema() *Schema {
	return StringSchema("date")
}

// UUIDSchema creates a UUID string schema
func UUIDSchema() *Schema {
	return StringSchema("uuid")
}

// PasswordSchema creates a password string schema
func PasswordSchema() *Schema {
	return StringSchema("password")
}

// Int32Schema creates an int32 integer schema
func Int32Schema() *Schema {
	schema := NewIntegerSchema().WithFormat("int32")
	return &schema
}

// Int64Schema creates an int64 integer schema
func Int64Schema() *Schema {
	schema := NewIntegerSchema().WithFormat("int64")
	return &schema
}

// FloatSchema creates a float number schema
func FloatSchema() *Schema {
	schema := NewNumberSchema().WithFormat("float")
	return &schema
}

// DoubleSchema creates a double number schema
func DoubleSchema() *Schema {
	schema := NewNumberSchema().WithFormat("double")
	return &schema
}

// IDSchema creates a common ID schema (integer with int64 format)
func IDSchema() *Schema {
	schema := Int64Schema().WithDescription("Unique identifier")
	return &schema
}

// PaginationSchema creates a common pagination object schema
func PaginationSchema() *Schema {
	pageSchema := *Int32Schema()
	pageSchema = pageSchema.WithDescription("Current page number")

	limitSchema := *Int32Schema()
	limitSchema = limitSchema.WithDescription("Number of items per page")

	totalSchema := *Int32Schema()
	totalSchema = totalSchema.WithDescription("Total number of items")

	hasNextSchema := NewBooleanSchema().WithDescription("Whether there are more pages")

	schema := NewObjectSchema().
		WithRequiredProperty("page", &pageSchema).
		WithRequiredProperty("limit", &limitSchema).
		WithRequiredProperty("total", &totalSchema).
		WithProperty("hasNext", &hasNextSchema)
	return &schema
}
