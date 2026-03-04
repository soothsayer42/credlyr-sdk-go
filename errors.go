package credlyr

import "fmt"

// APIError is a generic API error.
type APIError struct {
	Code       string
	Message    string
	StatusCode int
	RequestID  string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("credlyr: %s - %s (status %d)", e.Code, e.Message, e.StatusCode)
}

// APIErrorResponse is the error response from the API.
type APIErrorResponse struct {
	Error *struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	} `json:"error"`
}

// AuthenticationError is returned when authentication fails.
type AuthenticationError struct {
	Message   string
	RequestID string
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("credlyr: authentication failed - %s", e.Message)
}

// RateLimitError is returned when rate limit is exceeded.
type RateLimitError struct {
	Message    string
	RetryAfter int
	RequestID  string
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("credlyr: rate limit exceeded, retry after %d seconds", e.RetryAfter)
}

// NotFoundError is returned when a resource is not found.
type NotFoundError struct {
	Message    string
	ResourceID string
	RequestID  string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("credlyr: not found - %s", e.Message)
}

// ValidationError is returned when validation fails.
type ValidationError struct {
	Code      string
	Message   string
	Param     string
	RequestID string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("credlyr: validation error - %s", e.Message)
}
