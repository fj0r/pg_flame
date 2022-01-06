package config

import (
	"fmt"
	"os"
)

type PG_CONFIG struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Schema   string
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func Init() PG_CONFIG {
	return PG_CONFIG{
		Host:     Getenv("POSTGRES_HOST", "localhost"),
		Port:     Getenv("POSTGRES_PORT", "5432"),
		User:     Getenv("POSTGRES_USER", "postgres"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: Getenv("POSTGRES_DB", "postgres"),
		Schema:   Getenv("POSTGRES_SCHEMA", "public"),
	}
}

func (c *PG_CONFIG) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.Database)
}
