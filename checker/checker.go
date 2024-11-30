package main

import (
	"fmt"
	"io/ioutil"
	"os"
	pushswap "pushswap/programs"
	"pushswap/lib"
	"sort"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	// if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeNamedPipe == 0 {
	// 	fmt.Println("Error")
	// 	return
	// }
	instructions, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error")
		return
	}
	nums, err := lib.UniqueNums(args[0])
	if err != nil {
		fmt.Println("Error")
		return
	}
	stack := []*lib.Stack{lib.NewStack(nums), lib.NewStack(make([]int, 0, len(nums)))}
	operations := strings.Split(string(instructions), "\n")
	l := len(operations) - 1
	if l > 0 && operations[l-1] == "" {
		l--
	}
	operations = operations[:l]
	for _, v := range operations {
		if f, ok := pushswap.Arrange[v]; ok {
			f(stack)
		} else {
			fmt.Println("Error")
			return
		}
	}
	res := stack[0].Nums
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	if sort.IntsAreSorted(res) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
	// fmt.Println("Stack a", stack[0].Nums)
	// fmt.Println("Stack b", stack[1].Nums)
}
