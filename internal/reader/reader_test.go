package reader_test

import (
	"os"
	"testing"

	"github.com/adnanbrq/goenv/internal/reader"
	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	file, err := os.Open("../../.env")

	assert.Nil(t, err)
	assert.NotEmpty(t, reader.ReadLines(file))
}
