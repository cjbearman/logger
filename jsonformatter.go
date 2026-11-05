package logger

import (
	"encoding/json"
)

type jsonFormatter struct{}

// NewJsonFormatter returns a formatter that outputs logs in JSON format
func NewJsonFormatter() LogFormatter {
	return &jsonFormatter{}
}

// Format formats the log
func (f *jsonFormatter) Format(log *AppLogger) string {

	data, err := json.Marshal(log)
	if err != nil {
		return "failed to format log" // should not happen
	}
	return string(data) + "\n"

}
