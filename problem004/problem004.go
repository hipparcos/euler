package main

import (
	"flag"
	"fmt"
)

// IsPalindromic tells if a number is a palindrome in base 10.
func IsPalindromic(n int) bool {
	var invn int
	for q := n; q > 0; q /= 10 {
		invn = invn*10 + q%10
	}
	return n == invn
}

// palindricNumbers returns the list of palindromic numbers obtained by
// multiplying 2 numbers between loBound & upBound.
func palindricNumbers(loBound, upBound int) (ps []int) {
	for i := upBound; i >= loBound; i-- {
		for j := i; j >= loBound; j-- {
			p := i * j
			if IsPalindromic(p) {
				ps = append(ps, p)
			}
		}
	}
	return
}

// maxOfSLiceOfInt returns the maximum of a slice of int.
func maxOfSLiceOfInt(ns []int) (max int) {
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return
}

func main() {
	var upBound, loBound int
	flag.IntVar(&loBound, "lower", 100, "lower bound")
	flag.IntVar(&upBound, "upper", 999, "upper bound")
	flag.Parse()

	ps := palindricNumbers(loBound, upBound)
	largest := maxOfSLiceOfInt(ps)
	fmt.Println(largest)
}
