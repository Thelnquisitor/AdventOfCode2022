package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func createSequenceFromStringRange(stringRange []string) []int {
	low, _ := strconv.Atoi(stringRange[0])
	high, _ := strconv.Atoi(stringRange[1])

	sequenceArray := make([]int, high-low+1)
	for i := range sequenceArray {
		sequenceArray[i] = i + low
	}

	return sequenceArray
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func overlaps(x string, y string) bool {
	xSplit := strings.Split(x, "-")
	ySplit := strings.Split(y, "-")

	xSequence := createSequenceFromStringRange(xSplit)
	ySequence := createSequenceFromStringRange(ySplit)

	// Does x overlap in y?
	for _, i := range xSequence {
		if contains(ySequence, i) {
			return true
		}
	}

	// Does y overlap in x?
	for _, i := range ySequence {
		if contains(xSequence, i) {
			return true
		}
	}

	return false
}

func main() {
	handle, err := os.Open("input.txt")
	assert(err)
	defer handle.Close()

	scanner := bufio.NewScanner(handle)

	totalCount := 0

	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")

		if overlaps(assignments[0], assignments[1]) {
			totalCount++
		}
	}

	assert(scanner.Err())
	fmt.Println(totalCount)
}
