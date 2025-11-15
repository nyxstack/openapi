package main

import (
	"fmt"
	"log"

	"github.com/nyxstack/openapi"
)

func main() {
	// Create a comprehensive OpenAPI document
	doc := buildCompleteAPI()

	// Marshal to JSON
	jsonBytes, err := doc.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
}

func buildCompleteAPI() *openapi.Document {
	// Create base document
	doc := openapi.NewDocument("Pet Store API", "1.0.0")

	// Add metadata
	doc.WithInfo("A sample Pet Store Server based on the OpenAPI 3.0 specification", "http://swagger.io/terms/").
		WithContact("API Support", "http://www.swagger.io/support", "support@swagger.io").
		WithLicense("Apache 2.0", "http://www.apache.org/licenses/LICENSE-2.0.html")

	// Add servers
	doc.AddServer("https://petstore3.swagger.io/api/v3", "Production server").
		AddServer("https://staging.petstore3.swagger.io/api/v3", "Staging server")

	// Add tags
	doc.AddTag("pet", "Everything about your Pets").
		AddTag("store", "Access to Petstore orders").
		AddTag("user", "Operations about user")

	// Add security schemes
	doc.AddSecurityScheme("petstore_auth", openapi.SecurityScheme{
		Type: "oauth2",
		Flows: &openapi.OAuthFlows{
			Implicit: &openapi.OAuthFlow{
				AuthorizationUrl: "https://petstore3.swagger.io/oauth/authorize",
				Scopes: map[string]string{
					"write:pets": "modify pets in your account",
					"read:pets":  "read your pets",
				},
			},
		},
	})

	doc.AddSecurityScheme("api_key", openapi.SecurityScheme{
		Type: "apiKey",
		Name: "api_key",
		In:   "header",
	})

	// Add reusable schemas
	addSchemas(doc)

	// Add operations
	addPetOperations(doc)

	return doc
}

func addSchemas(doc *openapi.Document) {
	// Pet schema
	petSchema := openapi.NewObjectSchema().
		WithRequiredProperty("id", openapi.Int64Schema()).
		WithRequiredProperty("name", openapi.StringSchema("")).
		WithProperty("category", &openapi.Schema{
			Ref: "#/components/schemas/Category",
		}).
		WithProperty("photoUrls", openapi.NewArraySchema(openapi.StringSchema(""))).
		WithProperty("tags", openapi.NewArraySchema(&openapi.Schema{
			Ref: "#/components/schemas/Tag",
		})).
		WithProperty("status", func() *openapi.Schema {
			schema := openapi.StringSchema("").WithEnum("available", "pending", "sold")
			return &schema
		}())

	doc.AddSchema("Pet", petSchema)

	// Category schema
	categorySchema := openapi.NewObjectSchema().
		WithProperty("id", openapi.Int64Schema()).
		WithProperty("name", openapi.StringSchema(""))

	doc.AddSchema("Category", categorySchema)

	// Tag schema
	tagSchema := openapi.NewObjectSchema().
		WithProperty("id", openapi.Int64Schema()).
		WithProperty("name", openapi.StringSchema(""))

	doc.AddSchema("Tag", tagSchema)

	// Error schema
	errorSchema := openapi.NewObjectSchema().
		WithRequiredProperty("code", openapi.Int32Schema()).
		WithRequiredProperty("message", openapi.StringSchema(""))

	doc.AddSchema("Error", errorSchema)
}

func addPetOperations(doc *openapi.Document) {
	// GET /pets - List pets
	listPetsOp := openapi.NewOperation("listPets", "List all pets", "Returns a list of pets").
		WithTag("pet").
		WithQueryParameter("limit", "How many items to return at one time (max 100)", false, openapi.Int32Schema()).
		WithOkResponse("A paged array of pets", openapi.NewArraySchema(&openapi.Schema{
			Ref: "#/components/schemas/Pet",
		})).
		WithResponse("default", "Unexpected error", openapi.Response{
			Description: "Unexpected error",
			Content: map[string]openapi.MediaType{
				"application/json": {
					Schema: &openapi.Schema{
						Ref: "#/components/schemas/Error",
					},
				},
			},
		})

	doc.AddOperation("/pets", "GET", listPetsOp)

	// POST /pets - Create a pet
	createPetOp := openapi.NewOperation("createPet", "Create a pet", "Create a new pet").
		WithTag("pet").
		WithJSONRequestBody("Pet to add to the store", true, &openapi.Schema{
			Ref: "#/components/schemas/Pet",
		}).
		WithCreatedResponse("Pet created", &openapi.Schema{
			Ref: "#/components/schemas/Pet",
		}).
		WithBadRequestResponse("Invalid input")

	createPetOp.Security = []openapi.SecurityRequirement{
		{"petstore_auth": []string{"write:pets"}},
	}

	doc.AddOperation("/pets", "POST", createPetOp)

	// GET /pets/{petId} - Get pet by ID
	getPetOp := openapi.NewOperation("getPetById", "Find pet by ID", "Returns a single pet").
		WithTag("pet").
		WithPathParameter("petId", "ID of pet to return", openapi.Int64Schema()).
		WithOkResponse("Successful operation", &openapi.Schema{
			Ref: "#/components/schemas/Pet",
		}).
		WithBadRequestResponse("Invalid ID supplied").
		WithNotFoundResponse("Pet not found")

	getPetOp.Security = []openapi.SecurityRequirement{
		{"api_key": []string{}},
		{"petstore_auth": []string{"read:pets"}},
	}

	doc.AddOperation("/pets/{petId}", "GET", getPetOp)
}
