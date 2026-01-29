// Package logger provides structured logging functionality for all microservices.
// It supports multiple log levels, file output, and structured logging with fields.
// The logger interface allows for easy testing and implementation swapping.
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Logger defines the interface for logging operations.
// This interface allows for easy mocking in tests and swapping implementations.
type Logger interface {
	// Info logs informational messages with optional fields
	Info(msg string, fields ...any)
	// Debug logs debug messages with optional fields (only if log level is debug)
	Debug(msg string, fields ...any)
	// Warn logs warning messages with optional fields
	Warn(msg string, fields ...any)
	// Error logs error messages with optional fields
	Error(msg string, fields ...any)
}

// logger is the concrete implementation of the Logger interface.
// It uses the standard log package with different loggers for each level.
type logger struct {
	out      io.Writer   // Output destination (file or stdout)
	infoLog  *log.Logger // Info level logger
	debugLog *log.Logger // Debug level logger (optional)
	warnLog  *log.Logger // Warning level logger
	errorLog *log.Logger // Error level logger
}

// NewLogger creates a new Logger instance with specified log level and output file.
// If filePath is provided, logs are written to file, otherwise to stdout.
// Debug logs are only enabled if logLevel is "debug".
//
// Parameters:
//   - logLevel: Logging level - debug, info, warn, error (default: "info")
//   - filePath: Path to log file (optional, logs to stdout if empty)
//
// Returns:
//   - Logger: Configured logger instance
//   - error: File open error or invalid log level error
//
// Example:
//
//	logger, err := NewLogger("debug", "./logs/app.log")
//	if err != nil {
//	    log.Fatal("Failed to create logger:", err)
//	}
func NewLogger(logLevel, filePath string) (Logger, error) {
	var out io.Writer = os.Stdout

	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		out = file
	}

	if logLevel == "" {
		logLevel = "info"
	}
	logLevel = strings.ToLower(logLevel)

	validLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	if !validLevels[logLevel] {
		return nil, fmt.Errorf("invalid log level: %s", logLevel)
	}

	flag := log.Ldate | log.Ltime | log.Lshortfile

	l := &logger{
		out:      out,
		infoLog:  log.New(out, "INFO\t", flag),
		warnLog:  log.New(out, "WARN\t", flag),
		errorLog: log.New(out, "ERROR\t", flag),
	}

	if logLevel == "debug" {
		l.debugLog = log.New(out, "DEBUG\t", flag)
	}

	return l, nil
}

// Info logs informational messages with optional fields.
// Fields are passed as key-value pairs and will be formatted by the logger.
//
// Example:
//
//	logger.Info("Server started", "port", 8080, "mode", "production")
func (l logger) Info(msg string, fields ...any) {
	l.infoLog.Printf(msg, fields...)
}

// Debug logs debug messages with optional fields.
// Only logs if debug level is enabled during logger creation.
// Safe to call even if debug logging is disabled.
//
// Example:
//
//	logger.Debug("Processing request", "id", reqID, "method", "GET")
func (l logger) Debug(msg string, fields ...any) {
	if l.debugLog != nil {
		l.debugLog.Printf(msg, fields...)
	}
}

// Warn logs warning messages with optional fields.
// Used for potentially problematic situations that don't prevent the application from running.
//
// Example:
//
//	logger.Warn("Slow query detected", "duration", time.Since(start), "query", sql)
func (l logger) Warn(msg string, fields ...any) {
	l.warnLog.Printf(msg, fields...)
}

// Error logs error messages with optional fields.
// Used for errors that should be investigated but don't necessarily crash the application.
//
// Example:
//
//	logger.Error("Database connection failed", "error", err, "retry", attempt)
func (l logger) Error(msg string, fields ...any) {
	l.errorLog.Printf(msg, fields...)
}
