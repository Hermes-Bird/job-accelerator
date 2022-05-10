package domain

import (
	"github.com/golang-jwt/jwt"
	"time"
)

const CompanyType = "company"
const EmployeeType = "employee"

type JwtPayload struct {
	jwt.StandardClaims
	Id        string    `json:"id"`
	UserType  string    `json:"user_type"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expired_at"`
}

func NewPayload(id string, userType string, duration time.Duration) *JwtPayload {
	return &JwtPayload{
		Id:        id,
		UserType:  userType,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
}

type PasswordPayload struct {
	PasswordHash string
	PasswordSalt string
}

type TokensPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
