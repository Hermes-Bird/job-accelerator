package repositories

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type RefreshTokenRepository interface {
	SaveToken(token string, id string, duration time.Duration) error
	GetIdByToken(token string) (string, error)
	RemoveToken(token string) error
}

type RefreshTokenRedisRepo struct {
	c *redis.Client
}

var tokenCtx = context.Background()

func NewRefreshTokenRepo(c *redis.Client) RefreshTokenRepository {
	return &RefreshTokenRedisRepo{c}
}

func (r RefreshTokenRedisRepo) GetIdByToken(token string) (string, error) {
	val, err := r.c.Get(tokenCtx, token).Result()

	if err == redis.Nil {
		return "", errors.New("token not found")
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func (r RefreshTokenRedisRepo) RemoveToken(token string) error {
	_, err := r.c.Del(tokenCtx, token).Result()
	return err
}

func (r RefreshTokenRedisRepo) SaveToken(token string, id string, duration time.Duration) error {
	err := r.c.Set(tokenCtx, token, id, duration).Err()
	return err
}
