package main

import (
	"fmt"
	"os"
	// "math"
	//    "strconv"
	//    "reflect"
)

func main() {
//	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, "  ", os.Args[i])
	}
	
//	fmt.Println(os.Args)
//	fmt.Println("test")

}
