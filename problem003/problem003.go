package main

import (
	"flag"
	"fmt"
	"math"
)

// factorize returns the list all prime factors of n.
func factorize(n int) (factors []int) {
	// Test 2.
	if n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	// Test all odd numbers.
	var limit int = int(math.Sqrt(float64(n))) + 1
	for f := 3; f < limit; f += 2 {
		if n%f == 0 {
			factors = append(factors, f)
			n /= f
		}
	}
	return
}

// What is the largest prime factor of the number 600851475143?
func main() {
	var n int
	flag.IntVar(&n, "n", 600851475143, "number to factorize")
	flag.Parse()

	factors := factorize(n)
	fmt.Println(factors[len(factors)-1])
}
