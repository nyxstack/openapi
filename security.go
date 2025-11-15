package openapi

// SecurityScheme represents a security scheme in OpenAPI
type SecurityScheme struct {
	Type             string      `json:"type"`
	Description      string      `json:"description,omitempty"`
	Name             string      `json:"name,omitempty"`
	In               string      `json:"in,omitempty"`
	Scheme           string      `json:"scheme,omitempty"`
	BearerFormat     string      `json:"bearerFormat,omitempty"`
	Flows            *OAuthFlows `json:"flows,omitempty"`
	OpenIdConnectUrl string      `json:"openIdConnectUrl,omitempty"`
}

// OAuthFlows represents OAuth flows in OpenAPI
type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty"`
}

// OAuthFlow represents an OAuth flow in OpenAPI
type OAuthFlow struct {
	AuthorizationUrl string            `json:"authorizationUrl,omitempty"`
	TokenUrl         string            `json:"tokenUrl,omitempty"`
	RefreshUrl       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes"`
}

// NewSecurityScheme creates a new security scheme
func NewSecurityScheme(schemeType string) *SecurityScheme {
	return &SecurityScheme{
		Type: schemeType,
	}
}

// WithDescription sets the description of the security scheme
func (s SecurityScheme) WithDescription(description string) SecurityScheme {
	s.Description = description
	return s
}

// WithName sets the name of the security scheme
func (s SecurityScheme) WithName(name string) SecurityScheme {
	s.Name = name
	return s
}

// WithIn sets the location of the API key
func (s SecurityScheme) WithIn(in string) SecurityScheme {
	s.In = in
	return s
}

// WithScheme sets the HTTP authorization scheme
func (s SecurityScheme) WithScheme(scheme string) SecurityScheme {
	s.Scheme = scheme
	return s
}

// WithBearerFormat sets the bearer format
func (s SecurityScheme) WithBearerFormat(format string) SecurityScheme {
	s.BearerFormat = format
	return s
}

// WithFlows sets the OAuth flows
func (s SecurityScheme) WithFlows(flows *OAuthFlows) SecurityScheme {
	s.Flows = flows
	return s
}

// WithOpenIdConnectUrl sets the OpenID Connect URL
func (s SecurityScheme) WithOpenIdConnectUrl(url string) SecurityScheme {
	s.OpenIdConnectUrl = url
	return s
}

// NewOAuthFlows creates a new OAuth flows object
func NewOAuthFlows() *OAuthFlows {
	return &OAuthFlows{}
}

// WithImplicit sets the implicit flow
func (f *OAuthFlows) WithImplicit(flow *OAuthFlow) *OAuthFlows {
	f.Implicit = flow
	return f
}

// WithPassword sets the password flow
func (f *OAuthFlows) WithPassword(flow *OAuthFlow) *OAuthFlows {
	f.Password = flow
	return f
}

// WithClientCredentials sets the client credentials flow
func (f *OAuthFlows) WithClientCredentials(flow *OAuthFlow) *OAuthFlows {
	f.ClientCredentials = flow
	return f
}

// WithAuthorizationCode sets the authorization code flow
func (f *OAuthFlows) WithAuthorizationCode(flow *OAuthFlow) *OAuthFlows {
	f.AuthorizationCode = flow
	return f
}

// NewOAuthFlow creates a new OAuth flow
func NewOAuthFlow() *OAuthFlow {
	return &OAuthFlow{
		Scopes: make(map[string]string),
	}
}

// WithAuthorizationUrl sets the authorization URL
func (f *OAuthFlow) WithAuthorizationUrl(url string) *OAuthFlow {
	f.AuthorizationUrl = url
	return f
}

// WithTokenUrl sets the token URL
func (f *OAuthFlow) WithTokenUrl(url string) *OAuthFlow {
	f.TokenUrl = url
	return f
}

// WithRefreshUrl sets the refresh URL
func (f *OAuthFlow) WithRefreshUrl(url string) *OAuthFlow {
	f.RefreshUrl = url
	return f
}

// WithScope adds a scope to the OAuth flow
func (f *OAuthFlow) WithScope(scope, description string) *OAuthFlow {
	f.Scopes[scope] = description
	return f
}

// Convenience functions for common security schemes

// NewAPIKeySecurityScheme creates a new API key security scheme
func NewAPIKeySecurityScheme(name, in string) *SecurityScheme {
	return &SecurityScheme{
		Type: "apiKey",
		Name: name,
		In:   in,
	}
}

// NewHTTPSecurityScheme creates a new HTTP security scheme
func NewHTTPSecurityScheme(scheme string) *SecurityScheme {
	return &SecurityScheme{
		Type:   "http",
		Scheme: scheme,
	}
}

// NewBearerSecurityScheme creates a new Bearer token security scheme
func NewBearerSecurityScheme() *SecurityScheme {
	return &SecurityScheme{
		Type:   "http",
		Scheme: "bearer",
	}
}

// NewOAuth2SecurityScheme creates a new OAuth2 security scheme
func NewOAuth2SecurityScheme() *SecurityScheme {
	return &SecurityScheme{
		Type: "oauth2",
	}
}

// NewOpenIdConnectSecurityScheme creates a new OpenID Connect security scheme
func NewOpenIdConnectSecurityScheme(url string) *SecurityScheme {
	return &SecurityScheme{
		Type:             "openIdConnect",
		OpenIdConnectUrl: url,
	}
}

// Convenience functions for common security schemes

// JWTAuth creates a JWT Bearer token security scheme
func JWTAuth() SecurityScheme {
	return SecurityScheme{
		Type:   "http",
		Scheme: "bearer",
	}
}

// APIKeyInHeader creates an API key security scheme in the header
func APIKeyInHeader(name string) SecurityScheme {
	return SecurityScheme{
		Type: "apiKey",
		Name: name,
		In:   "header",
	}
}

// APIKeyInQuery creates an API key security scheme in the query
func APIKeyInQuery(name string) SecurityScheme {
	return SecurityScheme{
		Type: "apiKey",
		Name: name,
		In:   "query",
	}
}

// APIKeyInCookie creates an API key security scheme in cookies
func APIKeyInCookie(name string) SecurityScheme {
	return SecurityScheme{
		Type: "apiKey",
		Name: name,
		In:   "cookie",
	}
}

// Security requirement helper functions

// RequireBearer creates a security requirement for Bearer auth
func RequireBearer(schemeName string) SecurityRequirement {
	return SecurityRequirement{
		schemeName: []string{},
	}
}

// RequireAPIKey creates a security requirement for API key auth
func RequireAPIKey(schemeName string) SecurityRequirement {
	return SecurityRequirement{
		schemeName: []string{},
	}
}

// RequireOAuth creates a security requirement for OAuth with scopes
func RequireOAuth(schemeName string, scopes ...string) SecurityRequirement {
	return SecurityRequirement{
		schemeName: scopes,
	}
}
