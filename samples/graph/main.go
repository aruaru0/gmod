package main

import (
	"fmt"

	gr "github.com/aruaru0/gmod/graph"
)

func dsuTest() {
	uf := gr.NewDsu(10)
	uf.Merge(0, 1)
	uf.Merge(1, 2)
	uf.Merge(5, 6)
	fmt.Println(uf.Same(0, 2))
	fmt.Println(uf.Same(0, 5))
	fmt.Println(uf.Groups())
}

func main() {
	dsuTest()
}
