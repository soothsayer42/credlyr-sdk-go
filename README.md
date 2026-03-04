# Credlyr Go SDK

Official Go SDK for the Credlyr API - verifiable credentials infrastructure for identity verification.

## Installation

```bash
go get github.com/credlyr/sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    credlyr "github.com/credlyr/sdk-go"
)

func main() {
    client := credlyr.NewClient("sk_live_xxx")

    verification, err := client.Verifications.Create(context.Background(), &credlyr.CreateVerificationParams{
        RequestedClaims: []string{"name", "email"},
        ReturnURL:       "https://example.com/callback",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Verification URL:", verification.HostedURL)
}
```

## Configuration

```go
client := credlyr.NewClient("sk_live_xxx",
    credlyr.WithBaseURL("https://custom.api.com"),
    credlyr.WithTimeout(60 * time.Second),
    credlyr.WithMaxRetries(5),
)
```

## Resources

### Verifications

```go
// Create
verification, err := client.Verifications.Create(ctx, &credlyr.CreateVerificationParams{
    PolicyID:        "pol_xxx",
    RequestedClaims: []string{"name", "email"},
    ReturnURL:       "https://example.com/callback",
})

// Get
verification, err := client.Verifications.Get(ctx, "ver_xxx")

// List
verifications, err := client.Verifications.List(ctx)
```

### Issuance

```go
session, err := client.Issuance.CreateSession(ctx, &credlyr.CreateIssuanceParams{
    CredentialType: "EmploymentCredential",
    Claims: map[string]interface{}{
        "employer": "Acme Corp",
        "position": "Engineer",
    },
})
```

### Policies

```go
policy, err := client.Policies.Create(ctx, &credlyr.CreatePolicyParams{
    Name: "Age Verification",
    Rules: map[string]interface{}{"required_claims": []string{"birthDate"}},
})
```

## Webhook Verification

```go
func handleWebhook(w http.ResponseWriter, r *http.Request) {
    payload, _ := io.ReadAll(r.Body)
    signature := r.Header.Get("X-Credlyr-Signature")

    if !credlyr.VerifyWebhookSignature(payload, signature, webhookSecret) {
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }

    event, _ := credlyr.ParseWebhookEvent(payload)
    switch event.Type {
    case "verification.completed":
        // Handle
    }

    w.WriteHeader(http.StatusOK)
}
```

## Error Handling

```go
verification, err := client.Verifications.Get(ctx, "ver_xxx")
if err != nil {
    switch e := err.(type) {
    case *credlyr.AuthenticationError:
        log.Fatal("Invalid API key")
    case *credlyr.NotFoundError:
        log.Printf("Not found: %s", e.Message)
    case *credlyr.RateLimitError:
        log.Printf("Rate limited, retry after %d seconds", e.RetryAfter)
    case *credlyr.ValidationError:
        log.Printf("Validation error: %s", e.Message)
    }
}
```

## Requirements

- Go 1.21+
- No external dependencies

## License

MIT
