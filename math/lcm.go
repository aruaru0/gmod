package math

// LCM calculates the Least Common Multiple of two integers.
//
// Parameters a and b are the two integers for which the LCM is calculated.
// Returns the LCM of a and b.
func LCM(a, b int) int {
	return a / GCD(a, b) * b
}
