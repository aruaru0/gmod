package math

import (
	"testing"
)

func TestGCD(t *testing.T) {

	ab := [][3]int{{6, 2, 2}, {2, 3, 1}, {100, 20, 20}, {9, 0, 9}}

	for _, e := range ab {
		ret := GCD(e[0], e[1])
		if ret != e[2] {
			t.Errorf("GCD(%d, %d) = %d, want %d", e[0], e[1], ret, e[2])
		}
	}

}

func TestExtGCD(t *testing.T) {

	ab := [][5]int{{6, 2, 2, 0, 1}, {2, 3, 1, -1, 1}, {100, 20, 20, 0, 1}, {9, 0, 9, 1, 0}}

	for _, e := range ab {
		d, x, y := ExtGCD(e[0], e[1])
		if d != e[2] && x != e[3] && y != e[4] {
			t.Errorf("ExtGCD(%d, %d) = %d, want %d", e[0], e[1], d, e[2])
		}
	}
}
