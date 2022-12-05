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
	lose = "X"
	draw = "Y"
	win  = "Z"
)

// Ints signify value of shape we choose to satisfy a draw/lose/win
var scenarios = map[string]int{
	"AX": 3,
	"AY": 1,
	"AZ": 2,
	"BX": 1,
	"BY": 2,
	"BZ": 3,
	"CX": 2,
	"CY": 3,
	"CZ": 1,
}

// Calculate score based on if the round was a win, loss or draw and what shape we chose
func calculateRoundScore(opponentChoice string, desiredOutcome string) int {

	// Assume we lost
	roundScore := 0
	shapeScore, _ := scenarios[opponentChoice+desiredOutcome]

	switch desiredOutcome {
	case draw:
		roundScore = 3
	case win:
		roundScore = 6
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
		scenario := strings.Split(scanner.Text(), " ")
		opponentChoice := scenario[0]
		desiredOutcome := scenario[1]

		totalScore += calculateRoundScore(opponentChoice, desiredOutcome)
	}

	assert(scanner.Err())
	fmt.Println(totalScore)
}
