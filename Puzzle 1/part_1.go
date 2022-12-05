package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	calories := 0
	maxCalories := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
			continue
		}

		i, err := strconv.Atoi(scanner.Text())
		assert(err)
		calories += i
	}

	assert(scanner.Err())
	fmt.Println(maxCalories)
}
