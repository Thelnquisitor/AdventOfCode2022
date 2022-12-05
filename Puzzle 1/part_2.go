package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	caloriesSlice := []int{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			caloriesSlice = append(caloriesSlice, calories)
			calories = 0
			continue
		}

		i, err := strconv.Atoi(scanner.Text())
		assert(err)
		calories += i
	}
	assert(scanner.Err())

	sort.Ints(caloriesSlice)
	totalCalories := caloriesSlice[len(caloriesSlice)-1] + caloriesSlice[len(caloriesSlice)-2] + caloriesSlice[len(caloriesSlice)-3]
	fmt.Println(totalCalories)
}
