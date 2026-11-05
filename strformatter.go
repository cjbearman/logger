package logger

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type colorMode int

const (
	CMNone colorMode = 0
	CM8    colorMode = 8
	CM256  colorMode = 256
)

type stringFormatter struct {
	color bool
	mode  colorMode
}

// NewStringFormatter returns a basic string formatter
func NewStringFormatter() LogFormatter {
	return &stringFormatter{}
}

// NewColorStringFormatter returns a colorized string formatter
// with support for either 8 or 256 bit color terminals
// and backup support for uncolored terminals
func NewColorStringFormatter() LogFormatter {

	// Because we're enabling colors, we have to try and figure out
	// how many colors, if any, we can support.
	// We use terminfo via infocmp and a backup of evaluating TERM variable
	// Assume we have no colors available as the default option
	mode := CMNone

	// Use infocmp -c to get information about the terminal
	terminfo, err := exec.Command("infocmp", "-c").Output()
	if err == nil {
		rdr := bufio.NewScanner(bytes.NewReader(terminfo))

		// GO through output line by line
		re := regexp.MustCompile(`colors\s*=\s*([0-9]+)`)
		for rdr.Scan() {
			line := rdr.Text()
			trimmed := strings.TrimSpace(string(line))
			matches := re.FindSubmatch([]byte(trimmed))

			if len(matches) > 1 {
				// I.E. colors= xxx., this is what we want
				if ncolors, err := strconv.Atoi(string(matches[1])); err == nil {
					if ncolors >= 256 {
						mode = CM256
					} else if ncolors >= 8 {
						mode = CM8
					}
				}
				break
			}
		}
	}

	if mode == CMNone {
		// Haven't figured it out yet, check TERM directly
		term := strings.ToLower(os.Getenv("TERM"))
		if strings.Contains(term, "256") {
			mode = CM256
		} else if strings.Contains(term, "color") {
			mode = CM8
		}

	}

	return &stringFormatter{color: true, mode: mode}
}

// Format formats the log
func (f *stringFormatter) Format(log *AppLogger) string {
	var bldr strings.Builder

	if f.color {
		bldr.WriteString(f.colorLevel(log.Level))
	} else {
		bldr.WriteString(log.Level)
	}
	bldr.WriteRune(' ')
	bldr.WriteString(log.LoggerName)
	bldr.WriteString(" - ")
	bldr.WriteString(log.Message)

	if log.Err != "" {
		bldr.WriteString(": ")
		bldr.WriteString(log.Err)
	}

	if len(log.Fields) > 0 {
		if f.color && f.mode != CMNone {
			if f.mode == CM256 {
				bldr.WriteString(" \x1b[38;5;16;48;5;159m[")
			} else {
				bldr.WriteString(" \x1b[46;37m[")
			}
		} else {
			bldr.WriteString(" [")
		}
		first := true

		fieldKeys := make([]string, 0, len(log.Fields))
		for fk := range log.Fields {
			fieldKeys = append(fieldKeys, fk)
		}
		sort.Strings(fieldKeys)

		for _, k := range fieldKeys {
			if !first {
				bldr.WriteString(", ")
			}
			first = false
			bldr.WriteString(k)
			bldr.WriteString(" = ")
			bldr.WriteString(fmt.Sprintf("%v", log.Fields[k]))
		}
		bldr.WriteRune(']')
		if f.color {
			bldr.WriteString("\x1b[0m")
		}
	}

	bldr.WriteRune('\n')

	return bldr.String()
}

func (s *stringFormatter) colorLevel(l string) string {
	switch s.mode {
	case CM256:
		switch l {
		case "DEBUG":
			return "\x1b[38;5;16;48;5;145m DEBUG \x1b[0m"
		case "INFO":
			return "\x1b[38;5;16;48;5;150m INFO  \x1b[0m"
		case "WARN":
			return "\x1b[38;5;16;48;5;214m WARN  \x1b[0m"
		case "ERROR":
			return "\x1b[38;5;231;48;5;196m ERROR \x1b[0m"
		default:
			return "\x1b[38;5;16;48;5;116m" + l + "\x1b[0m"
		}
	case CM8:
		switch l {
		case "DEBUG":
			return "\x1b[30;47m DEBUG \x1b[0m"
		case "INFO":
			return "\x1b[37;42m INFO  \x1b[0m"
		case "WARN":
			return "\x1b[30;43m WARN  \x1b[0m"
		case "ERROR":
			return "\x1b[37;41m ERROR \x1b[0m"
		default:
			return "\x1b[37;45m" + l + "\x1b[0m"
		}

	default:
		return l
	}
}
