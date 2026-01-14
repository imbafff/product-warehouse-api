package db

import (
	"testing"

	"github.com/imbafff/product-warehouse-api/internal/infrastructure/config"
)

func TestNewPostgresDB_InvalidConfig(t *testing.T) {
	cfg := &config.Config{
		DBHost: "invalid-host-that-does-not-exist",
		DBPort: "5432",
		DBUser: "postgres",
		DBPass: "wrongpass",
		DBName: "nonexistent",
		DBSSL:  "disable",
	}

	db, err := NewPostgresDB(cfg)

	// Should fail to connect to invalid host
	if err == nil {
		t.Error("Expected error for invalid database config, got nil")
	}

	if db != nil {
		t.Error("Expected nil db connection for invalid config")
	}
}

func TestNewPostgresDB_ValidConfigReturnNonNil(t *testing.T) {
	// This test would need a real DB running
	// Skipping as we cannot guarantee DB availability in tests
	t.Skip("Skipping DB connection test - requires running PostgreSQL instance")
}

func TestDSN_Construction(t *testing.T) {
	cfg := &config.Config{
		DBHost: "localhost",
		DBPort: "5432",
		DBUser: "postgres",
		DBPass: "postgres",
		DBName: "testdb",
		DBSSL:  "disable",
	}

	// Test that config can be passed to NewPostgresDB without panic
	// Even if connection fails, config parsing should work
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("NewPostgresDB panicked: %v", r)
		}
	}()

	// Call function (it will fail to connect, but that's okay)
	_ , _ = NewPostgresDB(cfg)
}
