package util

import "testing"

func TestFindMinimumOfPowerOf2ThanMore(t *testing.T) {
	doTestMinimumPower(0, 0, t)
	doTestMinimumPower(1, 1, t)
	doTestMinimumPower(2, 2, t)
	doTestMinimumPower(4, 3, t)
	doTestMinimumPower(4, 4, t)
	doTestMinimumPower(8, 5, t)
	doTestMinimumPower(8, 6, t)
	doTestMinimumPower(8, 7, t)
	doTestMinimumPower(8, 8, t)
	doTestMinimumPower(16, 9, t)
	doTestMinimumPower(16, 10, t)
	doTestMinimumPower(16, 11, t)
	doTestMinimumPower(16, 12, t)
	doTestMinimumPower(16, 13, t)
	doTestMinimumPower(16, 14, t)
	doTestMinimumPower(16, 15, t)
	doTestMinimumPower(16, 16, t)

}

func TestIsPowerOf2(t *testing.T) {
	doTestPowerOf(true, []int{0}, t)
	doTestPowerOf(true, []int{1}, t)
	doTestPowerOf(true, []int{0, 1}, t)
	doTestPowerOf(false, []int{1, 1}, t)
	doTestPowerOf(true, []int{0, 0, 1}, t)
	doTestPowerOf(false, []int{1, 0, 1}, t)
	doTestPowerOf(false, []int{0, 1, 1}, t)
	doTestPowerOf(false, []int{1, 1, 1}, t)
	doTestPowerOf(true, []int{0, 0, 0, 1}, t)
	doTestPowerOf(false, []int{1, 0, 0, 1}, t)
	doTestPowerOf(false, []int{0, 1, 0, 1}, t)
	doTestPowerOf(false, []int{1, 1, 0, 1}, t)
	doTestPowerOf(false, []int{0, 0, 1, 1}, t)
	doTestPowerOf(false, []int{1, 0, 1, 1}, t)
	doTestPowerOf(false, []int{0, 1, 1, 1}, t)
	doTestPowerOf(false, []int{1, 1, 1, 1}, t)
	doTestPowerOf(true, []int{0, 0, 0, 0, 1}, t)
}

func TestSplitIntoBits(t *testing.T) {
	doTestBitArray([]int{0}, 0, t)
	doTestBitArray([]int{1}, 1, t)
	doTestBitArray([]int{0, 1}, 2, t)
	doTestBitArray([]int{1, 1}, 3, t)
	doTestBitArray([]int{0, 0, 1}, 4, t)
	doTestBitArray([]int{1, 0, 1}, 5, t)
	doTestBitArray([]int{0, 1, 1}, 6, t)
	doTestBitArray([]int{1, 1, 1}, 7, t)
	doTestBitArray([]int{0, 0, 0, 1}, 8, t)
	doTestBitArray([]int{1, 0, 0, 1}, 9, t)
	doTestBitArray([]int{0, 1, 0, 1}, 10, t)
	doTestBitArray([]int{1, 1, 0, 1}, 11, t)
	doTestBitArray([]int{0, 0, 1, 1}, 12, t)
	doTestBitArray([]int{1, 0, 1, 1}, 13, t)
	doTestBitArray([]int{0, 1, 1, 1}, 14, t)
	doTestBitArray([]int{1, 1, 1, 1}, 15, t)
	doTestBitArray([]int{0, 0, 0, 0, 1}, 16, t)
}

func doTestMinimumPower(expected int, n int, t *testing.T) {
	actual := FindMinimumOfPowerOf2ThanMore(n)

	if expected != actual {
		t.Errorf("actual '%v' is not equal expected '%v' for '%v'", actual, expected, n)
	}
}

func doTestPowerOf(expected bool, bitarray []int, t *testing.T) {
	actual := IsPowerOf2(bitarray)

	if actual && !expected {
		t.Errorf("check asserts '%v' is power of 2 but it is not", bitarray)
	}

	if !actual && expected {
		t.Errorf("check asserts '%v' is not power of 2 but it is", bitarray)
	}
}

func doTestBitArray(expected []int, n int, t *testing.T) {
	actual := SplitIntoBits(n)

	if len(expected) != len(actual) {
		t.Errorf("len of actual array '%v' is not equal len of expected array '%v'", actual, expected)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("actual '%v' is not equal expected '%v'", actual, expected)
		}
	}
}
