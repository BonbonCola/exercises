package main

import (
	"exercises/intset/lib"
	"fmt"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(3)

	y.Add(4)

	fmt.Println(x.String())
	fmt.Println(y.Len())
}

// 1 2 3 4
// 1 0 1 0
