package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

// parseInput converts command-line arguments to integers.
func parseInput(args []string) (Stack, error) {
	var stack Stack
	seen := make(map[int]bool)

	for _, arg := range strings.Fields(strings.Join(args, " ")) {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}

		if seen[num] {
			return nil, fmt.Errorf("duplicate number")
		}
		seen[num] = true
		stack = append(stack, num)
	}

	return stack, nil
}

// readInstructions reads instructions from standard input.
func readInstructions() ([]string, error) {
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inst := strings.TrimSpace(scanner.Text())
		if inst == "" {
			break
		}
		instructions = append(instructions, inst)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}

// push moves the top element from source to destination stack.
func push(dest, src Stack) (Stack, Stack) {
	if len(src) == 0 {
		return dest, src
	}
	return append(dest, src[0]), src[1:]
}

// swap swaps the first two elements of a stack.
func swap(stack Stack) Stack {
	if len(stack) < 2 {
		return stack
	}
	stack[0], stack[1] = stack[1], stack[0]
	return stack
}

// rotate shifts all elements up by 1.
func rotate(stack Stack) Stack {
	if len(stack) <= 1 {
		return stack
	}
	return append(stack[1:], stack[0])
}

// reverseRotate shifts all elements down by 1.
func reverseRotate(stack Stack) Stack {
	if len(stack) <= 1 {
		return stack
	}
	last := stack[len(stack)-1]
	return append([]int{last}, stack[:len(stack)-1]...)
}

// isSorted checks if the stack is sorted in ascending order.
func isSorted(stack Stack) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[i-1] {
			return false
		}
	}
	return true
}

func main() {
	// Handle no arguments case.
	if len(os.Args) < 2 {
		return
	}

	// Parse and validate input arguments.
	stackA, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	stackB := Stack{}

	// Read and validate instructions.
	instructions, err := readInstructions()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Execute instructions.
	for _, inst := range instructions {
		switch inst {
		case "pa":
			stackA, stackB = push(stackA, stackB)
		case "pb":
			stackB, stackA = push(stackB, stackA)
		case "sa":
			stackA = swap(stackA)
		case "sb":
			stackB = swap(stackB)
		case "ss":
			stackA = swap(stackA)
			stackB = swap(stackB)
		case "ra":
			stackA = rotate(stackA)
		case "rb":
			stackB = rotate(stackB)
		case "rr":
			stackA = rotate(stackA)
			stackB = rotate(stackB)
		case "rra":
			stackA = reverseRotate(stackA)
		case "rrb":
			stackB = reverseRotate(stackB)
		case "rrr":
			stackA = reverseRotate(stackA)
			stackB = reverseRotate(stackB)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	// Check sorting condition.
	if isSorted(stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
