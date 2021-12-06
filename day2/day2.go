package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
	"strconv"
	"strings"
)

const Forward = "forward"
const Up = "up"
const Down = "down"

func main() {
	input := common.ReadStrings("day2.values")
	fmt.Println(calculatePath(input, false))
	fmt.Println(calculatePath(input, true))
}

func calculatePath(values []string, withAim bool) int {
	x := 0
	y := 0
	aim := 0

	for _, command := range values {
		parts := strings.Split(command, " ")
		direction := fmt.Sprint(parts[0])
		steps, _ := strconv.Atoi(parts[1])

		if withAim {
			switch direction {
			case Forward:
				x += steps
				y += aim * steps
			case Up:
				aim -= steps
			case Down:
				aim += steps
			}
		} else {
			switch direction {
			case Forward:
				x += steps
			case Up:
				y -= steps
			case Down:
				y += steps
			}
		}
	}
	return x * y
}
