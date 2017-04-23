package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	// "math"
	//    "strconv"
	//"reflect"
)

func main() {
	counts := make(map[string]int)
	//	names := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "error")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			//names = counts
			countLines(f, counts, arg)
			f.Close()
			//	fmt.Println(arg)

		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, arg string) {
	//names := make(map[string]int)
	input := bufio.NewScanner(f)
	i := 0
	//names = counts
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 2 {
			i++
		}

	}
	if i > 0 {
		fmt.Println(arg)
	}

}
