package token

import "time"

// Maker is an interface to manage tokens
type Maker interface {
	// create  a new token for a specific username and valid duration
	CreateToken(username string, duration time.Duration) (string, error)
	// check if the inputed token is valid or not
	VerifyToken(token string) (*Payload, error)
}
