package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		Name     string
		Values   []int
		Expected int
	}{
		{"test1 -> 5", []int{1, 3, 6, 4, 1, 2}, 5},
		{"test2 -> 1", []int{-1, -3}, 1},
		{"test3 -> 1", generateRandomSlice(100000), 1},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			r := Solution(tc.Values)
			if r != tc.Expected {
				t.Errorf("%d != %d", r, tc.Expected)
			} else {
				t.Logf("%d == %d", r, tc.Expected)
			}
		})
	}
}

func generateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		x := rand.Intn(999) - rand.Intn(999)
		if x != 1 {
			slice[i] = x
		} else {
			slice[i] = x + rand.Intn(999)
		}
	}
	return slice
}
