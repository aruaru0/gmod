# math
## GCD
GCD calculates the Greatest Common Divisor of two integers using the Euclidean algorithm.

`func GCD(a, b int) int`
returns the greatest common multiple of a and b

```go
package main

import (
    "fmt"

    "github.com/aruaru0/gmod/math"
)

func main() {
    fmt.Println("gcd(6, 2) = ", math.GCD(6, 2))
}
```

## ExtGCD
ExtGCD calculates the Extended Greatest Common Divisor of two integers.

two numbers 𝑎,𝑏 For 𝑎=𝑎′𝑔, 𝑏=𝑏′𝑔 (𝑎′,𝑏′ are relatively prime) such that 𝑔=𝑔𝑐𝑑(𝑎,𝑏) Calculate.

In the extended GCD, 𝑎𝑥+𝑏𝑦=𝑔 
𝑥,𝑦,𝑔 such that Calculate.

`func ExtGCD(a, b int) (g int, x int, y int)`


```go
package main

import (
    "fmt"

    "github.com/aruaru0/gmod/math"
)

func main() {
    g, x, y := math.ExtGCD(18, 20)
    fmt.Println("ext gcd(18, 20) = ", g, x, y)}
```

## LCM
LCM calculates the Least Common Multiple of two integers.

`func LCM(a, b int) int`
returns the least common multiple of a and b


```go
package main

import (
	"fmt"

	"github.com/aruaru0/gmod/math"
)

func main() {
	fmt.Println("lcm(6, 15) = ", math.LCM(6, 15))
}
```

## Pfs, PfaMap
prime factorization


`func Pfs(n int) []int`
Returns the result of factorizing n as a list
`func PfsMap(n int) map[int]int`
Returns the result of factorizing n as a map


```go
import (
	"fmt"

	"github.com/aruaru0/gmod/math"
)

func main() {
	fmt.Println("pfs(12345678) = ", math.Pfs(123456780))
	fmt.Println("pfsMap(12345678) = ", math.PfsMap(123456780))
}
```

## Divisor
enumerate divisor

`func Divisor(n int) []int`
Returns a list of divisors of n (unsorted)

```go
import (
	"fmt"

	"github.com/aruaru0/gmod/math"
)

func main() {
	fmt.Println("Divisor(3*5*7) = ", math.Divisor(3*5*7))
}
```