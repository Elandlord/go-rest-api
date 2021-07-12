package config

import (
	"testing"
)

func TestConfigSetup(t *testing.T) {
	config := GetConfig()

	if config.DB.Connection == "" {
		t.Fatalf("Config DB Connection not properly set.")
	}

	if config.DB.Host == "" {
		t.Fatalf("Config DB Host not properly set.")
	}

	if config.DB.Port == "" {
		t.Fatalf("Config DB Port not properly set.")
	}

	if config.DB.Username == "" {
		t.Fatalf("Config DB Username not properly set.")
	}

	if config.DB.Password == "" {
		t.Fatalf("Config DB Password not properly set.")
	}

	if config.DB.Charset == "" {
		t.Fatalf("Config DB Charset not properly set.")
	}
}
