# AGENTS.md - AI Assistant Guide for OpenAPI Package

## Quick Overview

This package provides **Go types and utilities for building OpenAPI 3.1 specifications** programmatically. It's a document builder library - you use these types to manually construct OpenAPI documents, then marshal them to JSON/YAML.

**Key Purpose**: Provide Go types representing OpenAPI 3.1 spec → Build documents manually → Marshal to JSON/YAML

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

## AI Assistant Guidelines

### When to Suggest This Package
✅ **Use when**:
- User wants to build OpenAPI/Swagger specifications programmatically
- Building REST API documentation in Go
- Need to generate OpenAPI JSON/YAML from Go code
- Want type-safe OpenAPI document construction

❌ **Don't suggest when**:
- User wants GraphQL schema
- Only need simple HTTP client
- Building non-API applications

**Note**: If user wants automatic generation, this package is perfect as the foundation - use these types to build generation tools!

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
- **Validation**: Recommend online validators like swagger.io or redocly
- **Type Issues**: Ensure Go types map correctly to OpenAPI types
- **Compilation**: If it compiles and marshals, the spec structure is correct

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
→ Use online validators like swagger.io to check the generated JSON

**"Schema not found"**  
→ Verify reference format: `#/components/schemas/SchemaName`

**"Type errors"**
→ Ensure parameter has `name`, `in`, and `schema` fields

**"Empty responses"**
→ Every operation needs at least one response (usually "200")

## Usage Examples

### Basic Document Building
```go
package main

import (
    "fmt"
    "github.com/nyxstack/openapi"
)

func main() {
    doc := openapi.NewDocument("My API", "1.0.0")
    doc.WithInfo("Sample API", "").WithContact("Team", "https://example.com", "team@example.com")
    
    operation := openapi.NewOperation("getUser", "Get User", "Get user by ID")
    operation.WithPathParameter("id", "User ID", openapi.StringSchema(""))
    operation.WithOkResponse("User found", userSchema())
    
    doc.AddOperation("/users/{id}", "GET", operation)
    
    jsonBytes, _ := doc.ToJSON()
    fmt.Println(string(jsonBytes))
}
```

### Complete API Documentation
```go
// See example/ directory for a comprehensive Pet Store API example
// that demonstrates all major OpenAPI 3.1 features
```

---

**Remember**: This package provides Go types for building OpenAPI 3.1 specifications manually. The types themselves ensure structural correctness - if it compiles and marshals, the document structure is valid.