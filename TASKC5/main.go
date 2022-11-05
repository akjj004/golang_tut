package main

import (
	"fmt"
	"mat/mat"
)

func main() {

	numbers := []int{10, 2, 24, 13, 20}
	min, max := mat.MinMax(numbers)
	fmt.Println("Min: ", min, "Max: ", max)
	mat.Test()
}
