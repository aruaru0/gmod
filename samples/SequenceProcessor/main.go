package main

import (
	"fmt"
	"sort"

	sp "github.com/aruaru0/gmod/SequenceProcessor"
)

func main() {
	a := []int{1, 2, 3}

	fmt.Println("Permutation:")
	for {
		fmt.Println(a, "")
		if !sp.NextPermutation(sort.IntSlice(a)) {
			break
		}
	}

	fmt.Println("Combination:")
	for e := range sp.Combinations(a, 2, 1) {
		fmt.Println(e, "")
	}
}
