package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
	"strconv"
	"strings"
)

type Board [][]int
type ScoreCard [5][5]int

func main() {
	input := common.ReadStrings("day4.values")

	numbers := parseNumbers(input)
	input = input[1:]

	var boards = make([]Board, 100)
	var scoreCards = make([]ScoreCard, 100)
	k := 0
	for i := range boards {
		boards[i] = make(Board, 5)
		scoreCards[i] = ScoreCard{}
		for j := range boards[i] {
			boards[i][j] = parseBoardNumbers(input[k])
			k++
		}
	}

	fmt.Printf("Value 1: %v\n", play(numbers, boards, scoreCards, true))
	fmt.Printf("Value 2: %v\n", play(numbers, boards, scoreCards, false))
}

func play(numbers []int, boards []Board, scoreCards []ScoreCard, stopOnFirstWin bool) int {
	score := 0
	scoreCount := 0
	var scores []int

	for _, number := range numbers {
	NextBoard:
		for i, board := range boards {
			for j := range board {
				for k := 0; k < 5; k++ {
					if board[j][k] == number {
						scoreCards[i][j][k] = 1
						if win(scoreCards[i]) {
							score = calculateScore(scoreCards[i], board, number)
							if stopOnFirstWin {
								return score
							} else {
								if score == 0 {
									return scores[scoreCount-1]
								}
								scores = append(scores, score)
								scoreCount++
								continue NextBoard
							}
						}
					}
				}
			}
		}
	}
	return score
}

func win(scoreCard ScoreCard) bool {

	for _, rowScores := range scoreCard {
		rowTotal := 0
		for i := 0; i < 5; i++ {
			if rowScores[i] == 1 {
				rowTotal++
			}
			if rowTotal == 5 {
				return true
			}
		}

		columnTotal := make([]int, 5)
		for i := 0; i < 5; i++ {
			for j, colScores := range scoreCard {
				if colScores[j] == 1 {
					columnTotal[i]++
				}
				if columnTotal[i] == 5 {
					return true
				}
			}
		}
	}
	return false
}

func calculateScore(scoreCard ScoreCard, board Board, number int) int {
	total := 0
	for i, score := range scoreCard {
		for j := range score {
			if score[j] == 0 {
				value := board[i][j]
				total += value
			}
		}
	}
	return number * total
}

func parseNumbers(lines []string) []int {
	parts := strings.Split(lines[0], ",")
	numbers := make([]int, len(parts))
	for i, part := range parts {
		integer, _ := strconv.Atoi(part)
		numbers[i] = integer
	}
	return numbers
}

func parseBoardNumbers(line string) []int {
	boardNumbers := make([]int, 5)
	line = strings.Trim(strings.Replace(line, "  ", " ", -1), " ")
	for i, split := range strings.Split(line, " ") {
		imt, _ := strconv.Atoi(split)
		boardNumbers[i] = imt
	}
	return boardNumbers
}
