package parser_test

import (
	"testing"

	"github.com/adnanbrq/goenv/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseLines(t *testing.T) {
	expectation := [][]string{
		{"HELLO", "WORLD"},
		{"PORT", "4000"},
		{"METRIC", "4.5"},
		{"DATABASE_URL", "postgresql://test:test@localhost:5432/test"},
		{"DSN", "postgresql://test:test@localhost:5432/test"},
		{"DSN2", "postgresql://test:test@localhost:5432/test"},
		{"WITHOUT_COMMENT", "2345.345#hello"},
		{"WITH_COMMENT", "2345.345"},
		{"A.B.C", "HELLO#WORLD"},
		{"A.B.C.[key].D", "WHAT?"},
	}

	lines := []string{
		// Comments
		"# Comment",

		// Plain values
		"HELLO=WORLD",
		"PORT=4000",
		"METRIC=4.5",

		// Invalid key
		"1239SOMETHING=HELLO",

		// Invalid value
		"SOMETHING=",

		// Spacing
		"",

		// Complex values
		"DATABASE_URL=postgresql://test:test@localhost:5432/test",
		"DSN='postgresql://test:test@localhost:5432/test'",
		"DSN2=\"postgresql://test:test@localhost:5432/test\"",

		// Highlighter treats '#hello' as a comment but it is part of the value. Goenv will not treat '#hello' as a comment
		"WITHOUT_COMMENT=2345.345#hello",
		// here it will treat '#hello' as a comment as it is seperated between the value
		"WITH_COMMENT=2345.345 #hello",

		"A.B.C=HELLO#WORLD",
		"A.B.C.[key].D=WHAT?",
	}
	envs := parser.ParseLines(lines)
	for _, env := range envs {
		assert.Contains(t, expectation, []string{env.Key, env.Value})
	}
	assert.Equal(t, len(expectation), len(envs))
}
