package services

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"github.com/Hermes-Bird/job-accelerator/internal/repositories"
	"github.com/golang-jwt/jwt"
	"time"
)

type AuthService interface {
	GenerateTokenPair(id string, userType string, accessDuration time.Duration, refreshDuration time.Duration) (*domain.TokensPayload, error)
	RefreshToken(token string, accessDuration time.Duration, refreshDuration time.Duration) (*domain.TokensPayload, error)
	ValidateAccessToken(token string, userType string) (*domain.JwtPayload, error)
	HashPassword(password string) (*domain.PasswordPayload, error)
	HashPasswordWithSalt(password string, salt string) (string, error)
}

type AuthServiceImpl struct {
	tokenStore    repositories.RefreshTokenRepository
	accessSecret  string
	refreshSecret string
}

func (s AuthServiceImpl) HashPassword(password string) (*domain.PasswordPayload, error) {
	salt := make([]byte, 16)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	passwordBytes := []byte(password)
	passwordSaltBytes := append(passwordBytes, salt...)

	hash := sha512.Sum512(passwordSaltBytes)

	return &domain.PasswordPayload{
		PasswordHash: base64.StdEncoding.EncodeToString(hash[:]),
		PasswordSalt: base64.StdEncoding.EncodeToString(salt),
	}, nil
}

func (s AuthServiceImpl) HashPasswordWithSalt(password string, salt string) (string, error) {
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", fmt.Errorf("error while decoding: %s", err.Error())
	}
	passwordSaltBytes := append([]byte(password), saltBytes...)

	hash := sha512.Sum512(passwordSaltBytes)
	return base64.StdEncoding.EncodeToString(hash[:]), nil
}

func NewJwtServiceImpl(accessSecret string, refreshToken string, ts repositories.RefreshTokenRepository) AuthService {
	return &AuthServiceImpl{
		accessSecret:  accessSecret,
		refreshSecret: refreshToken,
		tokenStore:    ts,
	}
}

func (s AuthServiceImpl) GenerateAccessToken(id string, userType string, duration time.Duration) (string, error) {
	payload := domain.NewPayload(id, userType, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(s.accessSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s AuthServiceImpl) ValidateAccessToken(token string, userType string) (*domain.JwtPayload, error) {
	var payload domain.JwtPayload
	jwtToken, err := jwt.ParseWithClaims(token, &payload, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token alg: %s", token.Header["alg"])
		}
		return []byte(s.accessSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("invalid token provided")
	}

	if err = payload.Valid(); err != nil {
		return nil, err
	}

	if payload.UserType != userType {
		return nil, errors.New(fmt.Sprintf("wrong token type: %s", payload.UserType))
	}

	return &payload, nil
}

func (s AuthServiceImpl) GenerateTokenPair(id string, userType string, accessDuration time.Duration, refreshDuration time.Duration) (*domain.TokensPayload, error) {
	accessToken, err := s.GenerateAccessToken(id, userType, accessDuration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.GenerateRefreshToken(id, userType, refreshDuration)
	if err != nil {
		return nil, err
	}

	return &domain.TokensPayload{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s AuthServiceImpl) GenerateRefreshToken(id string, userType string, duration time.Duration) (string, error) {
	payload := domain.NewPayload(id, userType, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := jwtToken.SignedString([]byte(s.refreshSecret))
	if err != nil {
		return "", err
	}

	err = s.tokenStore.SaveToken(token, id, duration)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s AuthServiceImpl) RefreshToken(token string, accessDuration time.Duration, refreshDuration time.Duration) (*domain.TokensPayload, error) {
	var payload domain.JwtPayload
	_, err := jwt.ParseWithClaims(token, &payload, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token alg: %s", jwtToken.Header["alg"])
		}
		return []byte(s.refreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	id, err := s.tokenStore.GetIdByToken(token)
	if err != nil || id == "" {
		return nil, errors.New("refresh token not valid")
	}

	pair, err := s.GenerateTokenPair(payload.Id, payload.UserType, accessDuration, refreshDuration)
	if err != nil {
		return nil, err
	}

	err = s.tokenStore.RemoveToken(token)
	if err != nil {
		return nil, err
	}

	return pair, nil
}
