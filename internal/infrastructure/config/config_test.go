package config

import (
	"testing"
)

func TestConfigStructure(t *testing.T) {
	// Test that Config struct can be created and has expected fields
	cfg := &Config{
		DBHost: "localhost",
		DBPort: "5432",
		DBUser: "postgres",
		DBPass: "password",
		DBName: "testdb",
		DBSSL:  "disable",
	}

	if cfg.DBHost != "localhost" {
		t.Errorf("Expected DBHost 'localhost', got %s", cfg.DBHost)
	}

	if cfg.DBPort != "5432" {
		t.Errorf("Expected DBPort '5432', got %s", cfg.DBPort)
	}

	if cfg.DBUser != "postgres" {
		t.Errorf("Expected DBUser 'postgres', got %s", cfg.DBUser)
	}

	if cfg.DBPass != "password" {
		t.Errorf("Expected DBPass 'password', got %s", cfg.DBPass)
	}

	if cfg.DBName != "testdb" {
		t.Errorf("Expected DBName 'testdb', got %s", cfg.DBName)
	}

	if cfg.DBSSL != "disable" {
		t.Errorf("Expected DBSSL 'disable', got %s", cfg.DBSSL)
	}
}

func TestConfigEmpty(t *testing.T) {
	// Test that empty Config can be created
	cfg := &Config{}

	if cfg == nil {
		t.Error("Expected config object, got nil")
	}

	// Empty values should be empty strings
	if cfg.DBHost != "" {
		t.Errorf("Expected empty DBHost, got %s", cfg.DBHost)
	}
}
