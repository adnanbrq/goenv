package goenv_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/adnanbrq/goenv"
	"github.com/stretchr/testify/assert"
)

func Test_With_Open_File_Error(t *testing.T) {
	// Rename file to artificially force errors
	assert.Nil(t, os.Rename(".env", ".env.temp"))

	// Listen to Stdout
	scan := bufio.NewScanner(os.Stdout)

	// Load .env
	goenv.Load(false)

	// Assert that a message has been printed to the console
	for scan.Scan() {
		assert.NotEmpty(t, scan.Text())
	}

	// Restore old filename
	assert.Nil(t, os.Rename(".env.temp", ".env"))
}

func Test_Load_With_Override(t *testing.T) {

	os.Setenv("PORT", "3000")
	goenv.Load(true)

	assert.Equal(t, "4000", os.Getenv("PORT"))
}

func Test_Load_Without_Override(t *testing.T) {
	// Listen to Stdout
	// scan := bufio.NewScanner(os.Stdout)

	os.Setenv("PORT", "3000")
	goenv.Load(false)

	assert.Equal(t, "3000", os.Getenv("PORT"))
}
