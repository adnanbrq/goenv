package parser

import (
	"os"
	"regexp"
	"strings"
)

var (
	regExKey, _          = regexp.Compile("^[a-zA-Z]{1}(\\w+|\\.|\\[[\\w\\-]+\\])+$")
	regExValueBase, _    = regexp.Compile("^(\\w|\\S)+")
	regExValueSingleQ, _ = regexp.Compile("^'(\\w+|\\S+)'")
	regExValueDoubleQ, _ = regexp.Compile("^\"(\\w+|\\S+)\"")
)

type Env struct {
	Exists bool
	Key    string
	Value  string
}

func ParseLines(lines []string) []Env {
	res := make([]Env, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) < 2 {
			continue
		}

		key := parts[0]
		value := strings.Join(parts[1:], "=")

		if !regExKey.MatchString(key) {
			continue
		}

		switch true {
		case regExValueSingleQ.MatchString(value):
			{
				if groups := regExValueSingleQ.FindStringSubmatch(value); len(groups) > 1 {
					value = groups[1]
				}
				break
			}
		case regExValueDoubleQ.MatchString(value):
			{
				if groups := regExValueDoubleQ.FindStringSubmatch(value); len(groups) > 1 {
					value = groups[1]
				}
				break
			}
		case regExValueBase.MatchString(value):
			{
				value = regExValueBase.FindString(value)
				break
			}
		default:
			{
				value = ""
			}
		}

		if value != "" {
			_, exists := os.LookupEnv(key)
			res = append(res, Env{
				Exists: exists,
				Key:    key,
				Value:  value,
			})
		}
	}

	return res
}
