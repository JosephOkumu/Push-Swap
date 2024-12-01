package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

// parseInput converts command-line arguments to integers
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

// readInstructions reads instructions from standard input
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

// push moves top element from source to destination stack
func push(dest, src Stack) (Stack, Stack) {
	if len(src) == 0 {
		return dest, src
	}
	return append(dest, src[0]), src[1:]
}

// swap swaps first two elements of a stack
func swap(stack Stack) Stack {
	if len(stack) < 2 {
		return stack
	}
	stack[0], stack[1] = stack[1], stack[0]
	return stack
}





