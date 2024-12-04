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

// Reads instructions on the standard input and saves them. If gets non-existing insturction - returns
func GetInstructions(allInstructions []string, instructionsToUse *[]string) bool {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for i, k := range allInstructions {
			if scanner.Text() == "stop" {
				return false
			}

			if scanner.Text() == "" {
				continue
			}

			if k == scanner.Text() {
				*instructionsToUse = append(*instructionsToUse, scanner.Text())
				break
			}

			if i == len(allInstructions)-1 {
				fmt.Println("Error")
				return true
			}
		}
	}
	return false
}

// Append all number to A stack
func AppendNumbers(userInput []string, aStack *[]int) bool {
	for i := 0; i < len(userInput); i++ {
		numToAppend, err := strconv.Atoi(userInput[i])
		if err != nil {
			fmt.Println("Error")
			return true
		}

		// Check for duplicates
		if i != 0 {
			for k := 0; k < len(*aStack); k++ {
				if (*aStack)[k] == numToAppend {
					fmt.Println("Error")
					return true
				}
			}
		}

		(*aStack) = append((*aStack), numToAppend)
	}
	return false
}


