package logg

import (
	"encoding/json"
	"time"
)

// Event represents a single log event
type Event struct {
	Time    time.Time
	Level   Level
	Prefix  string
	Message string
	Payload interface{}
}

// TimeRFC3339 formattes event date by time.RFC3339
func (r *Event) TimeRFC3339() string {
	return r.Time.Format(time.RFC3339)
}

// MarshalJSON marshals Event to JSON
func (r *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TimeStr string      `json:"time"`
		Level   Level       `json:"level"`
		Prefix  string      `json:"prefix,omitempty"`
		Message string      `json:"message"`
		Payload interface{} `json:"payload,omitempty"`
	}{
		r.TimeRFC3339(),
		r.Level,
		r.Prefix,
		r.Message,
		r.Payload,
	})
}
