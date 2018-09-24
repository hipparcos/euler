package main

import (
	"flag"
	"fmt"
)

func sumMultiple(multiple int, limit int) (sum int) {
	// n is the number of multiple below limit.
	n := (limit - 1) / multiple
	// sum is the sum of all integers from 1 to n, multiplied by multiple.
	sum = multiple * (n * (n + 1) / 2)
	return
}

// Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	var limit int
	flag.IntVar(&limit, "limit", 1000, "upper limit of multiples")
	flag.Parse()

	var sum int
	// Inclusion.
	sum += sumMultiple(3, limit)
	sum += sumMultiple(5, limit)
	// Exclusion.
	sum -= sumMultiple(15, limit)

	fmt.Println(sum)
}
