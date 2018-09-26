package main

import (
	"flag"
	"fmt"
)

func sumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func sumOfN(n int) int {
	return n * (n + 1) / 2
}

// Find the difference between the sum of the squares of the first one hundred
// natural numbers and the square of the sum.
func main() {
	var n int
	flag.IntVar(&n, "n", 100, "apply for the first nth natural numbers")
	flag.Parse()

	ss := sumOfSquares(n)
	sn := sumOfN(n)
	fmt.Println(sn*sn - ss)
}
