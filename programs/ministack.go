package pushswap

import (
	"pushswap/lib"
	"slices"
)

func SortSmallNumbs(nums []int) string {
	instructions := ""
	stacks := []*lib.Stack{lib.NewStack(nums), lib.NewStack(make([]int, 0, len(nums)))}
	instructions = SortInstructions(nums, stacks)
	return instructions
}

func SortInstructions(nums []int, stacks []*lib.Stack) string {
	halfSize := len(nums) / 2
	instructions := ""
	for {
		stackA, stackB := stacks[0].Nums, stacks[1].Nums
		if isASorted(stackA) {
			break
		}
		op := Complete(OperationA(stackA, halfSize), stackA, stackB)
		Action(stacks, op)
		instructions += op + "\n"
	}
	for {
		stackA, stackB := stacks[0].Nums, stacks[1].Nums
		if len(stackA) == len(nums) && isASorted(stackA) {
			break
		}
		ins := OperationB(stackB)
		Action(stacks, ins)
		instructions += ins + "\n"
	}
	return instructions
}

// OperationA generates operations for stack A (sa, ra, rra, pb).
func OperationA(stackA []int, halfSize int) string {
	swapIdx := ASwapIndex(stackA)
	midPoint, lastIdx := len(stackA)/2, len(stackA)-1
	if stackA[lastIdx] < halfSize {
		return "pb"
	}
	if swapIdx > 0 {
		if isOrdered(stackA[lastIdx-1], stackA[lastIdx]) {
			return "sa"
		} else if swapIdx < midPoint {
			return "rra"
		}
	} else {
		minIdx := IndexOfSmall(stackA)
		if minIdx == lastIdx {
			return "pb"
		} else if minIdx < midPoint || isOrdered(stackA[lastIdx], stackA[0]) {
			return "rra"
		}
	}
	return "ra"
}

// OperationB generates operations for stack B (sb, rb, rrb, pa).
func OperationB(stackB []int) string {
	stackBCopy := make([]int, len(stackB))
	copy(stackBCopy, stackB)
	lib.UInts(stackBCopy)
	swapIdx := BSwapIndex(stackBCopy)
	midPoint, lastIdx := len(stackBCopy)/2, len(stackBCopy)-1
	if swapIdx > 0 {
		if isOrdered(stackBCopy[lastIdx], stackBCopy[lastIdx-1]) {
			return "sb"
		} else if swapIdx < midPoint {
			return "rrb"
		}
	} else {
		maxNumIdx := IndexOfBig(stackBCopy)
		if maxNumIdx == lastIdx {
			return "pa"
		} else if maxNumIdx < midPoint || isOrdered(stackBCopy[0], stackBCopy[lastIdx]) {
			return "rrb"
		}
	}
	return "rb"
}

// Complete adjusts operations to combine previous operations into a more efficient sequence.
func Complete(ins string, stackA, stackB []int) string {
	if len(stackB) < 2 || isBSorted(stackB) {
		return ins
	}
	bins := OperationB(stackB)
	switch ins { // ins is an instruction from OperationA(stackA)
	case "pb":
		if bins == "sb" {
			if t := len(stackA) - ASwapIndex(stackA); t < 3 {
				if t == 1 {
					return "ss"
				}
				return "ra"
			} else {
				return "sb"
			}
		}
	case "sa":
		if t := len(stackB) - BSwapIndex(stackB); t < 3 {
			if bins == "sb" {
				return "ss"
			}
			return "rb"
		}
	case "ra":
		if bins == "rb" || bins == "sb" && len(stackB) == 2 {
			return "rr"
		}
	case "rra":
		if bins == "rrb" || bins == "sb" && len(stackB) == 2 {
			return "rrr"
		}
	}
	return ins
}

// IndexOfSmall finds the index of the indexOfSmall element in a slice.
func IndexOfSmall(nums []int) int {
	minVal, minIdx := 0x3f3f3f3f, -1
	for i := 0; i < len(nums); i++ {
		if minVal > nums[i] {
			minVal = nums[i]
			minIdx = i
		}
	}
	return minIdx
}

// IndexOfBig finds the index of the largest element in a slice.
func IndexOfBig(nums []int) int {
	maxVal, maxIdx := -1, -1
	for i := 0; i < len(nums); i++ {
		if maxVal < nums[i] {
			maxVal = nums[i]
			maxIdx = i
		}
	}
	return maxIdx
}

func ASwapIndex(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i]-nums[i-1] == 1 {
			return i
		}
	}
	return -1
}

func BSwapIndex(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1]-nums[i] == 1 {
			return i
		}
	}
	return -1
}

// isOrdered checks if b is exactly one more than a.
func isOrdered(a, b int) bool {
	return b-a == 1
}

// isASorted checks if stack A is sorted in ascending order.
func isASorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] != 1 {
			return false
		}
	}
	return true
}

// isBSorted checks if stack B is sorted in ascending order
func isBSorted(nums []int) bool {
	return slices.IsSorted(nums)
}
