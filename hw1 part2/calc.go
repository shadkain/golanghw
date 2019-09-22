package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	defer handleError()

	input := readInput()
	regExp := createRegExp()

	symbols := extractSymbols(input, regExp)

	exp := prepareExpression(symbols)
	result := evaluateExpression(exp)

	fmt.Println(result)
}

// Function declarations
func handleError() {
	err := recover()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func readInput() string {
	args := os.Args

	nArgs := len(args)
	if nArgs != 2 {
		panic(fmt.Sprintf("expected 1 argument, but provided %d", nArgs-1))
	}

	return args[1]
}

func createRegExp() *regexp.Regexp {
	return regexp.MustCompile(`\d+|[()+/*-]`)
}

func extractSymbols(source string, regExp *regexp.Regexp) []string {
	return regExp.FindAllString(source, -1)
}

func prepareExpression(symbols []string) []string {
	output := make([]string, 0, len(symbols))
	stack := createStack(len(symbols))

	isNumber := createIsNumberFunc()

	for _, sym := range symbols {
		if isNumber(sym) {
			output = append(output, sym)
		} else {
			if isCloseBrace(sym) {
				for top := stack.Top(); !isOpenBrace(top); top = stack.Top() {
					output = append(output, top)
					stack.Pop()
				}
				stack.Pop()
			} else if isOpenBrace(sym) {
				stack.Push(sym)
			} else {
				if !stack.isEmpty() {
					top := stack.Top()
					if !isOpenBrace(top) && priority(top) > priority(sym) {
						output = append(output, top)
						stack.Pop()
					}
				}

				stack.Push(sym)
			}
		}
	}
	for !stack.isEmpty() {
		output = append(output, stack.Top())
		stack.Pop()
	}

	return output
}

func createIsNumberFunc() func(string) bool {
	regExp := regexp.MustCompile(`\d+`)

	return func(str string) bool {
		return regExp.FindString(str) != ""
	}
}

func isOpenBrace(str string) bool {
	return str == "("
}

func isCloseBrace(str string) bool {
	return str == ")"
}

func priority(str string) (priority int) {
	switch str {
	case "+":
		fallthrough
	case "-":
		priority = 1
	case "*":
		fallthrough
	case "/":
		priority = 2
	}

	return
}

func evaluateExpression(exp []string) int {
	stack := createIntStack(len(exp))
	isNumber := createIsNumberFunc()

	for _, sym := range exp {
		if isNumber(sym) {
			num, err := strconv.Atoi(sym)
			if err != nil {
				panic(err)
			}

			stack.Push(num)
		} else {
			r := stack.Top()
			stack.Pop()
			l := stack.Top()
			stack.Pop()

			stack.Push(evaluate(l, r, sym))
		}
	}

	return stack.Top()
}

func evaluate(l, r int, op string) (res int) {
	switch op {
	case "+":
		res = l + r
	case "-":
		res = l - r
	case "*":
		res = l * r
	case "/":
		res = l / r
	}

	return
}

// Stack class
type Stack struct {
	buffer []string
}

func (this *Stack) isEmpty() bool {
	return len(this.buffer) == 0
}

func (this *Stack) Push(value string) {
	this.buffer = append(this.buffer, value)
}

func (this *Stack) Top() string {
	return this.buffer[len(this.buffer)-1]
}

func (this *Stack) Pop() bool {
	size := len(this.buffer)

	if size > 0 {
		this.buffer = this.buffer[:size-1]
		return true
	}

	return false
}

func createStack(size int) Stack {
	return Stack{
		buffer: make([]string, 0, size),
	}
}

// IntStack class
type IntStack struct {
	buffer []int
}

func (this *IntStack) isEmpty() bool {
	return len(this.buffer) == 0
}

func (this *IntStack) Push(value int) {
	this.buffer = append(this.buffer, value)
}

func (this *IntStack) Top() int {
	return this.buffer[len(this.buffer)-1]
}

func (this *IntStack) Pop() bool {
	size := len(this.buffer)

	if size > 0 {
		this.buffer = this.buffer[:size-1]
		return true
	}

	return false
}

func createIntStack(size int) IntStack {
	return IntStack{
		buffer: make([]int, 0, size),
	}
}
