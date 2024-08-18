package SequenceProcessor

// combinations generates all possible combinations of the given list with the specified length.
//
// list is the input list of integers, k is the length of the combinations, and buf is the buffer size of the output channel.
// chan []int is the channel that outputs all combinations.
func Combinations(list []int, k, buf int) (c chan []int) {
	c = make(chan []int, buf)
	n := len(list)

	pattern := make([]int, k)

	var body func(pos, begin int)
	body = func(pos, begin int) {
		if pos == k {
			t := make([]int, k)
			copy(t, pattern)
			c <- t
			return
		}

		for num := begin; num < n+pos-k+1; num++ {
			pattern[pos] = list[num]
			body(pos+1, num+1)
		}
	}
	go func() {
		defer close(c)
		body(0, 0)
	}()

	return
}
