package config

import "os"

type PG_CONFIG struct {
    Host string
    Port string
    User string
    Password string
    Database string
    Schema string
}

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func Init() PG_CONFIG {
    return PG_CONFIG {
        Host: getenv("POSTGRES_HOST", "localhost"),
        Port: getenv("POSTGRES_PORT", "5432"),
        User: getenv("POSTGRES_USER", "postgres"),
        Password: os.Getenv("POSTGRES_PASSWORD"),
        Database: getenv("POSTGRES_DB", "postgres"),
        Schema: getenv("POSTGRES_SCHEMA", "public"),
    }
}
