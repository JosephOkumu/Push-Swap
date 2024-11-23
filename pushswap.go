package main

import (
	"fmt"
	"os"
	pushswap "pushswap/programs"
	"pushswap/lib"
	"sort"
)

func main() {
	var (
		nums []int
		err  error
	)
	if lib.EnableRandFlag() {
		nums, err = lib.GenerateRandNums()
		if err != nil {
			fmt.Println("Error")
			return
		}
	} else {
		args := os.Args[1:]
		if len(args) < 1 {
			return
		}
		nums, err = lib.UniqInts(args[0])
		if err != nil {
			fmt.Println("Error")
			return
		}
	}
	if sort.IntsAreSorted(nums) {
		return
	}
	file, _ := os.Create("push-swap-result.txt")
	defer file.Close()
	numsToStr := fmt.Sprint(nums)
	file.WriteString("\"" + numsToStr[1:len(numsToStr)-1] + "\"\n")
	instructions := ""
	lib.UInts(nums)
	if len(nums) <= 100 {
		instructions = pushswap.SortSmallNumbs(nums)
	} else {
		instructions = pushswap.Sort(nums)
	}
	fmt.Print(instructions)
	file.WriteString(instructions)
	//fmt.Println("step ", len(strings.Split(instructions, "\n"))-1)
}
