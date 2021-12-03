package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"strconv"
)

func readFileToIntegers(path string) []int {
	var lines []int

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}

	return lines
}

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

func p1(values []string) int {
	x := 0
	return x
}

func p2(values []string) int {
	x := 0
	return x
}

func main() {
	lines := read("values")

	fmt.Println("https://adventofcode.com/2021/day/3")
	fmt.Printf("Day 3, part 1: %v\n", p1(lines))
	fmt.Printf("Day 3, part 2: %v\n", p2(lines))
}
