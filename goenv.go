package goenv

import (
	"errors"
	"fmt"
	"os"

	"github.com/adnanbrq/goenv/internal/parser"
	"github.com/adnanbrq/goenv/internal/reader"
)

func Load(override bool) {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Printf("%v\n%v\n", errors.New("Failed to read .env file."), err)
		return
	}

	lines := reader.ReadLines(file)
	envs := parser.ParseLines(lines)
	for _, env := range envs {
		if env.Exists && !override {
			continue
		}

		os.Setenv(env.Key, env.Value)
	}
}
