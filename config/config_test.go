package config

import (
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestConfigAddr(t *testing.T) {
	cases := []struct {
		cfg      *Config
		expected string
	}{
		{cfg: &Config{Port: 3000, Host: "127.0.0.1"}, expected: "127.0.0.1:3000"},
	}
	for _, value := range cases {
		tests.AssertEqual(t, value.cfg.Addr(), value.expected)
	}

}
