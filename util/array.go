package util

func ArrEquals(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, _ := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func ArrayCopy(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}
