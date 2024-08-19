package math

// divisor returns a list of divisors of a given number n.
//
// Parameter n is the number for which the divisors are calculated.
// Returns a slice of integers representing the divisors of n.
func Divisor(n int) []int {
	ret := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if i*i != n {
				ret = append(ret, n/i)
			}
		}
	}
	return ret
}
