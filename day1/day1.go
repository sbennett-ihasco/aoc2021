package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
)

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
	lines := common.ReadIntegers("day1.values")

	fmt.Println("https://adventofcode.com/2021/day/1")
	fmt.Printf("Day 1, part 1: %v\n", calculateIncrements(lines, 1))
	fmt.Printf("Day 1, part 2: %v\n", calculateIncrements(lines, 3))
}
