package logger

import (
	"io"
	"sync"
)

// LoggerFactory is used to configure logger output
// and obtain loggers
// The log level set on the factory is used by all loggers that it generates
// and can be changed at will
type AppLoggerFactory struct {
	level     LogLevel
	levelLock sync.RWMutex
	writer    io.Writer
	formatter LogFormatter
}

// NewLoggerFactory generates a new logger factory, with an initial log level
// a formatter and a writer
func NewLoggerFactory(level LogLevel, formatter LogFormatter, writer io.Writer) *AppLoggerFactory {
	return &AppLoggerFactory{
		level:     level,
		formatter: formatter,
		writer:    writer,
	}
}

// SetLevel changes the log level for all loggers derived from this factory
func (l *AppLoggerFactory) SetLevel(level LogLevel) {
	l.levelLock.Lock()
	defer l.levelLock.Unlock()
	l.level = level
}

// GetLevel returns the current log level
func (l *AppLoggerFactory) GetLevel() LogLevel {
	l.levelLock.RLock()
	defer l.levelLock.RUnlock()
	return l.level
}

// GetLogger generates a new named logger
func (f *AppLoggerFactory) GetLogger(name string) Logger {
	return &AppLogger{
		Factory:    f,
		LoggerName: name,
	}
}
