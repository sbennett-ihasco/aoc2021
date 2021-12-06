package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
)

func main() {
	input := common.ReadIntegers("day1.values")
	fmt.Println(calculateIncrements(input, 1))
	fmt.Println(calculateIncrements(input, 3))
}

func calculateIncrements(depths []int, offset int) (increments int) {
	for i, current := range depths[offset:] {
		if current > depths[i] {
			increments++
		}
	}
	return
}
