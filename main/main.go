package main

import (
	"algo/algo"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("arr: '%v', '%v', '%v'\n", arr[:2], arr[2], arr[2:])

	fmt.Printf("5 >> 0 = %v\n", 5>>0)
	fmt.Printf("5 >> 10 = %v\n", 5>>10)
	fmt.Printf("5 >> 100 = %v\n", 5>>100)
	fmt.Printf("-1 >> 0 = %v\n", -1>>0)
	fmt.Printf("-1 >> 1 = %v\n", -1>>1)
	fmt.Printf("-1 >> 10 = %v\n", -1>>10)
	fmt.Printf("-1 >> 100 = %v\n", -1>>100)

	algo.RadixSortOpt(arr)
}
