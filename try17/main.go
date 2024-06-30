package main

import (
	"fmt"
	"iter"
	"slices"
)

func Map[FROM any, TO any](
	seq iter.Seq2[int, FROM],
	f func(int, FROM) TO,
) iter.Seq2[int, TO] {
	return func(yield func(int, TO) bool) {
		for i, v := range seq {
			vv := f(i, v)
			if !yield(i, vv) {
				break
			}
		}
	}
}

func main() {
	seq := Map[int, string](
		slices.All([]int{10, 20}),
		func(i int, from int) string {
			return fmt.Sprintf("0x%x", from)
		},
	)
	for i, v := range seq {
		fmt.Println(i, v)
	}
}
