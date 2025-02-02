package main

import (
	"fmt"

	"github.com/evlic/monkey"
)

func sum[T int | float64](a, b T) T {
	return a + b
}

func foo[T int | float64](a, b T) T {
	return a - b
}

type S1[T int | float64] struct {
	i T
}

func (s *S1[T]) Get() T {
	return s.i
}

type S1__monkey__[T int | float64] struct {
	S1[T]
}

func (s *S1__monkey__[T]) Get() T {
	return s.i + 1
}

func main() {
	monkey.Patch(sum[int], foo[int], monkey.OptGeneric)
	fmt.Println(sum(1, 2)) // display -1

	monkey.Patch((*S1[int]).Get, (*S1__monkey__[int]).Get, monkey.OptGeneric)
	s := S1[int]{i: 1}
	fmt.Println(s.Get()) // display 2
}
