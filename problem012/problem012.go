package main

import (
	"flag"
	"fmt"
	"math"
)

// divisors returns the list of all divisors of n.
func divisors(n int) []int {
	ds := make([]int, 0)
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			ds = append(ds, i)
			if i == n/i {
				continue
			}
			ds = append(ds, n/i)
		}
	}
	return ds
}

func triangularNumberGenerator() func() int {
	idx, val := 0, 0
	return func() int {
		idx++
		val += idx
		return val
	}
}

func firstTriangularNumberWithNDivisors(n int) int {
	triangle := triangularNumberGenerator()
	for {
		t := triangle()
		if len(divisors(t)) >= n {
			return t
		}
	}
}

// What is the value of the first triangle number to have over five hundred divisors?
func main() {
	var n int
	flag.IntVar(&n, "n", 500, "number of divisors")
	flag.Parse()

	fmt.Println(firstTriangularNumberWithNDivisors(n))
}
