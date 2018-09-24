package main

import (
	"reflect"
	"testing"
)

func TestFactorize(t *testing.T) {
	testcases := []struct {
		n       int
		factors []int
	}{
		{2, []int{2}},
		{7, []int{7}},
		{20, []int{2, 2, 5}},
	}
	for _, tt := range testcases {
		got := Factorize(tt.n)
		if !reflect.DeepEqual(got, tt.factors) {
			t.Errorf("Fails for %v, expected %v, got %v", tt.n, tt.factors, got)
		}
	}
}

func TestNewFactors(t *testing.T) {
	testcases := []struct {
		asslice []int
		asmap   Factors
	}{
		{[]int{2}, map[int]int{2: 1}},
		{[]int{2, 2, 5}, map[int]int{2: 2, 5: 1}},
	}
	for _, tt := range testcases {
		got := NewFactors(tt.asslice)
		if !reflect.DeepEqual(got, tt.asmap) {
			t.Errorf("Fails for %v, expected %v, got %v", tt.asslice, tt.asmap, got)
		}
	}
}

func TestMergeMax(t *testing.T) {
	testcases := []struct {
		xs  Factors
		ys  Factors
		exp Factors
	}{
		{map[int]int{2: 1}, map[int]int{5: 1}, map[int]int{2: 1, 5: 1}},
		{map[int]int{2: 1}, map[int]int{2: 3}, map[int]int{2: 3}},
	}
	for _, tt := range testcases {
		tt.xs.MergeMax(tt.ys)
		if !reflect.DeepEqual(tt.xs, tt.exp) {
			t.Errorf("Fails, expected %v, got %v", tt.exp, tt.xs)
		}
	}
}

func TestFold(t *testing.T) {
	testcases := []struct {
		factors Factors
		product int
	}{
		{map[int]int{2: 1}, 2},
		{map[int]int{2: 2, 5: 1}, 20},
	}
	for _, tt := range testcases {
		got := tt.factors.Product()
		if got != tt.product {
			t.Errorf("Fails for %v, expected %v, got %v", tt.factors, tt.product, got)
		}
	}
}
