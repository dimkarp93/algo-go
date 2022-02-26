package util

import "math/rand"

func RandomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
