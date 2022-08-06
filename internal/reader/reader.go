package reader

import (
	"bufio"
	"os"
)

func ReadLines(file *os.File) []string {
	res := make([]string, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		res = append(res, scan.Text())
	}

	return res
}
