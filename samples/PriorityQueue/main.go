package main

import (
	"fmt"

	pq "github.com/aruaru0/gmod/PriorityQueue"
)

func main() {
	pq := pq.New[int](func(i, j int) bool {
		return i < j
	})

	a := []int{4, 6, 7, 2, 43, 5, 3, 8, 33, 2, 1, 3, 12}

	for _, e := range a {
		pq.Push(e)
	}

	fmt.Println("len = ", pq.Len())

	for pq.Len() != 0 {
		v, ok := pq.Pop()
		fmt.Println("val =", v, "ok =", ok)
	}

}
