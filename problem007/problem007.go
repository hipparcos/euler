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

func primeN(n int) int {
	p := 2
	for i := 2; n > 0; i++ {
		if isPrime(i) {
			n--
			p = i
		}
	}
	return p
}

// What is the 10 001st prime number?
func main() {
	var n int
	flag.IntVar(&n, "n", 10001, "nth prime number")
	flag.Parse()

	fmt.Println(primeN(n))
}
