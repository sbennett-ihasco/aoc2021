package common

import (
	"bufio"
	"os"
)

func ReadStrings(path string) []string {
	var lines []string

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
