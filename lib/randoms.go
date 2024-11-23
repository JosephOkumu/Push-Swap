package lib

import (
	"fmt"
	"math/rand"
)

// GenerateRandNums prompts the user for input and generates random unique integers.
func GenerateRandNums() ([]int, error) {
	fmt.Println("Please enter the minimum value, maximum value, and the quantity of random numbers.")
	var min, max, count int
	_, err := fmt.Scan(&min, &max, &count)
	if err != nil {
		return nil, fmt.Errorf("invalid input: %v", err)
	}
	if min > max {
		return nil, fmt.Errorf("the minimum value cannot be larger than the maximum value %v", max)
	}
	if max-min < count {
		return nil, fmt.Errorf("the range is too small to generate the required amount of unique numbers")
	}
	return CreateUniqueNums(min, max, count)
}

// CreateUniqueNums generates a list of unique random numbers in a specified range.
func CreateUniqueNums(min, max, count int) ([]int, error) {
	randomNumbers := make([]int, count)
	randomPermutation := rand.Perm(max - min + 1)
	for i := 0; i < count; i++ {
		randomNumbers[i] = randomPermutation[i] + min
	}
	return randomNumbers, nil
}
