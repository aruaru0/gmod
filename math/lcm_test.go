package math

import "testing"

func TestLCM(t *testing.T) {

	ab := [][3]int{{6, 2, 6}, {2, 3, 6}, {88, 20, 440}, {9, 0, 0}}

	for _, e := range ab {
		ret := LCM(e[0], e[1])
		if ret != e[2] {
			t.Errorf("LCM(%d, %d) = %d, want %d", e[0], e[1], ret, e[2])
		}
	}
}
