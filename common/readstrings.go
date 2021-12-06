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
		t := scanner.Text()
		lines = append(lines, t)
	}

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	return lines
}
