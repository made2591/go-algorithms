package common

import (
	"math/rand"
	"time"
)

func Sorted(vals []int) bool {
	for i, n := range vals[:len(vals)-1] {
		if n > vals[i+1] {
			return false
		}
	}
	return true
}

func Shuffle(vals []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}