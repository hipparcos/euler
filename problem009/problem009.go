package main

import (
	"flag"
	"fmt"
)

// mn returns a pair of integers suitable to generate a triple such as `a + b + c = target`.
// mn uses brute-force (formula obtained by replacing a, b & c by their m & n
// formulation in the constraint `a + b + c = target`).
func mn(target int) (m, n int) {
	t := target / 2
	for n = 1; n < t; n++ {
		for m = 1; m < t; m++ {
			if m*m+m*n == t {
				return
			}
		}
	}
	return 0, 0
}

// triple returns a Pythagorean triple using Euclid's formula.
func triple(m, n int) (a, b, c int) {
	a = m*m - n*n
	b = 2 * m * n
	c = m*m + n*n
	return
}

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
func main() {
	var target int
	flag.IntVar(&target, "target", 1000, "value of a + b + c")
	flag.Parse()

	a, b, c := triple(mn(target))
	fmt.Println(a * b * c)
}
