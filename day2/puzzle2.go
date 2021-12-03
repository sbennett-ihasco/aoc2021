package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read(path string) []string {
	var lines []string

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

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
	lines := read("values")

	fmt.Println("https://adventofcode.com/2021/day/2")
	fmt.Printf("Day 2, part 1: %v\n", calculatePath(lines, false))
	fmt.Printf("Day 2, part 2: %v\n", calculatePath(lines, true))
}
