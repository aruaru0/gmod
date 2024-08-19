package math

// PfsMap returns a map of prime factors of a given number n.
//
// Parameter n is the number for which the prime factors are calculated.
// Returns a map where the keys are the prime factors and the values are their respective counts.
func PfsMap(n int) map[int]int {
	pfs := make(map[int]int)
	for n%2 == 0 {
		pfs[2] = pfs[2] + 1
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i] = pfs[i] + 1
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = pfs[n] + 1
	}

	return pfs
}

// Pfs calculates the prime factorization of a given number.
//
// Parameter n is the integer representing the number for which the prime factorization is calculated.
// Returns: a slice of integers representing the prime factorization of the given number.
func Pfs(n int) []int {
	pfs := make([]int, 0)

	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	if n > 2 {
		pfs = append(pfs, n)
	}

	return pfs
}
