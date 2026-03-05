package credlyr

import "time"

// Verification represents a verification session.
type Verification struct {
	ID              string                 `json:"id"`
	Status          string                 `json:"status"`
	PolicyID        string                 `json:"policy_id,omitempty"`
	RequestedClaims []string               `json:"requested_claims"`
	// OutputClaims contains full claim values (sandbox environment only).
	// In production, use VerifiedClaims instead.
	OutputClaims map[string]interface{} `json:"output_claims,omitempty"`
	// VerifiedClaims contains list of claim names that were verified (production).
	// Actual values are masked for PII protection. In sandbox, use OutputClaims.
	VerifiedClaims []string               `json:"verified_claims,omitempty"`
	HostedURL      string                 `json:"hosted_url"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
	CompletedAt    *time.Time             `json:"completed_at,omitempty"`
	ExpiresAt      *time.Time             `json:"expires_at,omitempty"`
}

// CreateVerificationParams are the parameters for creating a verification.
type CreateVerificationParams struct {
	PolicyID        string                 `json:"policy_id,omitempty"`
	RequestedClaims []string               `json:"requested_claims,omitempty"`
	ReturnURL       string                 `json:"return_url,omitempty"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

// IssuanceSession represents an issuance session.
type IssuanceSession struct {
	ID             string                 `json:"id"`
	CredentialType string                 `json:"credential_type"`
	Status         string                 `json:"status"`
	Claims         map[string]interface{} `json:"claims"`
	HostedURL      string                 `json:"hosted_url"`
	CreatedAt      time.Time              `json:"created_at"`
	CompletedAt    *time.Time             `json:"completed_at,omitempty"`
}

// CreateIssuanceParams are the parameters for creating an issuance session.
type CreateIssuanceParams struct {
	CredentialType string                 `json:"credential_type"`
	Claims         map[string]interface{} `json:"claims"`
	ExpiresIn      int                    `json:"expires_in,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

// Credential represents an issued credential.
type Credential struct {
	ID             string     `json:"id"`
	Handle         string     `json:"handle"`
	CredentialType string     `json:"credential_type"`
	Status         string     `json:"status"`
	CreatedAt      time.Time  `json:"created_at"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty"`
}

// Policy represents a verification policy.
type Policy struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	PackType    string                 `json:"pack_type,omitempty"`
	Rules       map[string]interface{} `json:"rules,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
}

// CreatePolicyParams are the parameters for creating a policy.
type CreatePolicyParams struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	PackType    string                 `json:"pack_type,omitempty"`
	Rules       map[string]interface{} `json:"rules,omitempty"`
}

// Issuer represents a trusted credential issuer.
type Issuer struct {
	ID                     string    `json:"id"`
	Name                   string    `json:"name"`
	Domain                 string    `json:"domain,omitempty"`
	AllowedCredentialTypes []string  `json:"allowed_credential_types"`
	AssuranceLevel         string    `json:"assurance_level"`
	Active                 bool      `json:"active"`
	CreatedAt              time.Time `json:"created_at"`
}

// Project represents a project.
type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Environment string    `json:"environment"`
	CreatedAt   time.Time `json:"created_at"`
}

// APIKey represents an API key.
type APIKey struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	KeyPrefix   string     `json:"key_prefix"`
	Key         string     `json:"key,omitempty"` // Only present on creation
	ProjectID   string     `json:"project_id"`
	Environment string     `json:"environment"`
	CreatedAt   time.Time  `json:"created_at"`
	LastUsedAt  *time.Time `json:"last_used_at,omitempty"`
}

// Webhook represents a webhook endpoint.
type Webhook struct {
	ID            string    `json:"id"`
	URL           string    `json:"url"`
	Events        []string  `json:"events"`
	Enabled       bool      `json:"enabled"`
	ProjectID     string    `json:"project_id"`
	SigningSecret string    `json:"signing_secret,omitempty"` // Only present on creation
	CreatedAt     time.Time `json:"created_at"`
}

// TeamMember represents a team member.
type TeamMember struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// Organization represents an organization.
type Organization struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Plan      string    `json:"plan"`
	CreatedAt time.Time `json:"created_at"`
}

// Export represents an evidence export.
type Export struct {
	ID          string                 `json:"id"`
	Status      string                 `json:"status"`
	Filter      map[string]interface{} `json:"filter"`
	ArtifactURL string                 `json:"artifact_url,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
	CompletedAt *time.Time             `json:"completed_at,omitempty"`
}

// ListResponse is a generic list response wrapper.
type ListResponse[T any] struct {
	Data    []T  `json:"data"`
	HasMore bool `json:"has_more,omitempty"`
	Total   int  `json:"total,omitempty"`
}
