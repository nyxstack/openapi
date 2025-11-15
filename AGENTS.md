# AGENTS.md - AI Assistant Guide for OpenAPI Package

## Quick Overview

This package provides **OpenAPI 3.0 specification generation** for Go applications, specifically designed for the Nyx Framework. It allows automatic generation of interactive API documentation from Go structs and route definitions.

**Key Purpose**: Transform Go code into OpenAPI specs → Generate interactive docs at `/docs` endpoint

## Core Usage Pattern

```go
// 1. Create document
doc := openapi.NewDocument("My API", "1.0.0")

// 2. Chain configuration methods
doc.WithInfo("API description", "").
    WithContact("Team", "https://example.com", "team@example.com").
    WithLicense("MIT", "https://opensource.org/licenses/MIT").
    AddServer("https://api.example.com", "Production")

// 3. Add paths and operations
doc.AddOperation("/users/{id}", "GET", openapi.Operation{
    Summary:     "Get user by ID",
    OperationID: "getUser",
    Parameters: []*openapi.Parameter{{
        Name:     "id",
        In:       "path", 
        Required: true,
        Schema:   &openapi.Schema{Type: "string"},
    }},
    Responses: map[string]*openapi.Response{
        "200": {
            Description: "User found",
            Content: map[string]*openapi.MediaType{
                "application/json": {
                    Schema: &openapi.Schema{
                        Type: "object",
                        Properties: map[string]*openapi.Schema{
                            "id":   {Type: "string"},
                            "name": {Type: "string"},
                        },
                    },
                },
            },
        },
    },
})
```

## Key Types & Their Roles

| Type | Purpose | When to Use |
|------|---------|-------------|
| `Document` | Root OpenAPI container | Always - your starting point |
| `PathItem` | Groups operations for one URL path | When defining API endpoints |
| `Operation` | Single HTTP method handler | For each GET, POST, PUT, DELETE, etc. |
| `Schema` | Data structure definitions | Request/response body shapes |
| `Components` | Reusable elements | Shared schemas, auth schemes |
| `Parameter` | Query/path/header params | URL params, query strings |
| `Response` | HTTP response definition | Different status codes |
| `MediaType` | Content-Type specifications | JSON, XML, form data |

## Essential Methods

### Document Creation & Configuration
- `NewDocument(title, version)` - Create new doc
- `WithInfo(description, terms)` - Add API description
- `WithContact(name, url, email)` - Contact info
- `WithLicense(name, url)` - License info
- `AddServer(url, description)` - Add server

### Adding Content
- `AddOperation(path, method, operation)` - Add endpoint
- `AddSchema(name, schema)` - Add reusable schema
- `AddSecurityScheme(name, scheme)` - Add auth method
- `AddTag(name, description)` - Group endpoints

### Output
- `ToJSON()` - Get JSON bytes
- `ToJSONString()` - Get JSON string

## Common Patterns

### Basic API Documentation
```go
doc := openapi.NewDocument("User API", "1.0.0").
    WithInfo("Manages user accounts", "").
    AddServer("https://api.example.com", "Production")
```

### With Authentication
```go
doc.AddSecurityScheme("bearerAuth", openapi.SecurityScheme{
    Type:         "http",
    Scheme:       "bearer",
    BearerFormat: "JWT",
})

// Apply to operations
operation.Security = []openapi.SecurityRequirement{
    {"bearerAuth": []string{}},
}
```

### Reusable Schemas
```go
doc.AddSchema("User", openapi.Schema{
    Type: "object",
    Properties: map[string]*openapi.Schema{
        "id":    {Type: "string"},
        "email": {Type: "string", Format: "email"},
        "name":  {Type: "string"},
    },
    Required: []string{"id", "email"},
})

// Reference in operations
schema := &openapi.Schema{
    Ref: "#/components/schemas/User",
}
```

## Framework Integration Notes

- **Automatic Generation**: Nyx Framework auto-generates OpenAPI docs from controller structure
- **Endpoints**: 
  - `/openapi.json` - JSON specification
  - `/openapi.yaml` - YAML specification  
  - `/docs` - Interactive ScalarUI documentation
- **Controller Mapping**: URL paths derived from embedded controller hierarchy
- **Type Safety**: Schemas generated from Go struct tags

## AI Assistant Guidelines

### When to Suggest This Package
✅ **Use when**:
- User wants OpenAPI/Swagger documentation
- Building REST APIs in Go
- Need interactive API documentation
- Want type-safe spec generation

❌ **Don't suggest when**:
- User wants GraphQL schema
- Only need simple HTTP client
- Building non-API applications

### Common Mistakes to Avoid
- **Missing Required Fields**: Always set `Type` in schemas
- **Invalid Refs**: Reference format must be `#/components/schemas/SchemaName`
- **Wrong Parameter Location**: Use "query", "path", "header", or "cookie"
- **Circular References**: Avoid infinite loops in schema references

### Best Practices to Recommend
1. **Start Simple**: Begin with `NewDocument()`, add complexity gradually
2. **Reuse Components**: Put common schemas in `Components` section
3. **Validate Output**: Use online OpenAPI validators
4. **Example Values**: Add examples for better documentation
5. **Error Responses**: Document 4xx/5xx responses

### Debugging Help
- **JSON Output**: Use `doc.ToJSONString()` to inspect generated spec
- **Validation**: Recommend online validators like swagger.io
- **Missing Fields**: Check for required fields in OpenAPI 3.0 spec
- **Type Issues**: Ensure Go types map correctly to OpenAPI types

## Schema Type Mapping

| Go Type | OpenAPI Type | Notes |
|---------|--------------|-------|
| `string` | `string` | |
| `int`, `int32` | `integer` | format: int32 |
| `int64` | `integer` | format: int64 |
| `float32` | `number` | format: float |
| `float64` | `number` | format: double |
| `bool` | `boolean` | |
| `[]T` | `array` | items: T schema |
| `map[string]T` | `object` | additionalProperties: T |
| `struct` | `object` | properties from fields |
| `*T` | T schema | nullable: true |

## Quick Troubleshooting

**"Invalid OpenAPI spec"**
→ Check required fields: Document needs `openapi`, `info`, `paths`

**"Schema not found"**  
→ Verify reference format: `#/components/schemas/SchemaName`

**"Parameter validation error"**
→ Ensure parameter has `name`, `in`, and `schema` fields

**"Missing response"**
→ Every operation needs at least one response (usually "200")

## Integration Examples

### Nyx Framework Controller
```go
type UserController struct {
    APIController // Embedded for /api/users path
}

// Automatically documented as:
// GET /api/users/{id}
func (c *UserController) GetUser(ctx *nyx.Context) {
    // Implementation
}
```

### Manual OpenAPI Generation
```go
// Build complete spec
doc := buildOpenAPISpec()
json, _ := doc.ToJSONString()

// Serve documentation
http.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(json))
})
```

---

**Remember**: This package generates OpenAPI 3.0 specifications. Always validate output and test with real API documentation tools.