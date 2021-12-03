package day2

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
	"strconv"
	"strings"
)

func calculatePath(values []string, withAim bool) int {
	x := 0
	y := 0
	aim := 0

	for _, command := range values {
		parts := strings.Split(command, " ")
		direction := fmt.Sprint(parts[0])
		steps, _ := strconv.Atoi(parts[1])

		if withAim == false {

			if direction == "forward" {
				x += steps
			} else if direction == "up" {
				y -= steps
			} else if direction == "down" {
				y += steps
			}

		} else {

			if direction == "forward" {
				x += steps
				y += aim * steps
			} else if direction == "up" {
				aim -= steps
			} else if direction == "down" {
				aim += steps
			}
		}
	}

	return x * y
}

func main() {
	lines := common.ReadStrings("day2.values")

	fmt.Println("https://adventofcode.com/2021/day/1")
	fmt.Printf("Day 2, part 1: %v\n", calculatePath(lines, false))
	fmt.Printf("Day 1, part 2: %v\n", calculatePath(lines, true))
}
