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

