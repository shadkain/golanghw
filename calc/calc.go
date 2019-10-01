package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	in, err := readInput()
	if err != nil {
		crash(err)
	}

	out, err := calculate(in)
	if err != nil {
		crash(err)
	}

	fmt.Println(out)
}

func crash(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}

// Function declarations
func calculate(in string) (out int, err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf("%s", perr)
		}
	}()

	regExp := createRegExp()
	symbols := extractSymbols(in, regExp)
	checkBraces(symbols)

	exp := prepareExpression(symbols)
	out = evaluateExpression(exp)

	return
}

func readInput() (string, error) {
	args := os.Args

	nArgs := len(args)
	if nArgs != 2 {
		return "", fmt.Errorf("expected 1 argument, but provided %d", nArgs-1)
	}

	return args[1], nil
}

func createRegExp() *regexp.Regexp {
	return regexp.MustCompile(`\d+|[()+/*-]`)
}

func extractSymbols(source string, regExp *regexp.Regexp) []string {
	return regExp.FindAllString(source, -1)
}

func checkBraces(symbols []string) {
	bracketLevel := 0
	for _, sym := range symbols {
		if isOpenBrace(sym) {
			bracketLevel++
		} else if isCloseBrace(sym) {
			bracketLevel--
		}
	}

	if bracketLevel != 0 {
		panic("invalid braces")
	}
}

func prepareExpression(symbols []string) []string {
	output := make([]string, 0, len(symbols))
	stack := NewStringStack(len(symbols))

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
			} else if isOpenBrace(sym) || stack.IsEmpty() {
				stack.Push(sym)
			} else {
				top := stack.Top()
				if !isOpenBrace(top) && priority(top) >= priority(sym) {
					output = append(output, top)
					stack.Pop()
				}

				stack.Push(sym)
			}
		}
	}
	for !stack.IsEmpty() {
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
	default:
		priority = 0
	}

	return
}

func evaluateExpression(exp []string) int {
	stack := NewIntStack(len(exp))
	isNumber := createIsNumberFunc()

	for _, sym := range exp {
		if isNumber(sym) {
			num, _ := strconv.Atoi(sym)

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
