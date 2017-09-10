package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//ReadLines reads lines from stdin
func ReadLines(path string) []string {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

//Stack type
type Stack []int

func (s *Stack) push(num int) {
	*s = append(*s, num)
}

func (s *Stack) pop() int {
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return el
}

func main() {
	line := ReadLines("./postfix.in")[0]

	operands := new(Stack)

	opExecutors := map[string]func(op1 int, op2 int) int{
		"+": sum,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	var result int

	for _, r := range line {
		//skip space symbol
		if r == 32 {
			continue
		}

		number, err := strconv.Atoi(string(r))

		//if current symbol is not an operand then its operation
		//perform this operation on the last two operands
		if err != nil {
			//pop 2 operands from the stack (indexes swaped)
			operand2 := operands.pop()
			operand1 := operands.pop()

			//execute operation
			operation := string(r)
			number = opExecutors[operation](operand1, operand2)

			result = number
		}

		operands.push(number)
	}

	fmt.Println(result)
}

func sum(op1 int, op2 int) int {
	return op1 + op2
}

func multiply(op1 int, op2 int) int {
	return op1 * op2
}

func subtract(op1 int, op2 int) int {
	return op1 - op2
}

func divide(op1 int, op2 int) int {
	return op1 / op2
}
