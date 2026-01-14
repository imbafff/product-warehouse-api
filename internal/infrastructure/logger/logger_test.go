package logger

import (
	"log"
	"testing"
)

func TestNewLogger_ReturnsLogger(t *testing.T) {
	logger := New()

	if logger == nil {
		t.Error("Expected logger, got nil")
	}
}

func TestNewLogger_Type(t *testing.T) {
	logger := New()

	// Check that it's a *log.Logger
	_, ok := interface{}(logger).(*log.Logger)
	if !ok {
		t.Errorf("Expected *log.Logger, got %T", logger)
	}
}

func TestLogger_CanLog(t *testing.T) {
	logger := New()

	// Should not panic when logging
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Logger panicked: %v", r)
		}
	}()

	logger.Println("Test message")
	logger.Printf("Test format: %d\n", 42)
}
