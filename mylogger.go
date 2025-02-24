// Package mylog provides a simple logging utility for Go projects.
package mylogger

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/lmittmann/tint"
)

// Set the LOG_LEVEL environment variable to the desired log level (e.g.,
// DEBUG, INFO, WARN, ERROR) before running your application.

// MyLogger represents a customized logger with separate streams for standard output and errors.
type MyLogger struct {
	stdLogger *slog.Logger // Standard output logger
	errLogger *slog.Logger // Error output logger
	level     slog.Level
}

// Info logs an informational message.
func (l *MyLogger) Info(msg string, keyvals ...interface{}) {
	l.stdLogger.Info(msg, keyvals...)
}

// Debug logs a debug message.
func (l *MyLogger) Debug(msg string, keyvals ...interface{}) {
	l.stdLogger.Debug(msg, keyvals...)
}

// Error logs an error message.
func (l *MyLogger) Error(msg string, keyvals ...interface{}) {
	// Access caller information (requires runtime package)
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		file, line := runtime.FuncForPC(pc).FileLine(pc)
		// ommit the full path of the file
		file = file[strings.LastIndex(file, "/")+1:]
		msg = fmt.Sprintf("%s:%d %s", file, line, msg)
	}
	l.errLogger.Error(msg, keyvals...)
}

// Warn logs a warning message.
func (l *MyLogger) Warn(msg string, keyvals ...interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		file, line := runtime.FuncForPC(pc).FileLine(pc)
		// ommit the full path of the file
		file = file[strings.LastIndex(file, "/")+1:]
		msg = fmt.Sprintf("%s:%d %s", file, line, msg)
	}
	l.errLogger.Warn(msg, keyvals...)
}

// WarnWithStack logs a warning message with stack trace.
func (l *MyLogger) WarnWithStack(msg string, keyvals ...interface{}) {
	buffer := make([]byte, 1<<16)
	stackSize := runtime.Stack(buffer, true)
	stack := string(buffer[:stackSize])
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		file, line := runtime.FuncForPC(pc).FileLine(pc)
		// ommit the full path of the file
		file = file[strings.LastIndex(file, "/")+1:]
		msg = fmt.Sprintf("%s:%d %s", file, line, msg)
	}
	l.errLogger.Warn(msg, keyvals...)

	fmt.Println(stack)
}

// SetLevel sets the log level.
func (l *MyLogger) Enabled(_ context.Context, level slog.Level) bool {
	return level >= l.level.Level()
}

// NewLogger creates a new instance of MyLogger.
func NewLogger() *MyLogger {
	logLevel := os.Getenv("LOG_LEVEL")
	var level slog.Level
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Initialize error logger
	errLogger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      level, // Log warnings and errors
		TimeFormat: "2006/01/02 15:04:05",
	}))

	// Initialize standard output logger
	stdLogger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      level, // Log debug, info
		TimeFormat: "2006/01/02 15:04:05",
	}))

	// Create and return MyLogger instance
	return &MyLogger{
		errLogger: errLogger,
		stdLogger: stdLogger,
		level:     level,
	}
}

// NewLoggerBuffers creates a new instance of MyLogger with the specified output buffers.
// This function is useful for testing purposes.
func NewLoggerBuffers(stdOut, errOut *bytes.Buffer) *MyLogger {
	logLevel := os.Getenv("LOG_LEVEL")
	var level slog.Level
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Initialize error logger
	errLogger := slog.New(tint.NewHandler(errOut, &tint.Options{
		Level:      level, // Log warnings and errors
		TimeFormat: "2006/01/02 15:04:05",
	}))

	// Initialize standard output logger
	stdLogger := slog.New(tint.NewHandler(stdOut, &tint.Options{
		Level:      level, // Log debug, info
		TimeFormat: "2006/01/02 15:04:05",
	}))

	// Create and return MyLogger instance
	return &MyLogger{
		errLogger: errLogger,
		stdLogger: stdLogger,
		level:     level,
	}
}
