package main

import (
	"fmt"
	"math"
)

func main() {
	A := 6000
	B := 7000
	var best int
	var max int
	var i int
	for i = A; i <= B; i++ {
		count := 0
		j := int(math.Sqrt(float64(i)))
		z := i
		for j*j == z {
			fmt.Printf("found sqrt: %d\n", j)
			count++
			x := j
			y := int(math.Sqrt(float64(x)))
			if y*y != x {
				break
			}
			z = x
			j = y
		}
		if count > max {
			max = count
			best = i
			fmt.Printf("new best: %d\n", best)
		}
	}
	fmt.Printf("best: %d\n", best)
}

