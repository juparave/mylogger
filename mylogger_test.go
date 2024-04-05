package mylogger_test

import (
	"bytes"
	"testing"

	"github.com/juparave/mylogger"
)

func TestLogging(t *testing.T) {
	// Create buffers to capture output
	var stdOut bytes.Buffer
	var errOut bytes.Buffer

	// Create a new logger with the buffer as output
	logger := mylogger.NewLoggerBuffers(&stdOut, &errOut)

	// Log some messages
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	// Check if the log output contains the expected messages
	expected := "Debug message\nInfo message"
	actual := stdOut.String()
	if expected != actual {
		t.Errorf("Expected log output:\n%s\nActual log output:\n%s", expected, actual)
	}

	// Check if the log output contains the expected messages
	expected = "Warning message\nError message"
	actual = errOut.String()
	if expected != actual {
		t.Errorf("Expected log output:\n%s\nActual log output:\n%s", expected, actual)
	}
}

func TestCustomLogging(t *testing.T) {
	// Create buffers to capture output
	var stdOut bytes.Buffer
	var errOut bytes.Buffer

	// Create a new logger with the buffer as output
	logger := mylogger.NewLoggerBuffers(&stdOut, &errOut)

	// Log some custom messages with key-value pairs
	logger.Info("User logged in", "username", "john_doe", "ip", "192.168.1.1")
	logger.Error("Failed to connect to database", "error", "connection refused")

	// Check if the log output contains the expected messages
	expected := "User logged in username=john_doe ip=192.168.1.1\n"
	actual := stdOut.String()
	if expected != actual {
		t.Errorf("Expected log output:\n%s\nActual log output:\n%s", expected, actual)
	}

	expected = "Failed to connect to database error=connection refused\n"
	actual = errOut.String()
	if expected != actual {
		t.Errorf("Expected log output:\n%s\nActual log output:\n%s", expected, actual)
	}
}

func Example() {
	// Create a new logger
	logger := mylogger.NewLogger()

	// Log some messages
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// Output:
	// This is an informational message
	// This is a warning message
	// This is an error message
}
