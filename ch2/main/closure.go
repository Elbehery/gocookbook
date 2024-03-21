package main

import "fmt"

func SeqGenerator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	seq := SeqGenerator()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
}
