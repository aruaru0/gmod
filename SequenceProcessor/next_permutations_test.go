package SequenceProcessor

import (
	"sort"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	a := []int{1, 2, 3}

	exp := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	cnt := 0
	for pos := 0; ; pos++ {
		// fmt.Println(a, exp[pos])
		for i := 0; i < len(a); i++ {
			if a[i] != exp[pos][i] {
				t.Error("mismatch a=1", a, "!=", exp[pos])
			}
		}
		cnt++
		if !NextPermutation(sort.IntSlice(a)) {
			break
		}
	}
	if cnt != len(exp) {
		t.Error("mismatch cnt", cnt, "!=", len(exp))
	}
}
