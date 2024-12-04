package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(0)
	}
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	userInput := strings.Split(os.Args[1], " ")
	var aStack []int
	var bStack []int

	err := AppendNumbers(userInput, &aStack)
	if err {
		return
	}

	var allInstructions []string = []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"}
	var instructionsToUse []string

	errorDetected := GetInstructions(allInstructions, &instructionsToUse)
	if errorDetected {
		return
	}

	for i := 0; i < len(instructionsToUse); i++ {
		if instructionsToUse[i] == "pa" {
			PushTop(&bStack, &aStack)
		} else if instructionsToUse[i] == "pb" {
			PushTop(&aStack, &bStack)
		} else if instructionsToUse[i] == "sa" {
			Swap(&aStack)
		} else if instructionsToUse[i] == "sb" {
			Swap(&bStack)
		} else if instructionsToUse[i] == "ss" {
			SwapBoth(&aStack, &bStack)
		} else if instructionsToUse[i] == "ra" {
			Rotate(&aStack)
		} else if instructionsToUse[i] == "rb" {
			Rotate(&bStack)
		} else if instructionsToUse[i] == "rr" {
			RotateBoth(&aStack, &bStack)
		} else if instructionsToUse[i] == "rra" {
			ReverseRotate(&aStack)
		} else if instructionsToUse[i] == "rrb" {
			ReverseRotate(&bStack)
		} else {
			ReverseRotateBoth(&aStack, &bStack)
		}
	}

	if sort.IntsAreSorted(aStack) && len(bStack) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}


