// Package mylog provides a simple logging utility for Go projects.
package mylogger

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/lmittmann/tint"
)

// MyLogger represents a customized logger with separate streams for standard output and errors.
type MyLogger struct {
	stdLogger *slog.Logger // Standard output logger
	errLogger *slog.Logger // Error output logger
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

// NewLogger creates a new instance of MyLogger.
func NewLogger() *MyLogger {
	// Initialize error logger
	errLogger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelWarn,
		TimeFormat: "2006/01/02 15:04:05",
		// AddSource:  true,
		// ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
		// 	fmt.Println(groups)
		// 	if attr.Key == slog.SourceKey {
		// 		// Access caller information (requires runtime package)
		// 		pc, _, _, ok := runtime.Caller(2)
		// 		if ok {
		// 			file, line := runtime.FuncForPC(pc).FileLine(pc)
		// 			attr = slog.String("s", fmt.Sprintf("%s:%d", file, line))
		// 		}
		// 	}
		// 	return attr
		// },
		// NoColor: true,
	}))

	// Initialize standard output logger
	stdLogger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug, // Log debug, info
		TimeFormat: "2006/01/02 15:04:05",
	}))

	// Create and return MyLogger instance
	return &MyLogger{
		errLogger: errLogger,
		stdLogger: stdLogger,
	}
}

// NewLoggerBuffers creates a new instance of MyLogger with the specified output buffers.
// This function is useful for testing purposes.
func NewLoggerBuffers(stdOut, errOut *bytes.Buffer) *MyLogger {
	// Initialize error logger
	errLogger := slog.New(tint.NewHandler(errOut, &tint.Options{
		Level:      slog.LevelWarn, // Log warnings and errors
		TimeFormat: "2006/01/02 15:04:05",
		// AddSource:  true, // Include source file and line number
		// ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
		// 	fmt.Println(groups)
		// 	if attr.Key == slog.SourceKey {
		// 		// Access caller information (requires runtime package)
		// 		pc, _, _, ok := runtime.Caller(4)
		// 		if ok {
		// 			file, line := runtime.FuncForPC(pc).FileLine(pc)
		// 			attr = slog.String("s", fmt.Sprintf("%s:%d", file, line))
		// 		}
		// 	}
		// 	return attr
		// },
	}))

	// Initialize standard output logger
	stdLogger := slog.New(tint.NewHandler(stdOut, &tint.Options{
		Level:      slog.LevelDebug, // Log debug, info, warnings, and errors
		TimeFormat: "2006/01/02 15:04:05",
	}))

	// Create and return MyLogger instance
	return &MyLogger{
		errLogger: errLogger,
		stdLogger: stdLogger,
	}
}
