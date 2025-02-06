package mylogger_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/juparave/mylogger"
)

func removeColorCodes(s string) string {
	// Remove color codes from the string
	for _, c := range []string{"\x1b[2m", "\x1b[92m", "\x1b[93m", "\x1b[91m", "\x1b[0m"} {
		s = strings.ReplaceAll(s, c, "")
	}
	return s
}

func TestLogging(t *testing.T) {
	// Create buffers to capture output
	var stdOut bytes.Buffer
	var errOut bytes.Buffer

	// Create a new logger with the buffer as output
	logger := mylogger.NewLoggerBuffers(&stdOut, &errOut)

	// Log some messages
	logger.Info("Info message")
	logger.Debug("Debug message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	// date and time
	now := time.Now().Format("2006/01/02 15:04:05")

	// Check if the log output contains the expected messages
	expected := fmt.Sprintf("%s INF Info message\n%s DBG Debug message\n", now, now)
	actual := stdOut.String()
	actual = removeColorCodes(actual)
	// remove color codes
	if expected != actual {
		t.Errorf("Expected log output:\n%s\nActual log output:\n%s", expected, actual)
	}

	// Check if the log output contains the expected messages
	expected = fmt.Sprintf("%s WRN mylogger_test.go:32 Warning message\n%s ERR mylogger_test.go:33 Error message\n", now, now)
	actual = errOut.String()
	actual = removeColorCodes(actual)
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

	// date and time
	now := time.Now().Format("2006/01/02 15:04:05")

	// Check if the log output contains the expected messages
	expected := fmt.Sprintf("%s INF User logged in username=john_doe ip=192.168.1.1\n", now)
	actual := stdOut.String()
	actual = removeColorCodes(actual)
	if expected != actual {
		t.Errorf("Expected log output:\n%s\nActual log output:\n%s", expected, actual)
	}

	expected = fmt.Sprintf("%s ERR mylogger_test.go:66 Failed to connect to database error=\"connection refused\"\n", now)
	actual = errOut.String()
	actual = removeColorCodes(actual)
	if expected != actual {
		t.Errorf("Expected log output:\n%s\nActual log output:\n%s", expected, actual)
	}
}

func TestWithStack(t *testing.T) {
	// Create buffers to capture output
	var stdOut bytes.Buffer
	var errOut bytes.Buffer

	// Create a new logger with the buffer as output
	logger := mylogger.NewLoggerBuffers(&stdOut, &errOut)

	// Log some messages with stack trace
	logger.WarnWithStack("Something went wrong")

	// date and time
	now := time.Now().Format("2006/01/02 15:04:05")

	// Check if the log output contains the expected messages
	expected := fmt.Sprintf("%s WRN mylogger_test.go:96 Something went wrong", now)
	actual := errOut.String()
	actual = removeColorCodes(actual)
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("Expected log output:\n%s\nActual log output:\n%s", expected, actual)
	}
}

func example() {
	// Create a new logger
	logger := mylogger.NewLogger()

	// Log some messages
	fmt.Println("\nLogging messages...")
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.WarnWithStack("This is a warning message with stack trace")

	// Output:
	// This is an informational message
	// This is a warning message
	// This is an error message
	// This is a warning message with stack trace
	// goroutine 1 [running]:
	// runtime/debug.Stack(0xc0000a0000, 0x10000, 0x10000)
	// ...
}
