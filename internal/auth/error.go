package auth

import "fmt"

// CreateTokenError contains reason of token creation failure.
type CreateTokenError struct {
	Cause  error
	Reason string
}

// Error implements error interface.
func (e *CreateTokenError) Error() string {
	return fmt.Sprintf("metadata: create token error: %s", e.Reason)
}

func (e *CreateTokenError) Unwrap() error {
	return e.Cause
}
