package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
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
	TestDbName  string
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

func (c *Config) Copy() Config {
	newConfig := *c

	return newConfig
}

type TestConfig struct {
	Config Config
	DbName string
}

type GetTestDbNameRes struct {
	DbName string
	Dsn    string
}

func GetTestDbName(dsn string) (res GetTestDbNameRes, err error) {
	if dsn == "" {
		return res, fmt.Errorf("database DSN is empty")
	}
	dbUrl, err := url.Parse(dsn)
	if err != nil {
		return res, err
	}
	newPath := strings.Replace(dbUrl.Path, "/", "/test_", 1)
	dbUrl.Path = newPath
	res.DbName = strings.ReplaceAll(newPath, "/", "")

	res.Dsn = dbUrl.String()
	return res, nil
}

func (c *Config) TestConfig() (cfg *TestConfig, err error) {
	newConfig := c.Copy()
	newDbName, err := GetTestDbName(c.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	newConfig.DatabaseDSN = newDbName.Dsn
	cfg = &TestConfig{Config: newConfig, DbName: newDbName.DbName}
	return cfg, nil
}
