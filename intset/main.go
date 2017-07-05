package main

import (
	"exercises/intset/lib"
	"fmt"
)

func main() {
	var x intset.IntSet
	x.Add(1)
	x.Add(3)
	x.Add(4)

	fmt.Println(x.String())
	fmt.Println(x.Len())
	x.Remove(3)
	fmt.Println(x.String())
}

// 1 2 3 4
// 1 0 1 0
// 1
