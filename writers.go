package logg

import (
	"encoding/json"
	"io"
	"sync"
	"text/template"
)

type simpleWriter struct {
	io.Writer
	Lock      sync.Mutex
	WriteFunc func(io.Writer, *Event)
}

func (b *simpleWriter) WriteEvent(r *Event) {
	b.Lock.Lock()
	defer b.Lock.Unlock()
	b.WriteFunc(b, r)
}

// Default template for plaintext writer.
const DefaultTpl = `{{.TimeRFC3339}} [{{.Level}}]{{ if .Prefix }} {{.Prefix}}{{ end }}: {{.Message}}`

// NewPlainWriter creates writer for plaintext output
func NewPlainWriter(w io.Writer, tpl string) Writer {
	if tpl[len(tpl)-1:] != "\n" {
		tpl = tpl + "\n"
	}
	templ := template.Must(template.New("PlainWriter").Parse(tpl))
	return &simpleWriter{
		Writer:    w,
		WriteFunc: func(w io.Writer, r *Event) { templ.Execute(w, r) },
	}
}

// NewJSONWriter creates writer for JSON output
func NewJSONWriter(w io.Writer) Writer {
	return &simpleWriter{
		Writer:    w,
		WriteFunc: func(w io.Writer, r *Event) { json.NewEncoder(w).Encode(r) },
	}
}
