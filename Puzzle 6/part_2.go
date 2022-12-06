package main

import (
	"bufio"
	"fmt"
	"os"
)

const messageLength = 14

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func hasDuplicate(sequence string) bool {
	seen := map[rune]struct{}{}

	for _, char := range sequence {
		_, exists := seen[char]
		if !exists {
			seen[char] = struct{}{}
			continue
		}

		return true
	}

	return false
}

func main() {
	handle, err := os.Open("input.txt")
	assert(err)
	defer handle.Close()

	scanner := bufio.NewScanner(handle)
	scanner.Scan()
	assert(scanner.Err())
	data := scanner.Bytes()

	totalCount := 0
	for i := messageLength; i < len(data); i += 1 {
		low := i - messageLength
		if !hasDuplicate(string(data[low:i])) {
			totalCount += i
			break
		}
	}

	fmt.Println(totalCount)
}
