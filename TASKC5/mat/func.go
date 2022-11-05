package mat

func MinMax(tab []int) (int, int) {
	if len(tab) < 1 {
		return 0, 0
	}
	min, max := tab[0], tab[0]
	for i := 1; i < len(tab); i++ {
		if tab[i] < min {
			min = tab[i]
		}
		if tab[i] > max {
			max = tab[i]
		}
	}
	return min, max
}
