# OpenAPI Package

This package provides OpenAPI 3.0 specification generation for the Nyx Framework. It allows automatic generation of API documentation from Go structs and route definitions.

## Overview

The OpenAPI package implements the complete OpenAPI 3.0 specification structure, enabling the Nyx Framework to automatically generate comprehensive API documentation. This documentation is served at the `/docs` endpoint using ScalarUI.

## Package Structure

### Core Components

- **`document.go`** - Main OpenAPI document structure and builder
- **`paths.go`** - HTTP paths and operations management
- **`operation.go`** - Individual HTTP operation definitions
- **`components.go`** - Reusable components (schemas, parameters, etc.)
- **`schema.go`** - JSON Schema definitions for request/response bodies

### Supporting Files

- **`info.go`** - API metadata (title, version, description)
- **`server.go`** - Server definitions and base URLs
- **`parameter.go`** - Query, path, and header parameter definitions
- **`request_body.go`** - Request body specifications
- **`response.go`** - Response definitions with status codes
- **`media_type.go`** - Content type specifications
- **`tag.go`** - API endpoint grouping and organization

### Additional Features

- **`security.go`** - Authentication and authorization schemes
- **`callback.go`** - Webhook and callback definitions
- **`link.go`** - Response links and relationships
- **`example.go`** - Example values for documentation
- **`header.go`** - HTTP header definitions
- **`encoding.go`** - Content encoding specifications
- **`contact.go`** - API contact information
- **`license.go`** - API licensing information

## Usage

### Basic Document Creation

```go
doc := &openapi.Document{
    OpenAPI: "3.0.3",
    Info: &openapi.Info{
        Title:       "My API",
        Version:     "1.0.0",
        Description: "API documentation",
    },
    Paths: make(map[string]*openapi.PathItem),
}
```

### Adding Operations

```go
// Add a GET operation
pathItem := &openapi.PathItem{
    Get: &openapi.Operation{
        Summary:     "Get user",
        Description: "Retrieve user by ID",
        OperationID: "getUser",
        Tags:        []string{"users"},
        Parameters: []*openapi.Parameter{
            {
                Name:     "id",
                In:       "path",
                Required: true,
                Schema:   &openapi.Schema{Type: "string"},
            },
        },
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
    },
}

doc.Paths["/users/{id}"] = pathItem
```

### Schema Integration

The OpenAPI package integrates with the Nyx Framework's schema system to automatically generate schemas from Go structs:

```go
type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// Schema generation happens automatically when registering controllers
```

## Integration with Nyx Framework

### Automatic Documentation

The Nyx Framework automatically generates OpenAPI documentation based on:

1. **Controller Structure** - Paths derived from embedded controller hierarchy
2. **Route Definitions** - HTTP methods and URL patterns
3. **Go Structs** - Request/response schemas from struct tags
4. **Method Signatures** - Parameter extraction from handler functions

### Documentation Endpoints

- **`/openapi.json`** - OpenAPI specification in JSON format
- **`/openapi.yaml`** - OpenAPI specification in YAML format  
- **`/docs`** - Interactive API documentation using ScalarUI

### Example Controller Integration

```go
type UserController struct {
    APIController // Embedded for /api/users path
}

// GET /api/users/{id}
func (c *UserController) GetUser(ctx *nyx.Context) {
    // Implementation automatically documented
}
```

## Features

### Supported OpenAPI 3.0 Features

- ✅ **Complete Specification Support** - All OpenAPI 3.0 objects implemented
- ✅ **JSON Schema Integration** - Full schema validation and documentation
- ✅ **Multiple Content Types** - JSON, XML, form data, file uploads
- ✅ **Authentication Schemes** - API keys, OAuth2, HTTP auth
- ✅ **Response Linking** - Connect operations through links
- ✅ **Webhooks & Callbacks** - Document async operations
- ✅ **Examples & Descriptions** - Rich documentation with examples

### Framework Integration

- ✅ **Automatic Generation** - No manual OpenAPI writing required
- ✅ **Type Safety** - Generated from Go types and struct tags
- ✅ **Live Documentation** - Updates with code changes
- ✅ **Multiple Formats** - JSON and YAML output
- ✅ **Interactive UI** - ScalarUI for testing and exploration

## Development Notes

### Adding New Features

1. **Follow OpenAPI 3.0 Spec** - Ensure compliance with official specification
2. **Maintain Type Safety** - Use Go types to represent OpenAPI objects
3. **Support Serialization** - All structs should marshal to correct JSON/YAML
4. **Add Examples** - Include usage examples in documentation
5. **Test Integration** - Verify with Nyx Framework controller system

### Common Patterns

- **Pointer Fields** - Optional OpenAPI fields should be pointers
- **Omitempty Tags** - Use `json:",omitempty"` for optional fields
- **Validation** - Add validation methods where appropriate
- **Extensions** - Support OpenAPI extensions with `x-*` fields

### Testing

Test OpenAPI generation with:

```bash
# Build and run test server
go run main.go

# Check generated documentation
curl http://localhost:8080/openapi.json
curl http://localhost:8080/docs
```

## Future Enhancements

- [ ] **Schema Validation** - Runtime validation against OpenAPI schemas
- [ ] **Code Generation** - Generate client SDKs from OpenAPI specs
- [ ] **Advanced Examples** - More sophisticated example generation
- [ ] **Custom Extensions** - Framework-specific OpenAPI extensions
- [ ] **Performance Optimization** - Caching and lazy generation
- [ ] **Testing Utilities** - Test helpers for OpenAPI validation

## Related Packages

- **`../schema/`** - JSON Schema generation and validation
- **`../openapischema/`** - OpenAPI to JSON Schema conversion
- **`gitlab.com/go-nyx/scalarui`** - API documentation UI

This package forms the foundation for automatic API documentation in the Nyx Framework, making it easy to maintain up-to-date, interactive documentation for your APIs.