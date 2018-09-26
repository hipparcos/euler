package main

import (
	"flag"
	"fmt"
)

func isPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n <= 3:
		return true
	case n%2 == 0:
		return false
	case n%3 == 0:
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// primes returns a slice containing all prime numbers up to limit.
func primes(limit int) (ps []int) {
	if limit < 2 {
		return
	}
	ps = append(ps, 2)
	for i := 3; i < limit; i += 2 {
		if isPrime(i) {
			ps = append(ps, i)
		}
	}
	return
}

func sum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

// Find the sum of all the primes below two million.
func main() {
	var limit int
	flag.IntVar(&limit, "limit", 2000000, "limit of prime summation")
	flag.Parse()

	ps := primes(limit)
	fmt.Println(sum(ps))
}
