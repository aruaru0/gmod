# SequenceProcessor
##　NextPermutations
NextPermutation generates the next permutation of the sortable collection x in lexical order.

`func NextPermutation(x sort.Interface) bool`

Enumerate permutations of x.

```Go
a := []int{1, 2, 3}

fmt.Println("Permutation:")
for {
        fmt.Println(a, "")
        if !sp.NextPermutation(sort.IntSlice(a)) {
                break
        }
}
```


## Combinations
Combinations generates all possible combinations of the given list with the specified length.

`func Combinations(list []int, k, buf int) (c chan []int)`

Enumerate the combinations (nCk) of selecting k items from n different items (len(list)).

```Go
a := []int{1, 2, 3}

fmt.Println("Combination:")
for e := range sp.Combinations(a, 2, 1) {
        fmt.Println(e, "")
}
```


Below is the entire executable code
```go 
package main

import (
        "fmt"
        "sort"

        sp "github.com/aruaru0/gmod/SequenceProcessor"
)

func main() {
        a := []int{1, 2, 3}

        fmt.Println("Permutation:")
        for {
                fmt.Println(a, "")
                if !sp.NextPermutation(sort.IntSlice(a)) {
                        break
                }
        }

        fmt.Println("Combination:")
        for e := range sp.Combinations(a, 2, 1) {
                fmt.Println(e, "")
        }
}
```