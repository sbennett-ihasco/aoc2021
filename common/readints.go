package common

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntegers(path string) []int {
	var lines []int

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}

	return lines
}
