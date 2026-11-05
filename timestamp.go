package logger

import "time"

type timestamper struct {
	format string
	tz     *time.Location
}

func (t *timestamper) String() string {
	return time.Now().In(t.tz).Format(t.format)
}

func NewTimestamper(format string, tz *time.Location) *timestamper {
	return &timestamper{
		format: format,
		tz:     tz,
	}
}
