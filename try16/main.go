package main

import "fmt"

func Alphabet(yield func(int) bool) {
	for i := 65; i < 133; i++ {
		if !yield(i) {
			break
		}
	}
}

func Alphabet2(yield func(int, string) bool) {
	for i := 65; i < 133; i++ {
		if !yield(i, fmt.Sprintf("%x", i)) {
			break
		}
	}
}

func main() {
	for c := range Alphabet {
		fmt.Printf("%c", c)
		if c == 'C' {
			break
		}
	}
	fmt.Println("")

	for i, c := range Alphabet2 {
		fmt.Printf("%c:%d:%s\n", i, i, c)
	}
}
