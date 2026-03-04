package credlyr

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

// DefaultWebhookTolerance is the default time tolerance for webhook timestamps.
const DefaultWebhookTolerance = 5 * time.Minute

// WebhookEvent represents a webhook event payload.
type WebhookEvent struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
	CreatedAt time.Time              `json:"created_at"`
}

// VerifyWebhookSignature verifies a webhook signature.
func VerifyWebhookSignature(payload []byte, signature, secret string) bool {
	return VerifyWebhookSignatureWithTolerance(payload, signature, "", secret, 0)
}

// VerifyWebhookSignatureWithTimestamp verifies a webhook signature with timestamp.
func VerifyWebhookSignatureWithTimestamp(payload []byte, signature, timestamp, secret string) bool {
	return VerifyWebhookSignatureWithTolerance(payload, signature, timestamp, secret, DefaultWebhookTolerance)
}

// VerifyWebhookSignatureWithTolerance verifies a webhook signature with custom tolerance.
func VerifyWebhookSignatureWithTolerance(payload []byte, signature, timestamp, secret string, tolerance time.Duration) bool {
	// Parse signature header (format: "t=timestamp,v1=signature")
	parts := make(map[string]string)
	for _, part := range strings.Split(signature, ",") {
		if kv := strings.SplitN(part, "=", 2); len(kv) == 2 {
			parts[kv[0]] = kv[1]
		}
	}

	ts := parts["t"]
	if ts == "" {
		ts = timestamp
	}
	sig := parts["v1"]
	if sig == "" {
		sig = signature
	}

	// Check timestamp tolerance
	if ts != "" && tolerance > 0 {
		tsInt, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			return false
		}
		if time.Since(time.Unix(tsInt, 0)) > tolerance {
			return false
		}
	}

	// Compute expected signature
	var message []byte
	if ts != "" {
		message = []byte(ts + "." + string(payload))
	} else {
		message = payload
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(message)
	expectedSig := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(sig), []byte(expectedSig))
}

// ParseWebhookEvent parses a webhook payload after verification.
func ParseWebhookEvent(payload []byte) (*WebhookEvent, error) {
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, errors.New("failed to parse webhook event")
	}
	return &event, nil
}

// ComputeWebhookSignature computes a webhook signature for testing.
func ComputeWebhookSignature(payload []byte, timestamp, secret string) string {
	var message []byte
	if timestamp != "" {
		message = []byte(timestamp + "." + string(payload))
	} else {
		message = payload
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(message)
	sig := hex.EncodeToString(mac.Sum(nil))

	if timestamp != "" {
		return "t=" + timestamp + ",v1=" + sig
	}
	return "v1=" + sig
}
