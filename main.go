package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	args := os.Args[1:]

	// validate and convert the arguments to integers
	stackA := make([]int, 0, len(args))
	seen := make(map[int]bool)

	for _, arg := range args {
		arg = strings.TrimSpace(arg)

		// Converting to integer
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}

		// Check for duplicates
		if seen[num] {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		seen[num] = true

		// Add the number to the stack
		stackA = append(stackA, num)
	}
	fmt.Println("Stack A:", stackA)
}


