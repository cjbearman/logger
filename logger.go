package logger

import (
	"context"
	"fmt"
)

// Debug emits a log when log level is set to debug
func (l *AppLogger) Debug(msg string) {
	if l.Factory.GetLevel() <= Debug {
		l.Message = msg
		l.Level = "DEBUG"
		l.Factory.writer.Write([]byte(l.Factory.formatter.Format(l)))
	}
}

// Info emits a log when log level is set to debug or info
func (l *AppLogger) Info(msg string) {
	if l.Factory.GetLevel() <= Info {
		l.Message = msg
		l.Level = "INFO"
		l.Factory.writer.Write([]byte(l.Factory.formatter.Format(l)))
	}
}

// Warn emits a log when log level is set to debug, info or warn
func (l *AppLogger) Warn(msg string) {
	if l.Factory.GetLevel() <= Warn {
		l.Message = msg
		l.Level = "WARN"
		l.Factory.writer.Write([]byte(l.Factory.formatter.Format(l)))
	}
}

// Error emits a log level when log level is set to debug, info, warn or error
func (l *AppLogger) Error(msg string) {
	if l.Factory.GetLevel() <= Error {
		l.Message = msg
		l.Level = "ERROR"
		l.Factory.writer.Write([]byte(l.Factory.formatter.Format(l)))
	}
}

// Debugf emits a log when log level is set to debug
func (l *AppLogger) Debugf(format string, v ...interface{}) {
	if l.Factory.GetLevel() <= Debug {
		l.Message = fmt.Sprintf(format, v...)
		l.Level = "DEBUG"
		l.Factory.writer.Write([]byte(l.Factory.formatter.Format(l)))
	}
}

// Infof emits a log when log level is set to debug or info
func (l *AppLogger) Infof(format string, v ...interface{}) {
	if l.Factory.GetLevel() <= Info {
		l.Message = fmt.Sprintf(format, v...)
		l.Level = "INFO"
		l.Factory.writer.Write([]byte(l.Factory.formatter.Format(l)))
	}
}

// Warnf emits a log when log level is set to debug, info or warn
func (l *AppLogger) Warnf(format string, v ...interface{}) {
	if l.Factory.GetLevel() <= Warn {
		l.Message = fmt.Sprintf(format, v...)
		l.Level = "WARN"
		l.Factory.writer.Write([]byte(l.Factory.formatter.Format(l)))
	}
}

// Errorf emits a log level when log level is set to debug, info, warn or error
func (l *AppLogger) Errorf(format string, v ...interface{}) {
	if l.Factory.GetLevel() <= Error {
		l.Message = fmt.Sprintf(format, v...)
		l.Level = "ERROR"
		l.Factory.writer.Write([]byte(l.Factory.formatter.Format(l)))
	}
}

// WithError returns a new logger with an error attached
func (l *AppLogger) WithError(err error) *AppLogger {

	var errMsg = "nil"

	if err != nil {
		errMsg = err.Error()
	}
	return &AppLogger{
		Factory:    l.Factory,
		LoggerName: l.LoggerName,
		Message:    l.Message,
		Level:      l.Level,
		Err:        errMsg,
		Fields:     l.Fields,
	}
}

// WithField returns a new logger with a key/value field attached
func (l *AppLogger) WithField(key string, value any) *AppLogger {

	newFields := make(map[string]any, len(l.Fields)+1)
	for k, v := range l.Fields {
		newFields[k] = v
	}
	newFields[key] = value

	return &AppLogger{
		Factory:    l.Factory,
		LoggerName: l.LoggerName,
		Message:    l.Message,
		Level:      l.Level,
		Err:        l.Err,
		Fields:     newFields,
	}
}

// WithContextFields returns a new logger with any fields found in the context attached
func (l *AppLogger) WithContextFields(ctx context.Context) *AppLogger {
	newFields := make(map[string]any, len(l.Fields)+1)
	for k, v := range l.Fields {
		newFields[k] = v
	}

	if ctxValues, ok := ctx.Value(contextLoggingFieldMapKey).(map[string]any); ok {
		for key, value := range ctxValues {
			newFields[key] = value
		}
	}
	return &AppLogger{
		Factory:    l.Factory,
		LoggerName: l.LoggerName,
		Message:    l.Message,
		Level:      l.Level,
		Err:        l.Err,
		Fields:     newFields,
	}
}

func (l *AppLogger) GetLogger(name string) Logger {
	newFields := make(map[string]any, len(l.Fields)+1)
	for k, v := range l.Fields {
		newFields[k] = v
	}

	return &AppLogger{
		Factory:    l.Factory,
		LoggerName: name,
		Message:    l.Message,
		Level:      l.Level,
		Err:        l.Err,
		Fields:     newFields,
	}

}
