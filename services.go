package credlyr

import "context"

// VerificationsService handles verification operations.
type VerificationsService struct {
	client *Client
}

// Create creates a new verification session.
func (s *VerificationsService) Create(ctx context.Context, params *CreateVerificationParams) (*Verification, error) {
	var resp struct {
		Verification *Verification `json:"verification"`
	}
	if err := s.client.post(ctx, "/v1/verifications", params, &resp); err != nil {
		return nil, err
	}
	return resp.Verification, nil
}

// Get retrieves a verification by ID.
func (s *VerificationsService) Get(ctx context.Context, id string) (*Verification, error) {
	var resp struct {
		Verification *Verification `json:"verification"`
	}
	if err := s.client.get(ctx, "/v1/verifications/"+id, &resp); err != nil {
		return nil, err
	}
	return resp.Verification, nil
}

// List lists verifications.
func (s *VerificationsService) List(ctx context.Context) ([]Verification, error) {
	var resp struct {
		Verifications []Verification `json:"verifications"`
	}
	if err := s.client.get(ctx, "/v1/verifications", &resp); err != nil {
		return nil, err
	}
	return resp.Verifications, nil
}

// IssuanceService handles issuance operations.
type IssuanceService struct {
	client *Client
}

// CreateSession creates a new issuance session.
func (s *IssuanceService) CreateSession(ctx context.Context, params *CreateIssuanceParams) (*IssuanceSession, error) {
	var resp struct {
		IssuanceSession *IssuanceSession `json:"issuance_session"`
	}
	if err := s.client.post(ctx, "/v1/issuance_sessions", params, &resp); err != nil {
		return nil, err
	}
	return resp.IssuanceSession, nil
}

// GetSession retrieves an issuance session.
func (s *IssuanceService) GetSession(ctx context.Context, id string) (*IssuanceSession, error) {
	var resp struct {
		IssuanceSession *IssuanceSession `json:"issuance_session"`
	}
	if err := s.client.get(ctx, "/v1/issuance_sessions/"+id, &resp); err != nil {
		return nil, err
	}
	return resp.IssuanceSession, nil
}

// CredentialsService handles credential operations.
type CredentialsService struct {
	client *Client
}

// List lists credentials.
func (s *CredentialsService) List(ctx context.Context) ([]Credential, error) {
	var resp struct {
		Credentials []Credential `json:"credentials"`
	}
	if err := s.client.get(ctx, "/v1/credentials", &resp); err != nil {
		return nil, err
	}
	return resp.Credentials, nil
}

// GetStatus gets credential status by handle.
func (s *CredentialsService) GetStatus(ctx context.Context, handle string) (*Credential, error) {
	var resp struct {
		Credential *Credential `json:"credential"`
	}
	if err := s.client.get(ctx, "/v1/status/"+handle, &resp); err != nil {
		return nil, err
	}
	return resp.Credential, nil
}

// PoliciesService handles policy operations.
type PoliciesService struct {
	client *Client
}

// Create creates a new policy.
func (s *PoliciesService) Create(ctx context.Context, params *CreatePolicyParams) (*Policy, error) {
	var resp struct {
		Policy *Policy `json:"policy"`
	}
	if err := s.client.post(ctx, "/policies", params, &resp); err != nil {
		return nil, err
	}
	return resp.Policy, nil
}

// Get retrieves a policy.
func (s *PoliciesService) Get(ctx context.Context, id string) (*Policy, error) {
	var resp struct {
		Policy *Policy `json:"policy"`
	}
	if err := s.client.get(ctx, "/policies/"+id, &resp); err != nil {
		return nil, err
	}
	return resp.Policy, nil
}

// List lists policies.
func (s *PoliciesService) List(ctx context.Context) ([]Policy, error) {
	var resp struct {
		Policies []Policy `json:"policies"`
	}
	if err := s.client.get(ctx, "/policies", &resp); err != nil {
		return nil, err
	}
	return resp.Policies, nil
}

// Delete deletes a policy.
func (s *PoliciesService) Delete(ctx context.Context, id string) error {
	return s.client.delete(ctx, "/policies/"+id, nil)
}

// IssuersService handles issuer operations.
type IssuersService struct {
	client *Client
}

// List lists trusted issuers.
func (s *IssuersService) List(ctx context.Context) ([]Issuer, error) {
	var resp struct {
		Issuers []Issuer `json:"issuers"`
	}
	if err := s.client.get(ctx, "/trust/issuers", &resp); err != nil {
		return nil, err
	}
	return resp.Issuers, nil
}

// ProjectsService handles project operations.
type ProjectsService struct {
	client *Client
}

// List lists projects.
func (s *ProjectsService) List(ctx context.Context) ([]Project, error) {
	var resp struct {
		Projects []Project `json:"projects"`
	}
	if err := s.client.get(ctx, "/projects", &resp); err != nil {
		return nil, err
	}
	return resp.Projects, nil
}

// APIKeysService handles API key operations.
type APIKeysService struct {
	client *Client
}

// List lists API keys.
func (s *APIKeysService) List(ctx context.Context) ([]APIKey, error) {
	var resp struct {
		APIKeys []APIKey `json:"api_keys"`
	}
	if err := s.client.get(ctx, "/api_keys", &resp); err != nil {
		return nil, err
	}
	return resp.APIKeys, nil
}

// WebhooksService handles webhook operations.
type WebhooksService struct {
	client *Client
}

// List lists webhooks.
func (s *WebhooksService) List(ctx context.Context) ([]Webhook, error) {
	var resp struct {
		Webhooks []Webhook `json:"webhooks"`
	}
	if err := s.client.get(ctx, "/webhooks", &resp); err != nil {
		return nil, err
	}
	return resp.Webhooks, nil
}

// TeamService handles team operations.
type TeamService struct {
	client *Client
}

// List lists team members.
func (s *TeamService) List(ctx context.Context) ([]TeamMember, error) {
	var resp struct {
		Members []TeamMember `json:"members"`
	}
	if err := s.client.get(ctx, "/v1/team", &resp); err != nil {
		return nil, err
	}
	return resp.Members, nil
}

// BillingService handles billing operations.
type BillingService struct {
	client *Client
}

// GetUsage gets usage data.
func (s *BillingService) GetUsage(ctx context.Context) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := s.client.get(ctx, "/v1/billing", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ExportsService handles export operations.
type ExportsService struct {
	client *Client
}

// List lists exports.
func (s *ExportsService) List(ctx context.Context) ([]Export, error) {
	var resp struct {
		Exports []Export `json:"exports"`
	}
	if err := s.client.get(ctx, "/v1/exports", &resp); err != nil {
		return nil, err
	}
	return resp.Exports, nil
}

// OrgService handles organization operations.
type OrgService struct {
	client *Client
}

// Get retrieves organization details.
func (s *OrgService) Get(ctx context.Context) (*Organization, error) {
	var resp struct {
		Org *Organization `json:"org"`
	}
	if err := s.client.get(ctx, "/v1/org", &resp); err != nil {
		return nil, err
	}
	return resp.Org, nil
}
