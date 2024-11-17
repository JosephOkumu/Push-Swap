package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // If no arguments provided, just return
    if len(os.Args) == 1 {
        return
    }

    // Split the input string into numbers
    input := strings.Fields(os.Args[1])
    
    stackA := NewStack("a")
    stackB := NewStack("b")

    // Validate and convert arguments
    seen := make(map[int]bool)
    // Load numbers in reverse order
    for i := len(input) - 1; i >= 0; i-- {
        num, err := strconv.Atoi(input[i])
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error")
            return
        }
        if seen[num] {
            fmt.Fprintln(os.Stderr, "Error")
            return
        }
        seen[num] = true
        stackA.Push(num)
    }

    // Get and print operations
    operations := Sort(stackA, stackB)
    for _, op := range operations {
        fmt.Println(op)
    }
}