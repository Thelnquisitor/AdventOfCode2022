package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	data []string
}

func (s *stack) pop() string {
	top := len(s.data) - 1
	value := s.data[top]
	s.data = s.data[:top]
	return value
}

func (s *stack) push(val string) {
	s.data = append(s.data, val)
}

func (s stack) peek() string {
	top := len(s.data) - 1
	return s.data[top]
}

var stacks []stack

func assert(e error) {
	if e != nil {
		panic(e)
	}
}

func initStacks() {
	stack1 := stack{data: []string{"G", "T", "R", "W"}}
	stack2 := stack{data: []string{"G", "C", "H", "P", "M", "S", "V", "W"}}
	stack3 := stack{data: []string{"C", "L", "T", "S", "G", "M"}}
	stack4 := stack{data: []string{"J", "H", "D", "M", "W", "R", "F"}}
	stack5 := stack{data: []string{"P", "Q", "L", "H", "S", "W", "F", "J"}}
	stack6 := stack{data: []string{"P", "J", "D", "N", "F", "M", "S"}}
	stack7 := stack{data: []string{"Z", "B", "D", "F", "G", "C", "S", "J"}}
	stack8 := stack{data: []string{"R", "T", "B"}}
	stack9 := stack{data: []string{"H", "N", "W", "L", "C"}}

	stacks = []stack{stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9}
}

func dumpStackTops() {
	fmt.Println(stacks[0].peek())
	fmt.Println(stacks[1].peek())
	fmt.Println(stacks[2].peek())
	fmt.Println(stacks[3].peek())
	fmt.Println(stacks[4].peek())
	fmt.Println(stacks[5].peek())
	fmt.Println(stacks[6].peek())
	fmt.Println(stacks[7].peek())
	fmt.Println(stacks[8].peek())
}

func sanitizeProcedure(proc string) string {
	buff := strings.ReplaceAll(proc, "move", "")
	buff = strings.ReplaceAll(buff, "from", "")
	buff = strings.ReplaceAll(buff, "to", "")
	buff = strings.ReplaceAll(buff, "  ", " ")
	buff = strings.TrimSpace(buff)
	return buff
}

func rearrange(proc string) {
	instructions := strings.Split(proc, " ")
	popNum, _ := strconv.Atoi(instructions[0])
	source, _ := strconv.Atoi(instructions[1])
	target, _ := strconv.Atoi(instructions[2])

	// aligning with array indexes.
	source--
	target--

	for i := 0; i < popNum; i++ {
		crate := stacks[source].pop()
		stacks[target].push(crate)
		fmt.Println("Moved", crate, "from Stack", source, "to Stack", target)
	}
}

func main() {
	handle, err := os.Open("input.txt")
	assert(err)
	defer handle.Close()

	scanner := bufio.NewScanner(handle)

	initStacks()

	count := 0
	for scanner.Scan() {
		count++
		if count < 11 {
			continue
		}

		currentProcedure := sanitizeProcedure(scanner.Text())
		rearrange(currentProcedure)
	}

	assert(scanner.Err())
	dumpStackTops()
}
