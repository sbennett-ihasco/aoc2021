package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
	"strconv"
)

var gLineLength int

func main() {
	values := common.ReadStrings("day3.values")

	gLineLength = len(values[0])

	integers := make([]int, len(values))
	for i, line := range values {
		converted, _ := strconv.ParseInt(line, 2, 64)
		integers[i] = int(converted)
	}

	fmt.Println("https://adventofcode.com/2021/day/3")
	fmt.Printf("Day 3, part 1: %v\n", part1(integers))
	fmt.Printf("Day 3, part 2: %v\n", part2(integers))
}

func part1(integers []int) int {
	ret1 := mostOrLeastCommon(integers, true)
	ret2 := mostOrLeastCommon(integers, false)
	return ret1 * ret2
}

func part2(integers []int) int {
	ret1 := reduceByMostMatches(integers, true)
	ret2 := reduceByMostMatches(integers, false)
	return ret1 * ret2
}

func mostOrLeastCommon(integers []int, matchMost bool) int {
	columnCounts := make([]int, gLineLength)
	for column := 0; column < gLineLength; column++ {
		bitMask := 1 << (gLineLength - column - 1)
		for _, integer := range integers {
			if integer&bitMask > 0 {
				columnCounts[column]++
			} else {
				columnCounts[column]--
			}
		}
	}

	match := 0
	for column, columnCount := range columnCounts {
		bitMask := 1 << (gLineLength - column - 1)
		if (columnCount >= 0) == matchMost {
			match |= bitMask
		}
	}
	return match
}

func reduceByMostMatches(integers []int, matchMost bool) int {
	for i := 0; i < gLineLength; i++ {
		t := mostOrLeastCommon(integers, matchMost)
		integers = sift(integers, t, i)
		if len(integers) == 1 {
			break
		}
	}
	return integers[0]
}

func sift(integers []int, match int, position int) []int {
	var sifted []int
	bitMask := 1 << (gLineLength - position - 1)
	for _, integer := range integers {
		if (integer&bitMask == 0) == (match&bitMask == 0) {
			sifted = append(sifted, integer)
		}
	}
	return sifted
}
