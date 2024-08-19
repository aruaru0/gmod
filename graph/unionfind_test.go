package graph

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewDsu(10)
	uf.Merge(0, 1)

	if uf.Same(0, 1) == false {
		t.Errorf("uf.Same(0, 1) = false")
	}

	uf.Merge(1, 2)
	if uf.Same(0, 2) == false {
		t.Errorf("uf.Same(0, 2) = false")
	}

	uf.Merge(5, 6)

	if uf.Same(0, 1) == false {
		t.Errorf("uf.Same(0, 1) = false")
	}
	if uf.Same(0, 2) == false {
		t.Errorf("uf.Same(0, 2) = false")
	}
	if uf.Same(1, 2) == false {
		t.Errorf("uf.Same(1, 2) = false")
	}
	if uf.Same(5, 6) == false {
		t.Errorf("uf.Same(1, 2) = false")
	}

	if uf.Same(0, 5) == true {
		t.Errorf("uf.Same(0, 5) = true")
	}

	if len(uf.Groups()) != 7 {
		t.Errorf("len(uf.Groups()) = %d", len(uf.Groups()))
	}
	t.Log("Groups:", uf.Groups())
}
