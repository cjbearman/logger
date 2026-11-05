package logger

import (
	"fmt"
	"strings"
)

// LevelFromString will return a log level from it's string representations
func LevelFromString(logLevel string) (LogLevel, error) {

	switch strings.ToLower(logLevel) {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn", "warning":
		return Warn, nil
	case "error", "err":
		return Error, nil
	case "none":
		return None, nil
	default:
		return None, fmt.Errorf("invalid log level: %s", logLevel)
	}
}

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case None:
		return "none"
	default:
		return "unknown"
	}
}
