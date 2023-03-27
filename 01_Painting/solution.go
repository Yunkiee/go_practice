package _1_Painting

func solution(n int, m int, section []int) int {
	count := 0
	lastPaintedIndex := 0
	for _, target := range section {
		if target > lastPaintedIndex {
			lastPaintedIndex = target + m - 1
			count++
		}
	}

	return count
}
