package pushswap

import (
	"pushswap/lib"
)

func Sort(nums []int) string {
	result := "" 
	stack := []*lib.Stack{lib.NewStack(nums), lib.NewStack(make ([]int, 0, len(nums)))}
	maxNum := len(stack[0].Nums)-1
	maxBits := 0
	for maxNum >> maxBits != 0{
		maxBits++
	}
	for i := 0; i < maxBits; i++ {
		for j := 0; j < len(nums); j++ {
			num := stack[0].Nums[len(stack[0].Nums)-1]
			if (num>>i)&1 == 1{
				Action(stack, "ra")
				result += "ra\n"
			}else {
				Action(stack, "pb")
				result += "pb\n"
			}
		}
		for len(stack[1].Nums) != 0 {
			Action(stack, "pa")
			result += "pa\n"
		}
	}
	return result
}