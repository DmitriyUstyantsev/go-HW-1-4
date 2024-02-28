package main

import (
	"fmt"
)

func mergeArrays(arr1 []int, arr2 []int) []int {
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
			merged[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}
	return merged
}

func main() {
	arr1 := []int{6, 3, 5, 4}
	arr2 := []int{1, 2, 7, 8, 9}
	merged := mergeArrays(arr1, arr2)
	fmt.Println(merged)
}
