package logg_test

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/davidmz/logg"
)

type testDataRec struct {
	Prefix  string
	Message string
	Payload interface{}
	Re      *regexp.Regexp
}

var plainStringRe = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2} \[[A-Z]+\]( [^\x00-\x1f]+)?: [^\x00-\x1f]*\n`)

var testData = []testDataRec{
	testDataRec{"", "Hello,\nworld!", nil, plainStringRe},
	testDataRec{"hi", "Hello,\nworld!", 123, plainStringRe},
	testDataRec{" ", "", nil, plainStringRe},
}

func TestPlain(t *testing.T) {
	for _, dat := range testData {
		wr := &bytes.Buffer{}
		l := logg.Default()
		slave := logg.New(logg.ALL, logg.NewPlainWriter(wr, logg.DefaultTpl))

		l.AddWriter(slave)
		l.Log(logg.WARN, dat.Message, dat.Payload)
		res := wr.String()
		if !dat.Re.MatchString(res) {
			t.Errorf("log returns %q, expected %q", res, dat.Re.String())
		}
	}
}
