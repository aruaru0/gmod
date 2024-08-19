# Graph
---
## DSU: Unionfind
Disjoint Set Union: Union Find Tree

`func NewDsu(n int) *DSU`
Initialize with number of elements n and return *DSU.

`func (d DSU) Merge(a, b int) int`
Merge elements of **a** and **b** into one set

`func (d DSU) Same(a, b int) bool`
Checks whether **a** and **b** are included in the same set. Returns true if included

`func (d DSU) Leader(a int) int`
Returns the representative of the set containing **a**.

`func (d DSU) Size(a int) int`
Returns the size of the set containing **a**


`func (d DSU) Groups() [][]int`
Returns the elements of each set


```go
package main

import (
	"fmt"

	gr "github.com/aruaru0/gmod/graph"
)

func dsuTest() {
	uf := gr.NewDsu(10)
	uf.Merge(0, 1)
	uf.Merge(1, 2)
	uf.Merge(5, 6)
	fmt.Println(uf.Same(0, 2)) // true
	fmt.Println(uf.Same(0, 5)) // false
	fmt.Println(uf.Groups())
}

func main() {
	dsuTest()
}

```