package main

import (
	"testing"
)

func TestIsPalindromic(t *testing.T) {
	testcases := []struct {
		n int
		p bool
	}{
		{1, true},
		{100, false},
		{101, true},
	}
	for _, tt := range testcases {
		if IsPalindromic(tt.n) != tt.p {
			t.Errorf("Failed for %v", tt.n)
		}
	}
}
