package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	Port           string
	AccessSecret   string
	RefreshSecret  string
	AccessTimeout  time.Duration
	RefreshTimeout time.Duration
	DbUsername     string
	DbPassword     string
	DbAddress      string
	RedisAddr      string
	RedisPassword  string
	RedisDb        int
}

func GetConfig() Config {
	var cfg = Config{
		Port:           ":8080",
		AccessSecret:   "access-secret",
		RefreshSecret:  "refresh-secret",
		AccessTimeout:  15 * time.Minute,
		RefreshTimeout: 30 * 24 * time.Hour,
		DbAddress:      "127.0.0.1:3306",
		DbUsername:     "root",
		DbPassword:     "",
		RedisAddr:      "localhost:6379",
		RedisPassword:  "",
		RedisDb:        0,
	}

	if err := env.Parse(&cfg); err != nil {
		fmt.Println("failed:", err)
	}
	return cfg
}
