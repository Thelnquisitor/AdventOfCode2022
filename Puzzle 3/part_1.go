package main

import (
	"bufio"
	"fmt"
	"os"
)

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

// Get common type found in both compartments in a single rucksack for one elf
func getCommonType(i string, j string) rune {

	for _, iChar := range i {
		for _, jChar := range j {
			if jChar == iChar {
				return iChar
			}
		}
	}

	// Doesn't matter, AOC doesn't seem to consider edge cases
	return -200
}

func offsetAscii(char rune) int {
	if char > 96 {
		return (int(char) - 96)
	}

	return int(char) - 38
}

func main() {
	handle, err := os.Open("input.txt")
	assert(err)
	defer handle.Close()

	scanner := bufio.NewScanner(handle)

	totalPriorityValue := 0

	for scanner.Scan() {
		halfwayPoint := len(scanner.Text()) / 2

		firstCompartmentItems := scanner.Text()[0:halfwayPoint]
		secondCompartmentItems := scanner.Text()[halfwayPoint:]

		commonChar := getCommonType(firstCompartmentItems, secondCompartmentItems)
		totalPriorityValue += offsetAscii(commonChar)
	}

	assert(scanner.Err())
	fmt.Println(totalPriorityValue)
}
