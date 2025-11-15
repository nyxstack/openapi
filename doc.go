// Package openapi provides Go types and utilities for building OpenAPI 3.1.0 specifications.
//
// This package implements the complete OpenAPI 3.1.0 specification as Go structs,
// allowing you to programmatically build OpenAPI documents and marshal them to JSON or YAML.
//
// Basic usage:
//
//	doc := openapi.NewDocument("My API", "1.0.0")
//	doc.WithInfo("API description", "").
//		WithContact("Team", "https://example.com", "team@example.com")
//
//	operation := openapi.NewOperation("getUser", "Get User", "Get user by ID")
//	doc.AddOperation("/users/{id}", "GET", operation)
//
//	jsonBytes, err := doc.ToJSON()
//	if err != nil {
//		// handle error
//	}
//
// The package supports all OpenAPI 3.1.0 features including schemas, parameters,
// responses, security schemes, callbacks, links, and more.
package openapi
