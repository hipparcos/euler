package main

import (
	"flag"
	"fmt"
)

func sumFibEven(limit int) (sum int) {
	n_1 := 1
	// Iterate through Fibonacci sequence terms.
	for n := 1; n < limit; n_1, n = n, n_1+n {
		// Check parity.
		if (n & 0x1) == 0 {
			sum += n
		}
	}
	return
}

// By considering the terms in the Fibonacci sequence whose values do not exceed
// four million, find the sum of the even-valued terms.
func main() {
	var limit int
	flag.IntVar(&limit, "limit", 4000000, "upper limit of fibonacci values")
	flag.Parse()

	fmt.Println(sumFibEven(limit))
}
