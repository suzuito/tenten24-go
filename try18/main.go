package main

import (
	"fmt"
	"iter"
	"sync"
)

type Node[T any] struct {
	Value    T
	Children []*Node[T]
}

func dfs[T any](
	root *Node[T],
	yield func(T) bool,
	isBreaked *bool,
	lockIsBreaked *sync.Mutex,
) {
	if *isBreaked {
		return
	}
	if !yield(root.Value) {
		lockIsBreaked.Lock()
		*isBreaked = true
		lockIsBreaked.Unlock()
		return
	}
	for _, child := range root.Children {
		dfs(child, yield, isBreaked, lockIsBreaked)
	}
}

func All[T any](root *Node[T]) iter.Seq[T] {
	isBreaked := false
	lockIsBreaked := sync.Mutex{}
	return func(yield func(T) bool) {
		dfs(
			root,
			yield,
			&isBreaked,
			&lockIsBreaked,
		)
	}
}

func main() {
	root := Node[string]{
		Value: "A",
		Children: []*Node[string]{
			{
				Value: "B",
				Children: []*Node[string]{
					{
						Value:    "C",
						Children: []*Node[string]{},
					},
					{
						Value:    "D",
						Children: []*Node[string]{},
					},
				},
			},
			{
				Value:    "E",
				Children: []*Node[string]{},
			},
		},
	}
	for a := range All(&root) {
		if a == "C" {
			break
		}
		fmt.Println(a)
	}
}
