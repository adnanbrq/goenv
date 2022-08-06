package parser_test

import (
	"testing"

	"github.com/adnanbrq/goenv/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseLines(t *testing.T) {
	expectedKeys := []string{
		"HELLO",
		"PORT",
		"METRIC",
		"DATABASE_URL",
		"DSN",
		"WITH_COMMENT",
		"HELLO",
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
		"WITH_COMMENT=2345.345#hello",
		"HELLO=\"WORLD\"",
	}
	envs := parser.ParseLines(lines)
	for _, env := range envs {
		assert.Contains(t, expectedKeys, env.Key)
	}
	assert.Equal(t, len(expectedKeys), len(envs))
}
