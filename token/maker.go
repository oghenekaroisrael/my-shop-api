package token

import "time"

type Maker interface {
	createToken(UserID int32, duration time.Duration) (string, error)

	verifyToken(token string) (*Payload, error)
}
