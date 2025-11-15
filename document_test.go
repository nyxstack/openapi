package openapi

import (
	"encoding/json"
	"testing"
)

func TestNewDocument(t *testing.T) {
	doc := NewDocument("Test API", "1.0.0")

	if doc.OpenAPI != "3.1.0" {
		t.Errorf("Expected OpenAPI version '3.1.0', got '%s'", doc.OpenAPI)
	}

	if doc.Info.Title != "Test API" {
		t.Errorf("Expected title 'Test API', got '%s'", doc.Info.Title)
	}

	if doc.Info.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", doc.Info.Version)
	}

	if doc.Paths == nil {
		t.Error("Expected Paths to be initialized")
	}

	if doc.Tags == nil {
		t.Error("Expected Tags to be initialized")
	}
}

func TestDocumentMarshalJSON(t *testing.T) {
	doc := NewDocument("Test API", "1.0.0")
	doc.WithInfo("A test API", "")

	data, err := json.Marshal(doc)
	if err != nil {
		t.Fatalf("Error marshaling document: %v", err)
	}

	// Test that we can unmarshal it back
	var unmarshaled Document
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Error unmarshaling document: %v", err)
	}

	if unmarshaled.Info.Title != doc.Info.Title {
		t.Errorf("Title mismatch after roundtrip: expected '%s', got '%s'", doc.Info.Title, unmarshaled.Info.Title)
	}
}

func TestDocumentWithContact(t *testing.T) {
	doc := NewDocument("Test API", "1.0.0")
	doc.WithContact("Test Team", "https://example.com", "test@example.com")

	if doc.Info.Contact == nil {
		t.Fatal("Expected Contact to be set")
	}

	if doc.Info.Contact.Name != "Test Team" {
		t.Errorf("Expected contact name 'Test Team', got '%s'", doc.Info.Contact.Name)
	}

	if doc.Info.Contact.URL != "https://example.com" {
		t.Errorf("Expected contact URL 'https://example.com', got '%s'", doc.Info.Contact.URL)
	}

	if doc.Info.Contact.Email != "test@example.com" {
		t.Errorf("Expected contact email 'test@example.com', got '%s'", doc.Info.Contact.Email)
	}
}

func TestDocumentWithLicense(t *testing.T) {
	doc := NewDocument("Test API", "1.0.0")
	doc.WithLicense("MIT", "https://opensource.org/licenses/MIT")

	if doc.Info.License == nil {
		t.Fatal("Expected License to be set")
	}

	if doc.Info.License.Name != "MIT" {
		t.Errorf("Expected license name 'MIT', got '%s'", doc.Info.License.Name)
	}

	if doc.Info.License.URL != "https://opensource.org/licenses/MIT" {
		t.Errorf("Expected license URL 'https://opensource.org/licenses/MIT', got '%s'", doc.Info.License.URL)
	}
}

func TestDocumentWithServer(t *testing.T) {
	doc := NewDocument("Test API", "1.0.0")
	doc.AddServer("https://api.example.com", "Production server")

	if len(doc.Servers) != 1 {
		t.Fatalf("Expected 1 server, got %d", len(doc.Servers))
	}

	server := doc.Servers[0]
	if server.URL != "https://api.example.com" {
		t.Errorf("Expected server URL 'https://api.example.com', got '%s'", server.URL)
	}

	if server.Description != "Production server" {
		t.Errorf("Expected server description 'Production server', got '%s'", server.Description)
	}
}
