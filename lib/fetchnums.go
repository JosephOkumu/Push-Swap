package lib

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//UniqInts returns slice of unique integers from the given string s.
//Only distinct numbers separated with single space are allowed
func UniqueNums(str string) ([]int, error) {
	input := strings.Split(str, " ")
	nums := make([]int, len(input))
	existing := make(map[int]bool, len(nums))
	for i, num := range input {
		num, err := strconv.Atoi(num)
		if err != nil || existing[num] {
			return nil, fmt.Errorf("invalid number")
		}
		existing[num] = true
		nums[i] = num
	}
	return nums, nil
}

//white spaces and non-digits are allowed
func GetNums(str string) []int {
	fields := strings.Fields(str)
	nums := make([]int, 0, len(fields))
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

//InputInts takes input from the User and returns numbers
func InputInts() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return GetNums(scanner.Text())
}

//UInts changes nums to unsigned integers by replacing them with their indexes when they're sorted
func UInts(nums []int) {
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)
	sort.Ints(copiedNums)
	for i, n := range nums {
		nums[i] = FindElementIndex(copiedNums, n)
	}
}
