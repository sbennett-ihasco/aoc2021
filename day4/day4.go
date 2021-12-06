package main

import (
	"fmt"
	"sbennett-ihasco/aoc2021/common"
	"strconv"
	"strings"
)

const CardSize = 5
const BoardCount = 100

type Board [][]int
type ScoreCard [CardSize][CardSize]int

func main() {
	input := common.ReadStrings("input.txt")

	numbers := parseNumbers(input)
	input = input[1:]

	var boards = make([]Board, BoardCount)
	var scoreCards = make([]ScoreCard, BoardCount)
	k := 0
	for i := range boards {
		boards[i] = make(Board, CardSize)
		scoreCards[i] = ScoreCard{}
		for j := range boards[i] {
			boards[i][j] = parseBoardNumbers(input[k])
			k++
		}
	}

	fmt.Println(play(numbers, boards, scoreCards, true), play(numbers, boards, scoreCards, false))
}

func play(numbers []int, boards []Board, scoreCards []ScoreCard, stopOnFirstWin bool) (score int) {
	var scoreCount int
	var scores []int

	for _, number := range numbers {
		for i, board := range boards {
		NextBoard:
			for j := range board {
				for k := 0; k < CardSize; k++ {
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

func win(scoreCard ScoreCard) (won bool) {
	for _, rowScores := range scoreCard {
		rowTotal := 0
		for i := 0; i < CardSize; i++ {
			if rowScores[i] == 1 {
				rowTotal++
			}
			if rowTotal == CardSize {
				won = true
			}
		}

		columnTotal := make([]int, CardSize)
		for i := 0; i < CardSize; i++ {
			for j, colScores := range scoreCard {
				if colScores[j] == 1 {
					columnTotal[i]++
				}
				if columnTotal[i] == CardSize {
					won = true
				}
			}
		}
	}
	return won
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

func parseNumbers(lines []string) (numbers []int) {
	parts := strings.Split(lines[0], ",")
	numbers = make([]int, len(parts))
	for i, part := range parts {
		integer, _ := strconv.Atoi(part)
		numbers[i] = integer
	}
	return numbers
}

func parseBoardNumbers(line string) (boardNumbers []int) {
	boardNumbers = make([]int, CardSize)
	line = strings.Trim(strings.Replace(line, "  ", " ", -1), " ")
	for i, split := range strings.Split(line, " ") {
		integer, _ := strconv.Atoi(split)
		boardNumbers[i] = integer
	}
	return boardNumbers
}
