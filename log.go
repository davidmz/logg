package logg

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

// Logger is a logger type
type Logger struct {
	// MinLevel is a minimal level of messages, accepted by this logger
	MinLevel Level
	// Prefix
	Prefix string

	writers []Writer
}

// Writer writes log events to log
type Writer interface {
	WriteEvent(*Event)
}

// DefaultWriter writes plaintext log to os.Stderr with default template.
var DefaultWriter = NewPlainWriter(os.Stderr, DefaultTpl)

// Default creates new logger with level WARN and writer DefaultWriter
func Default() *Logger { return New(WARN, DefaultWriter) }

// New creates new logger with level and set of writers
func New(level Level, writers ...Writer) *Logger {
	return &Logger{MinLevel: level, writers: writers}
}

var fmtPc = regexp.MustCompile(`(^|[^%])%`)

// Log is a common method to log a formatted message with level
func (l *Logger) Log(level Level, format string, vs ...interface{}) {
	if l == nil || level < l.MinLevel {
		return
	}
	npc := len(fmtPc.FindAllStringIndex(format, -1))
	var payload interface{}
	if len(vs) > npc {
		payload = vs[npc]
		vs = vs[:npc]
	}
	r := &Event{
		Time:    time.Now(),
		Level:   level,
		Prefix:  l.Prefix,
		Message: fmt.Sprintf(format, vs...),
		Payload: payload,
	}
	l.WriteEvent(r)
}

// AddWriter adds writer to logger
func (l *Logger) AddWriter(lw Writer) {
	l.writers = append(l.writers, lw)
}

// RemoveWriter removes writer from logger
func (l *Logger) RemoveWriter(lw Writer) bool {
	foundIdx := -1
	for i, w := range l.writers {
		if w == lw {
			foundIdx = i
			break
		}
	}
	if foundIdx >= 0 {
		l.writers[foundIdx] = nil
		l.writers = append(l.writers[:foundIdx], l.writers[foundIdx+1:]...)
	}
	return foundIdx >= 0
}

// WriteEvent writes Event (from parent logger) to log
func (l *Logger) WriteEvent(r *Event) {
	if r.Level < l.MinLevel {
		return
	}
	for _, w := range l.writers {
		w.WriteEvent(r)
	}
}

// Child creates logger with level ALL and this logger as writer
func (l *Logger) Child() *Logger { return New(ALL, l) }

// ChildWithPrefix creates logger with level ALL, prefix and this logger as writer
func (l *Logger) ChildWithPrefix(prefix string) *Logger {
	ch := l.Child()
	ch.Prefix = prefix
	return ch
}
