package main

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

func AddNumbers[T Number](a, b T) T {
	return a + b
}
