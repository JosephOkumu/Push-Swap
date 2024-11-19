package pushswap

import (
	"pushswap/lib"
)

// push the first top element of stack b to stack a
func pa(l []*lib.Stack) {
	if len(l[1].Nums) > 0 {
		l[0].PushBack(l[1].PopBack())
	}
}

// push the top first element of stack a to stack b
func pb(l []*lib.Stack) {
	if len(l[0].Nums) > 0 {
		l[1].PushBack(l[0].PopBack())
	}
}

// swap first 2 elements of stack a
func sa(l []*lib.Stack) {
	if len(l[0].Nums) > 1 {
		v1 := l[0].PopBack()
		v2 := l[0].PopBack()
		l[0].PushBack(v1)
		l[0].PushBack(v2)
	}
}

// swap first 2 elements of stack b
func sb(l []*lib.Stack) {
	if len(l[1].Nums) > 1 {
		v1 := l[1].PopBack()
		v2 := l[1].PopBack()
		l[1].PushBack(v1)
		l[1].PushBack(v2)
	}
}

// execute sa and sb
func ss(l []*lib.Stack) {
	sa(l)
	sb(l)
}

// rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func ra(l []*lib.Stack) {
	if len(l[0].Nums) == 0 {
		return
	}
	l[0].PushFront(l[0].PopBack())
}

// rotate stack b
func rb(l []*lib.Stack) {
	if len(l[1].Nums) == 0 {
		return
	}
	l[1].PushFront(l[1].PopBack())
}

// execute ra and rb
func rr(l []*lib.Stack) {
	ra(l)
	rb(l)
}

// reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func rra(l []*lib.Stack) {
	if len(l[0].Nums) == 0 {
		return
	}
	l[0].PushBack(l[0].PopFront())
}

// reverse rotate b
func rrb(l []*lib.Stack) {
	if len(l[1].Nums) == 0 {
		return
	}
	l[1].PushBack(l[1].PopFront())
}

func rrr(l []*lib.Stack) {
	rra(l)
	rrb(l)
}
