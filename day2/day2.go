package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
	"strconv"
	"strings"
)

func main() {
	input := common.ReadStrings("input.txt")
	fmt.Println(calculatePath(input, false), calculatePath(input, true))
}

func calculatePath(values []string, withAim bool) int {
	var x, y, aim int

	for _, command := range values {
		parts := strings.Split(command, " ")
		direction := fmt.Sprint(parts[0])
		steps, _ := strconv.Atoi(parts[1])
		if withAim {
			switch direction {
			case "forward":
				x += steps
				y += aim * steps
			case "up":
				aim -= steps
			case "down":
				aim += steps
			}
		} else {
			switch direction {
			case "forward":
				x += steps
			case "up":
				y -= steps
			case "down":
				y += steps
			}
		}
	}
	return x * y
}
