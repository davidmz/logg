package logg

import (
	"encoding/json"
	"errors"
	"strings"
)

// Level represents logging level.
type Level int

// Log levels
const (
	// ALL is a fake level, for show all events
	ALL Level = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	// OFF is a fake level, for hide all events
	OFF
)

var (
	levelNames = map[Level]string{
		ALL:   "ALL",
		TRACE: "TRACE",
		DEBUG: "DEBUG",
		INFO:  "INFO",
		WARN:  "WARN",
		ERROR: "ERROR",
		FATAL: "FATAL",
		OFF:   "OFF",
	}
)

// ErrUnknownLevelName returning if level name is unknown
var ErrUnknownLevelName = errors.New("unknown level name")

// LevelByName gets log level by its name (case insensitive)
func LevelByName(name string) (Level, error) {
	name = strings.ToUpper(name)
	for l, n := range levelNames {
		if n == name {
			return l, nil
		}
	}
	return OFF, ErrUnknownLevelName
}

// String converts Level value to string
func (l Level) String() string { return levelNames[l] }

// Set obtains Level value from string
func (l *Level) Set(s string) error {
	x, err := LevelByName(s)
	if err != nil {
		*l = x
	}
	return err
}

// MarshalJSON marshal Level value to JSON
func (l Level) MarshalJSON() ([]byte, error) { return json.Marshal(l.String()) }

// UnmarshalJSON unmarshal Level value from JSON
func (l *Level) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return l.Set(s)
}
