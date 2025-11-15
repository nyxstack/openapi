# OpenAPI Package

Go types and utilities for building OpenAPI 3.1.0 specifications programmatically.

## Features

### Supported OpenAPI 3.1.0 Features

- ✅ **Complete Specification Support** - All OpenAPI 3.1.0 objects implemented
- ✅ **JSON Schema Integration** - Full schema validation and documentation
- ✅ **Multiple Content Types** - JSON, XML, form data, file uploads
- ✅ **Authentication Schemes** - API keys, OAuth2, HTTP auth
- ✅ **Response Linking** - Connect operations through links
- ✅ **Webhooks & Callbacks** - Document async operations
- ✅ **Examples & Descriptions** - Rich documentation with examples

## Installation

```bash
go get github.com/nyxstack/openapi
```

## Overview

This package provides Go types that represent the complete OpenAPI 3.1.0 specification. You use these types to programmatically build OpenAPI documents, then marshal them to JSON or YAML format.

## Usage

### Basic Document Creation

```go
package main

import (
    "fmt"
    "github.com/nyxstack/openapi"
)

func main() {
    // Create a new OpenAPI document
    doc := openapi.NewDocument("My API", "1.0.0")
    
    // Add metadata
    doc.WithInfo("A sample API", "").
        WithContact("API Team", "https://example.com", "team@example.com").
        WithLicense("MIT", "https://opensource.org/licenses/MIT")
    
    // Add a server
    doc.AddServer("https://api.example.com", "Production server")
    
    // Build a simple operation
    operation := openapi.NewOperation("getUser", "Get User", "Retrieve user by ID")
    operation.WithPathParameter("id", "User ID", openapi.StringSchema(""))
    operation.WithOkResponse("User found", userSchema())
    
    // Add the operation to a path
    doc.AddOperation("/users/{id}", "GET", operation)
    
    // Marshal to JSON
    jsonBytes, err := doc.ToJSON()
    if err != nil {
        panic(err)
    }
    
    fmt.Println(string(jsonBytes))
}

func userSchema() *openapi.Schema {
    return openapi.NewObjectSchema().
        WithRequiredProperty("id", openapi.StringSchema("")).
        WithProperty("name", openapi.StringSchema("")).
        WithProperty("email", openapi.EmailSchema())
}
```

### Building Complex Schemas

```go
// Create reusable schemas
doc.AddSchema("User", *openapi.NewObjectSchema().
    WithRequiredProperty("id", openapi.StringSchema("")).
    WithRequiredProperty("email", openapi.EmailSchema()).
    WithProperty("name", openapi.StringSchema("")))

// Reference schemas in operations
schema := &openapi.Schema{
    Ref: "#/components/schemas/User",
}
```

### Authentication

```go
// Add Bearer token auth
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

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

