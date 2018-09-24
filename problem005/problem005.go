package main

import (
	"flag"
	"fmt"
	"math"
)

// factorize returns the prime factorization of n.
func Factorize(n int) (factors []int) {
	var limit int = int(math.Ceil(math.Sqrt(float64(n))))
	// Test 2.
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	// Test all odd numbers.
	for f := 3; f <= limit; f += 2 {
		for n%f == 0 {
			factors = append(factors, f)
			n /= f
		}
	}
	if len(factors) == 0 {
		// Is prime, add the number itself.
		factors = append(factors, n)
	}
	return
}

// Factors is a map of factors and their power: map{factor]power.
type Factors map[int]int

// NewFactors converts a list of factors to a map of factors and their power.
func NewFactors(factors []int) (mapf Factors) {
	mapf = make(map[int]int)
	for _, f := range factors {
		if _, ok := mapf[f]; ok {
			mapf[f]++
		} else {
			mapf[f] = 1
		}
	}
	return
}

// MergeMax merges ys into xs.
// If an element is present in both xs & yz, MergeMax keeps only the max.
func (fs *Factors) MergeMax(xs Factors) {
	for xf, xp := range xs {
		if p, ok := (*fs)[xf]; !ok || p < xp {
			(*fs)[xf] = xp
		}
	}
}

// Fold folds fs using fun.
func (fs *Factors) Fold(init int, fun func(acc int, key int, val int) int) (acc int) {
	acc = init
	for k, v := range *fs {
		acc = fun(acc, k, v)
	}
	return
}

// Product returns the product of all factors.
func (fs *Factors) Product() int {
	return (*fs).Fold(1, func(acc, key, val int) int {
		factor := 1
		for ; val > 0; val-- {
			factor *= key
		}
		return acc * factor
	})
}

// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
func main() {
	var upBound, loBound int
	flag.IntVar(&loBound, "lower", 1, "lower bound")
	flag.IntVar(&upBound, "upper", 20, "upper bound")
	flag.Parse()

	if upBound < loBound {
		loBound, upBound = upBound, loBound
	}

	factors := NewFactors([]int{})
	for n := loBound; n <= upBound; n++ {
		fs := NewFactors(Factorize(n))
		factors.MergeMax(fs)
	}
	fmt.Println(factors.Product())
}
