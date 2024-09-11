// Package logging provides a robust logging system for the Albatross project.
package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger initializes the global logger with custom settings.
func InitLogger() {
	// Set up the logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	// Set the global log level (can be changed based on environment)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// Fields is a type alias for log field key-value pairs
type Fields map[string]interface{}

// Info logs an info level message with optional fields
func Info(message string, fields Fields) {
	event := log.Info()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}

// Error logs an error level message with optional fields
func Error(message string, err error, fields Fields) {
	event := log.Error().Err(err)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}

// Debug logs a debug level message with optional fields
func Debug(message string, fields Fields) {
	event := log.Debug()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}

// Fatal logs a fatal level message with optional fields and then exits
func Fatal(message string, fields Fields) {
	event := log.Fatal()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}
