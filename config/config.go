package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	DEFAULT_HOST    = "localhost"
	DEFAULT_PORTSTR = "8080"
	DEFAULT_PORT    = 8080
)

type Config struct {
	DatabaseDSN string
	Port        uint
	Host        string
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func FromEnv() (*Config, error) {
	host, ok := os.LookupEnv("GO_HOST")
	if !ok {
		host = DEFAULT_HOST
	}

	portStr, ok := os.LookupEnv("GO_PORT")
	if !ok {
		portStr = DEFAULT_PORTSTR
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 0 {
		port = DEFAULT_PORT
	}

	databaseDsn := os.Getenv("GO_DATABASE_DSN")

	cfg := Config{Host: host, Port: uint(port), DatabaseDSN: databaseDsn}

	return &cfg, nil
}
