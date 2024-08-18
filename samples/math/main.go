package main

import (
	"fmt"

	"github.com/aruaru0/gmod/math"
)

func gcdtest() {
	fmt.Println("gcd(6, 2) = ", math.GCD(6, 2))
	fmt.Println("gcd(2, 3) = ", math.GCD(2, 3))
	fmt.Println("gcd(100, 20) = ", math.GCD(100, 20))
	fmt.Println("gcd(9, 0) = ", math.GCD(9, 0))

	g, x, y := math.ExtGCD(18, 20)
	fmt.Println("ext gcd(18, 20) = ", g, x, y)
}

func lcmtest() {
	fmt.Println("lcm(6, 15) = ", math.LCM(6, 15))
}

func main() {
	gcdtest()
	lcmtest()
}
