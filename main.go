package main

func Solution(A []int) int {
	s := make(map[int]bool, len(A))

	for i := 1; i < len(A)+1; i++ {
		s[i] = false
	}

	for _, a := range A {
		s[a] = true
	}

	for i := 1; i < len(A)+1; i++ {
		if s[i] == false {
			return i
		}
	}
	return len(A) + 1
}
