package main

import (
	"fmt"
	"math"
)

func main() {
	N := 18
	K := 2

	all := 0
	plays := 0
	for all < K && N > 2 {
		if math.Mod(float64(N), 2) == 0 {

			all++
			plays++
			N = N / 2
			fmt.Printf("all in at %d\n", N)
		} else {
			plays++
			N -= 1
		}
	}
	for N > 1 {
		plays++
		N -= 1
	}
	fmt.Println(plays)
}

