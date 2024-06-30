package main

import (
	"errors"
	"strconv"
	"strings"
)

func StackMachine(commands string) (int, error) {
	stack := []int{}
	commandList := strings.Split(commands, " ")

	for _, command := range commandList {
		if commandNumber, err := strconv.Atoi(command); err == nil {
			if commandNumber < 0 || commandNumber > 50000 {
				return 0, errors.New("number out of range, enter numbers between 0 and 50000")
			}
			stack = append(stack, commandNumber)
		} else {
			switch command {
			case "+":
				if len(stack) < 2 {
					return 0, errors.New("not enough numbers to add")
				}
				addedNumber1, addedNumber2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				result := addedNumber1 + addedNumber2
				if result > 50000 {
					return 0, errors.New("error: result too large")
				}
				stack = append(stack, result)
			case "-":
				if len(stack) < 2 {
					return 0, errors.New("not enough numbers to subtract")
				}
				subtractedNumber1, subtractednumber2 := stack[len(stack)-1], stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				result := subtractedNumber1 - subtractednumber2
				if result < 0 {
					return 0, errors.New("error: result is negative")
				}
				stack = append(stack, result)
			case "*":
				if len(stack) < 2 {
					return 0, errors.New("not enough numbers to multiply")
				}
				multipliedNumber1, multipliedNumber2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				result := multipliedNumber1 * multipliedNumber2
				if result > 50000 {
					return 0, errors.New("error: result is too large")
				}
				stack = append(stack, result)
			case "DUP":
				if len(stack) == 0 {
					continue
				}
				result := stack[len(stack)-1]
				stack = append(stack, result)
			case "POP":
				if len(stack) == 0 {
					return 0, errors.New("error: nothing to pop")
				}
				stack = stack[:len(stack)-1]
			case "SUM":
				if len(stack) == 0 {
					return 0, errors.New("error: nothing to add")
				}
				sum := 0
				for _, numberInCommand := range stack {
					sum += numberInCommand
				}
				if sum > 50000 {
					return 0, errors.New("error: result too large")
				}
				stack = []int{sum} // Reset stack to contain only the sum
			case "CLEAR":
				stack = nil
			default:
				return 0, errors.New("invalid input")
			}
		}
	}

	if len(stack) == 0 {
		return 0, errors.New("nothing in the stack")
	}
	return stack[len(stack)-1], nil
}

func main() {
}
