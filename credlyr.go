// Package credlyr provides the official Go SDK for the Credlyr API.
package credlyr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultBaseURL    = "https://api.credlyr.com"
	DefaultTimeout    = 30 * time.Second
	DefaultMaxRetries = 2
)

// Client is the Credlyr API client.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	maxRetries int

	// Services
	Verifications *VerificationsService
	Issuance      *IssuanceService
	Credentials   *CredentialsService
	Policies      *PoliciesService
	Issuers       *IssuersService
	Projects      *ProjectsService
	APIKeys       *APIKeysService
	Webhooks      *WebhooksService
	Team          *TeamService
	Billing       *BillingService
	Exports       *ExportsService
	Org           *OrgService
}

// ClientOption is a function that configures a Client.
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithMaxRetries sets the maximum number of retries.
func WithMaxRetries(retries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = retries
	}
}

// NewClient creates a new Credlyr API client.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		apiKey:     apiKey,
		baseURL:    DefaultBaseURL,
		httpClient: &http.Client{Timeout: DefaultTimeout},
		maxRetries: DefaultMaxRetries,
	}

	for _, opt := range opts {
		opt(c)
	}

	// Initialize services
	c.Verifications = &VerificationsService{client: c}
	c.Issuance = &IssuanceService{client: c}
	c.Credentials = &CredentialsService{client: c}
	c.Policies = &PoliciesService{client: c}
	c.Issuers = &IssuersService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.APIKeys = &APIKeysService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	c.Team = &TeamService{client: c}
	c.Billing = &BillingService{client: c}
	c.Exports = &ExportsService{client: c}
	c.Org = &OrgService{client: c}

	return c
}

// request makes an HTTP request to the API.
func (c *Client) request(ctx context.Context, method, path string, body, result interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Authorization", "Bearer "+c.apiKey)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "credlyr-go/1.0.0")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				time.Sleep(time.Duration(500*(1<<attempt)) * time.Millisecond)
				continue
			}
			return fmt.Errorf("request failed: %w", err)
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = err
			continue
		}

		if resp.StatusCode == 429 {
			if attempt < c.maxRetries {
				time.Sleep(time.Second)
				continue
			}
			return &RateLimitError{Message: "rate limit exceeded"}
		}

		if resp.StatusCode >= 500 && attempt < c.maxRetries {
			time.Sleep(time.Duration(500*(1<<attempt)) * time.Millisecond)
			continue
		}

		if resp.StatusCode >= 400 {
			var apiErr APIErrorResponse
			if err := json.Unmarshal(respBody, &apiErr); err == nil && apiErr.Error != nil {
				switch resp.StatusCode {
				case 400:
					return &ValidationError{Code: apiErr.Error.Code, Message: apiErr.Error.Message}
				case 401:
					return &AuthenticationError{Message: apiErr.Error.Message}
				case 404:
					return &NotFoundError{Message: apiErr.Error.Message}
				default:
					return &APIError{Code: apiErr.Error.Code, Message: apiErr.Error.Message, StatusCode: resp.StatusCode}
				}
			}
			return &APIError{Code: "unknown", Message: string(respBody), StatusCode: resp.StatusCode}
		}

		if result != nil && len(respBody) > 0 {
			if err := json.Unmarshal(respBody, result); err != nil {
				return fmt.Errorf("failed to unmarshal response: %w", err)
			}
		}

		return nil
	}

	if lastErr != nil {
		return lastErr
	}
	return fmt.Errorf("request failed after %d attempts", c.maxRetries+1)
}

func (c *Client) get(ctx context.Context, path string, result interface{}) error {
	return c.request(ctx, http.MethodGet, path, nil, result)
}

func (c *Client) post(ctx context.Context, path string, body, result interface{}) error {
	return c.request(ctx, http.MethodPost, path, body, result)
}

func (c *Client) patch(ctx context.Context, path string, body, result interface{}) error {
	return c.request(ctx, http.MethodPatch, path, body, result)
}

func (c *Client) delete(ctx context.Context, path string, result interface{}) error {
	return c.request(ctx, http.MethodDelete, path, nil, result)
}
