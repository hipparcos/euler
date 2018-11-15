package main

import (
	"flag"
	"fmt"
)

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func lenCollatz(start int) (length int) {
	n := start
	length = 1
	for n > 1 {
		length++
		n = nextCollatz(n)
	}
	return
}

// Longest Collatz sequence
// Which starting number, under one million, produces the longest chain?
func main() {
	var highest int
	flag.IntVar(&highest, "highest", 1000000, "highest starting number to test")
	flag.Parse()

	maxl, maxn := 1, 1
	for n := highest; n > 1; n-- {
		l := lenCollatz(n)
		if l > maxl {
			maxl = l
			maxn = n
		}
	}
	fmt.Println(maxn)
}
