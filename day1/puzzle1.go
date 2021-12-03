package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func read(path string) []int {
	var lines []int

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}

	return lines
}

func calculateIncrements(values []int, offset int) int {
	increments := 0

	for i, current := range values[offset:] {
		if current > values[i] {
			increments++
		}
	}

	return increments
}

func main() {
	lines := read("values")
	
	fmt.Println("https://adventofcode.com/2021/day/1")
	fmt.Printf("Day 1, part 1: %v\n", calculateIncrements(lines, 1))
	fmt.Printf("Day 1, part 2: %v\n", calculateIncrements(lines, 3))
}
