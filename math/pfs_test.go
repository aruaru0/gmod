package math

import (
	"testing"
)

func TestPfs(t *testing.T) {
	p := Pfs(12345678)

	res := []int{2, 3, 3, 47, 14593}

	if len(p) != len(res) {
		t.Errorf("Pfs(12345678) = %d, want %d", len(p), len(res))
	}

	for i, e := range p {
		if res[i] != e {
			t.Error("Pfs(12345678) = ", p, "want", res)
		}
	}
}

func TestPfsMap(t *testing.T) {
	p := PfsMap(12345678)

	res := map[int]int{2: 1, 3: 2, 47: 1, 14593: 1}

	if len(p) != len(res) {
		t.Errorf("PfsMap(12345678) = %d, want %d", len(p), len(res))
	}

	for e := range p {
		if res[e] != p[e] {
			t.Error("PfsMap(12345678) = ", p, "want", res)
		}
	}
}
