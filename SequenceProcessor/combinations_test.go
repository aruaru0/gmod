package SequenceProcessor

import (
	"testing"
)

func TestCombinations(t *testing.T) {
	a := []int{1, 2, 3, 4}

	exp := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}

	cnt := 0
	for e := range Combinations(a, 2, 1) {
		for i := 0; i < len(e); i++ {
			if e[i] != exp[cnt][i] {
				t.Error("mismatch ", e, "!=", exp[cnt])
			}
		}
		cnt++
	}

	if cnt != len(exp) {
		t.Error("mismatch cnt", cnt, "!=", len(exp))
	}
}
