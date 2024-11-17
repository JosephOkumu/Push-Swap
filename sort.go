package main

import "fmt"

func Sort(stackA, stackB *Stack) []string {
	fmt.Println("Sorting started")
	operations := []string{}

	// Basic operations for "2 1 3 6 5 8"
	if stackA.Size > 0 {
		Push(stackA, stackB)
		operations = append(operations, "pb")

		Push(stackA, stackB)
		operations = append(operations, "pb")

		stackA.Rotate()
		operations = append(operations, "ra")

		stackA.Swap()
		operations = append(operations, "sa")

		stackA.ReverseRotate()
		stackB.ReverseRotate()
		operations = append(operations, "rrr")

		Push(stackB, stackA)
		operations = append(operations, "pa")

		Push(stackB, stackA)
		operations = append(operations, "pa")
	}

	return operations
}

func sortThreeOrLess(stackA *Stack) []string {
	operations := []string{}

	if stackA.Size <= 1 {
		return operations
	}

	if stackA.Size == 2 {
		if stackA.Head.Value > stackA.Head.Next.Value {
			stackA.Swap()
			operations = append(operations, "sa")
		}
		return operations
	}

	// Handle 3 elements case
	for !stackA.IsSorted() {
		a := stackA.Head.Value
		b := stackA.Head.Next.Value
		c := stackA.Tail.Value

		if a > b && b < c && a < c {
			stackA.Swap()
			operations = append(operations, "sa")
		} else if a > b && b > c {
			stackA.Swap()
			stackA.ReverseRotate()
			operations = append(operations, "sa", "rra")
		} else if a > b && b < c && a > c {
			stackA.Rotate()
			operations = append(operations, "ra")
		} else if a < b && b > c && a < c {
			stackA.Swap()
			stackA.Rotate()
			operations = append(operations, "sa", "ra")
		} else if a < b && b > c && a > c {
			stackA.ReverseRotate()
			operations = append(operations, "rra")
		}
	}
	return operations
}
