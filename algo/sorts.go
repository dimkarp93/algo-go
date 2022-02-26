package algo

import (
	"algo/util"
	"math"
)

func SelectSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	result := util.ArrayCopy(arr)

	for i := 0; i < len(result); i++ {
		min := result[i]
		mi := i

		for j := i + 1; j < len(result); j++ {
			if result[j] < min {
				min = result[j]
				mi = j
			}
		}

		t := result[i]
		result[i] = result[mi]
		result[mi] = t
	}

	return result
}

func InsertSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var result []int

	for _, v := range arr {
		result = append(result, v)
		index := -1
		for i, x := range result {
			if x > v {
				index = i
				break
			}
		}

		if index == 0 {
			temp := []int{v}
			result = append(temp, result[:len(result)-1]...)
		} else if index > 0 {
			temp := []int{v}
			temp = append(temp, result[index:len(result)-1]...)
			result = append(result[:index], temp...)
		}

	}

	return result
}

func BubbleSort(arr []int) []int {
	result := util.ArrayCopy(arr)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1; j++ {
			if result[j] > result[j+1] {
				t := result[j]
				result[j] = result[j+1]
				result[j+1] = t
			}
		}
	}

	return result
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	midIdx := len(arr) / 2

	t1 := MergeSort(arr[:midIdx])
	t2 := MergeSort(arr[midIdx:])

	var result []int

	i := 0
	j := 0
	for i < len(t1) || j < len(t2) {
		if i >= len(t1) {
			r := append(result, t2[j:]...)
			return r
		}

		if j >= len(t2) {
			r := append(result, t1[i:]...)
			return r
		}

		if t1[i] < t2[j] {
			result = append(result, t1[i])
			i++
		} else {
			result = append(result, t2[j])
			j++
		}
	}

	return result
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	result := util.ArrayCopy(arr)

	quickSort0(result, 0, len(result)-1)

	return result
}

func quickSort0(arr []int, l int, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	x := arr[mid]
	i := l
	j := r

	for i <= j {
		for arr[i] < x && i <= j {
			i++
		}
		for arr[j] > x && i <= j {
			j--
		}

		if i > j {
			break
		}

		t := arr[i]
		arr[i] = arr[j]
		arr[j] = t

		i++
		j--
	}

	quickSort0(arr, l, j)
	quickSort0(arr, i, r)
}

func HeapSort(arr []int) []int {
	heap := util.ArrayCopy(arr)

	for i := (len(heap) / 2) + 1; i >= 0; i-- {
		heapify(heap, i)
	}

	var result []int
	for len(heap) > 0 {
		v := heap[0]
		result = append(result, v)

		heap[0] = heap[len(heap)-1]
		heap[len(heap)-1] = v

		heap = heap[:len(heap)-1]

		heapify(heap, 0)

	}

	return result
}

func heapify(heap []int, j int) {
	mi := -1

	for j < len(heap) {
		l := 2*j + 1
		r := 2*j + 2

		if l < len(heap) && r < len(heap) {
			if heap[l] < heap[r] {
				mi = l
			} else {
				mi = r
			}
		} else if l < len(heap) {
			mi = l
		} else {
			break
		}

		if heap[j] > heap[mi] {
			t := heap[j]
			heap[j] = heap[mi]
			heap[mi] = t
			j = mi
		} else {
			break
		}
	}
}

func CounterSort(arr []int) []int {
	result := make([]int, len(arr))

	min := math.MaxInt
	max := math.MinInt

	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	length := max - min + 1
	dict := make([]int, length)

	for _, v := range arr {
		dict[v-min]++
	}

	idx := 0
	for i, v := range dict {
		for j := 0; j < v; j++ {
			result[idx] = i + min
			idx++
		}
	}

	return result
}

func RadixSort(arr []int) []int {
	return radixSort0(arr, radixSortPositiveOpt)
}

func radixSort0(arr []int, sort func(arr []int) []int) []int {
	var pos []int
	var neg []int

	for _, v := range arr {
		if v >= 0 {
			pos = append(pos, v)
		} else {
			neg = append(neg, -v)
		}
	}

	posResult := sort(pos)
	negResult := sort(neg)

	result := make([]int, len(negResult))

	for i := len(negResult) - 1; i >= 0; i-- {
		result[len(negResult)-1-i] = -negResult[i]
	}

	result = append(result, posResult...)

	return result
}
func radixSortPositiveOpt(arr []int) []int {
	limit := 33
	length := len(arr)

	result := util.ArrayCopy(arr)

	layer := [][]int{make([]int, length), make([]int, length)}
	indexes := []int{0, 0}

	for i := 0; i <= limit; i++ {
		indexes[0] = 0
		indexes[1] = 0

		hasOnlyZeroRest := true
		for _, v := range result {
			rest := getRest(v, i)
			if rest != 0 {
				hasOnlyZeroRest = false
			}
			idx := getShift(v, i)
			layer[idx][indexes[idx]] = v
			indexes[idx]++
		}

		if hasOnlyZeroRest {
			return result
		}

		idx := 0
		for i := 0; i < 2; i++ {
			for j := 0; j < indexes[i]; j++ {
				result[idx] = layer[i][j]
				idx++
			}
		}
	}

	return result
}

func getShift(x, idx int) int {
	return (x >> idx) & 1
}

func getRest(x, idx int) int {
	return x >> idx
}
