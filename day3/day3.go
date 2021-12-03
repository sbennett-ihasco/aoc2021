package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
	"sort"
	"strconv"
	"strings"
)

func p1(values []string) int64 {

	ones := make(map[int]int)
	var gamma string
	var epsilon string
	threshold := len(values) / 2

	for _, value := range values {
		parts := strings.Split(value, "")
		for i, bit := range parts {
			if bit == "1" {
				ones[i]++
			}
		}
	}

	keys := make([]int, 0)
	for i, _ := range ones {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, i := range keys {
		if ones[i] > threshold {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	ret1, _ := strconv.ParseInt(gamma, 2, 64)
	ret2, _ := strconv.ParseInt(epsilon, 2, 64)

	return ret1 * ret2
}

func p2(values []string) int {
	x := 0
	return x
}

func main() {
	lines := common.ReadStrings("day3.values")

	fmt.Println("https://adventofcode.com/2021/day/3")
	fmt.Printf("Day 3, part 1: %v\n", p1(lines))
	fmt.Printf("Day 3, part 2: %v\n", p2(lines))
}
