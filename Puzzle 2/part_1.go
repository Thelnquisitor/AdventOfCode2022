package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// First column of input
	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"

	// Second Column of input
	rock     = "X"
	paper    = "Y"
	scissors = "Z"
)

// From our perspective
var scenarios = map[string]string{
	"AX": "Draw",
	"AY": "Win",
	"AZ": "Loss",
	"BX": "Loss",
	"BY": "Draw",
	"BZ": "Win",
	"CX": "Win",
	"CY": "Loss",
	"CZ": "Draw",
}

// Calculate score based on if the round was a win, loss or draw and what shape we chose
func calculateRoundScore(outcome string, shape string) int {

	// Assume we lost
	roundScore := 0

	switch outcome {
	case "Draw":
		roundScore = 3
	case "Win":
		roundScore = 6
	}

	shapeScore := 0
	switch shape {
	case rock:
		shapeScore = 1
	case paper:
		shapeScore = 2
	case scissors:
		shapeScore = 3
	}

	return roundScore + shapeScore
}

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	handle, err := os.Open("input.txt")
	assert(err)
	defer handle.Close()

	scanner := bufio.NewScanner(handle)

	totalScore := 0

	for scanner.Scan() {
		scenario := strings.ReplaceAll(scanner.Text(), " ", "")

		outcome, _ := scenarios[scenario]

		totalScore += calculateRoundScore(outcome, scenario[len(scenario)-1:])
	}

	assert(scanner.Err())
	fmt.Println(totalScore)
}
