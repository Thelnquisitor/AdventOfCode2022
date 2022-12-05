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

// Get common type found in a group of elves' rucksacks
func getCommonType(elfGroup []string) rune {

	firstElf := elfGroup[0]
	secondElf := elfGroup[1]
	thirdElf := elfGroup[2]

	for _, firstType := range firstElf {
		for _, secondType := range secondElf {
			for _, thirdType := range thirdElf {
				if thirdType == secondType && secondType == firstType {
					return thirdType
				}
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

	count := 0
	var processBuffer []string

	// Process three lines at a time
	for scanner.Scan() {
		count++
		processBuffer = append(processBuffer, scanner.Text())

		if count%3 == 0 {
			groupValue := getCommonType(processBuffer)
			processBuffer = []string{}
			totalPriorityValue += offsetAscii(groupValue)
		}
	}

	assert(scanner.Err())
	fmt.Println(totalPriorityValue)
}
