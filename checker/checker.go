package main

import (
	"fmt"
	"os"
	"strconv"
)

// Stack type to represent the stack
type Stack []int

// parseInput parses the stack from the command-line arguments
func parseInput(args []string) (Stack, error) {
	var stack Stack
	seen := make(map[int]bool)

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}

		// Check for duplicates
		if seen[num] {
			return nil, fmt.Errorf("duplicate number")
		}
		seen[num] = true
		stack = append(stack, num)
	}

	return stack, nil
}

// executeInstruction executes the given instruction on the stacks
func executeInstruction(inst string, stackA, stackB Stack) (Stack, Stack, error) {
	switch inst {
	case "pa":
		// Push from B to A
		if len(stackB) == 0 {
			return stackA, stackB, fmt.Errorf("Error")
		}
		return append([]int{stackB[0]}, stackA...), stackB[1:], nil
	case "pb":
		// Push from A to B
		if len(stackA) == 0 {
			return stackA, stackB, fmt.Errorf("Error")
		}
		return stackA[1:], append([]int{stackA[0]}, stackB...), nil
	case "sa":
		// Swap the first two elements of A
		if len(stackA) < 2 {
			return stackA, stackB, nil
		}
		stackA[0], stackA[1] = stackA[1], stackA[0]
		return stackA, stackB, nil
	case "sb":
		// Swap the first two elements of B
		if len(stackB) < 2 {
			return stackA, stackB, nil
		}
		stackB[0], stackB[1] = stackB[1], stackB[0]
		return stackA, stackB, nil
	case "ss":
		// Swap both A and B
		stackA, stackB, _ = executeInstruction("sa", stackA, stackB)
		stackA, stackB, _ = executeInstruction("sb", stackA, stackB)
		return stackA, stackB, nil
	case "ra":
		// Rotate A
		if len(stackA) <= 1 {
			return stackA, stackB, nil
		}
		return append(stackA[1:], stackA[0]), stackB, nil
	case "rb":
		// Rotate B
		if len(stackB) <= 1 {
			return stackA, stackB, nil
		}
		return stackA, append(stackB[1:], stackB[0]), nil
	case "rr":
		// Rotate both A and B
		stackA, stackB, _ = executeInstruction("ra", stackA, stackB)
		stackA, stackB, _ = executeInstruction("rb", stackA, stackB)
		return stackA, stackB, nil
	case "rra":
		// Reverse rotate A
		if len(stackA) <= 1 {
			return stackA, stackB, nil
		}
		last := stackA[len(stackA)-1]
		return append([]int{last}, stackA[:len(stackA)-1]...), stackB, nil
	case "rrb":
		// Reverse rotate B
		if len(stackB) <= 1 {
			return stackA, stackB, nil
		}
		last := stackB[len(stackB)-1]
		return stackA, append([]int{last}, stackB[:len(stackB)-1]...), nil
	case "rrr":
		// Reverse rotate both A and B
		stackA, stackB, _ = executeInstruction("rra", stackA, stackB)
		stackA, stackB, _ = executeInstruction("rrb", stackA, stackB)
		return stackA, stackB, nil
	default:
		return stackA, stackB, fmt.Errorf("invalid instruction")
	}
}

// isSorted checks if stack A is sorted in ascending order
func isSorted(stack Stack) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[i-1] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		// No arguments, nothing to do
		return
	}

	// Parse stack from the command-line arguments
	stackArgs := os.Args[1:]
	stackA, err := parseInput(stackArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Initialize stackB (empty)
	stackB := Stack{}

	// Process instructions from the command-line arguments
	for _, inst := range os.Args[1:] {
		// Stop processing when we reach the end of the instructions
		if inst == stackArgs[0] {
			break
		}

		// Execute the instruction on the stacks
		stackA, stackB, err = executeInstruction(inst, stackA, stackB)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	// Check if the stack is sorted and B is empty
	if isSorted(stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
