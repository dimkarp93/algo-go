package util

func FindMinimumOfPowerOf2ThanMore(n int) int {
	bits := SplitIntoBits(n)
	if IsPowerOf2(bits) {
		return n
	}

	return 1 << len(bits)
}

func IsPowerOf2(bits []int) bool {
	if len(bits) < 2 {
		return true
	}

	hasLeadingOne := bits[len(bits)-1] == 1
	hasNoLeadingOne := false

	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 1 {
			hasNoLeadingOne = true
		}
	}

	return hasLeadingOne && !hasNoLeadingOne
}

func SplitIntoBits(n int) []int {
	if n < 2 {
		return []int{n}
	}

	var result []int

	for ; n > 0; n >>= 1 {
		result = append(result, n&1)
	}

	return result
}
