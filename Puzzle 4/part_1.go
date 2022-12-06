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

func containsRange(sequence []int, potentialSubset []int) bool {

	for _, i := range potentialSubset {
		if !contains(sequence, i) {
			return false
		}
	}

	return true
}

// Determine whether assignment x is contained fully by assignment y or vice versa
func isFullyContained(x string, y string) bool {
	xSplit := strings.Split(x, "-")
	ySplit := strings.Split(y, "-")

	// transform into range of numbers?
	xSequence := createSequenceFromStringRange(xSplit)
	ySequence := createSequenceFromStringRange(ySplit)

	if containsRange(xSequence, ySequence) || containsRange(ySequence, xSequence) {
		return true
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

		if isFullyContained(assignments[0], assignments[1]) {
			totalCount++
		}
	}

	assert(scanner.Err())
	fmt.Println(totalCount)
}
