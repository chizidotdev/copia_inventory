package token_manager

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return nil, nil
}
func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return nil, nil
}
func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}
func (p *Payload) GetIssuer() (string, error) {
	return "", nil
}
func (p *Payload) GetSubject() (string, error) {
	return "", nil
}
func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}
