package token

import "time"

// We are implementing this interface so that we can easily switch between different token implementations
type Maker interface {
	// MakeToken creates a JWT token with the given username and expiration duration
	MakeToken(username string, duration time.Duration) (string, error)

	// VerifyToken verifies the token string and returns the payload
	VerifyToken(token string) (*Payload, error)
}
