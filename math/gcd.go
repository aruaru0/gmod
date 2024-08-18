package math

// GCD calculates the Greatest Common Divisor of two integers using the Euclidean algorithm.
//
// Parameters a and b are the two integers for which the GCD is calculated.
// Returns the GCD of a and b.
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// ExtGCD calculates the Extended Greatest Common Divisor of two integers.
//
// Parameters a and b are the two integers for which the Extended GCD is calculated.
// Returns the GCD and the coefficients of Bézout's identity, x and y, such that ax + by = gcd(a, b).
func ExtGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	d, y, x := ExtGCD(b, a%b)
	y -= a / b * x
	return d, x, y
}
