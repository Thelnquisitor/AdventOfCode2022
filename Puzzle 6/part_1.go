package main

import (
	"bufio"
	"fmt"
	"os"
)

const packetLength = 14

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func hasDuplicate(sequence string) bool {
	fmt.Println(sequence)
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
	for i := packetLength; i < len(data); i += 1 {
		low := i - packetLength
		if !hasDuplicate(string(data[low:i])) {
			totalCount += i
			break
		}
	}

	fmt.Println(totalCount)
}
