package main

import "fmt"

func minmax(tab []int) (int, int) {
	if len(tab) < 1 {
		return 0, 0
	}
	a, b := tab[0], tab[0]
	for i := 0; i < len(tab); i++ {
		if a > tab[i] {
			a = tab[i]
		}

		if b < tab[i] {
			b = tab[i]
		}

	}
	return a, b
}

func main() {
	// i, j, k := 5, true, "test"
	// fmt.Println(i, j, k)
	numbers := []int{10, 2, 24, 13, 20}
	a, b := minmax(numbers)
	fmt.Println("Min: ", a, "Max: ", b)

}
