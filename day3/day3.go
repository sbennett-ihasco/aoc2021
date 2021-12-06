package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
	"strconv"
)

var bits int

func main() {
	input := common.ReadStrings("day3.values")

	bits = len(input[0])

	integers := make([]int, len(input))
	for i, value := range input {
		converted, _ := strconv.ParseInt(value, 2, 64)
		integers[i] = int(converted)
	}

	fmt.Println(part1(integers))
	fmt.Println(part2(integers))
}

func part1(integers []int) int {
	most := mostOrLeastCommon(integers, true)
	least := mostOrLeastCommon(integers, false)
	return most * least
}

func part2(integers []int) int {
	most := reduceByMostOrLeastCommon(integers, true)
	least := reduceByMostOrLeastCommon(integers, false)
	return most * least
}

func mostOrLeastCommon(integers []int, mostCommon bool) (match int) {
	columnCounts := make([]int, bits)

	for bit := 0; bit < bits; bit++ {
		// Shift mask by 1 each time (from rhs), reducing its length
		mask := 1 << (bits - bit - 1)
		for _, integer := range integers {
			// Masking << [n] accumulating by column [0] >>
			if integer&mask > 0 {
				columnCounts[bit]++
			} else {
				columnCounts[bit]--
			}
		}
	}

	match = 0
	for bit, columnCount := range columnCounts {
		mask := 1 << (bits - bit - 1)
		if (columnCount >= 0) == mostCommon {
			match |= mask
		}
	}

	return match
}

func reduceByMostOrLeastCommon(integers []int, mostCommon bool) int {
	for column := 0; column < bits; column++ {
		matches := mostOrLeastCommon(integers, mostCommon)
		integers = reduce(integers, matches, column)
		if len(integers) == 1 {
			break
		}
	}
	return integers[0]
}

func reduce(integers []int, match int, position int) (reduced []int) {
	mask := 1 << (bits - position - 1)
	for _, integer := range integers {
		if (integer&mask == 0) == (match&mask == 0) {
			reduced = append(reduced, integer)
		}
	}
	return reduced
}
