package util

func GenIntSeries(n int, min int, max int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = RandomRange(min, max)
	}

	return result
}
