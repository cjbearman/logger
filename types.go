package logger

import (
	"context"
)

type LogLevel int

type contextKeyFields string

const (
	Debug LogLevel = 0
	Info  LogLevel = 1
	Warn  LogLevel = 2
	Error LogLevel = 3
	None  LogLevel = 4

	contextLoggingFieldMapKey contextKeyFields = "logger_fields"
)

// LogFormatter may be defined by any entity that wishes to encode logs into a string for output
type LogFormatter interface {
	// Format formats the provided log and returns the desired string form
	Format(log *AppLogger) string
}

// Logger is obtained by LoggerFactory and is used to log messages
// as well as to obtain derivative loggers containing additional fields
type AppLogger struct {
	// Factory is the logger factory, do not change this directly
	Factory *AppLoggerFactory `json:"-"`

	// LoggerName is the logger name, do not change this directly
	LoggerName string `json:"logger"`

	// Message is the message about to be logged, do not change this directly
	Message string `json:"msg"`

	// Level is the log level of the message being logged, do not change this directly
	Level string `json:"level"`

	// Err is the error set for a message about to be logged, do not change this directly
	Err string `json:"error,omitempty"`

	// Fields are the fields configured for the logger, do not change this directly
	Fields map[string]any `json:"fields,omitempty"`
}

type LoggerFactory interface {
	SetLevel(level LogLevel)
	GetLevel() LogLevel
	GetLogger(name string) Logger
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	WithError(err error) *AppLogger
	WithField(key string, value any) *AppLogger
	WithContextFields(ctx context.Context) *AppLogger
}
