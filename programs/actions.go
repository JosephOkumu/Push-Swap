package pushswap

import "pushswap/lib"

type f func(stack []*lib.Stack)

var Arrange = map[string]f {
	"pa": pa,
	"pb": pb,
	"sa": sa,
	"sb": sb,
	"ss": ss,
	"ra": ra,
	"rb": rb,
	"rr": rr,
	"rra": rra,
	"rrb": rrb,
	"rrr": rrr,
}

func Action(stack []*lib.Stack, instructs string) {
	Arrange[instructs](stack)
}