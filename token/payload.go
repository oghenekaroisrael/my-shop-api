package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("Token has expired")
var ErrInvalidToken = errors.New("Token is invalid")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    int32     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(UserID int32, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		UserID:    UserID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
