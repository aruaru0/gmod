package math

import (
	"testing"
)

func TestDivisor(t *testing.T) {
	d := Divisor(3 * 5 * 7)

	r := []int{1, 3 * 5 * 7, 3, 5 * 7, 5, 3 * 7, 7, 3 * 5}
	for i, e := range d {
		if e != r[i] {
			t.Errorf("divisor(3*5*7) = %d, want %d", e, r[i])
		}
	}

	d = Divisor(1)
	if len(d) != 1 || d[0] != 1 {
		t.Errorf("divisor(1) = %d, want %d", d[0], 1)
	}
}
