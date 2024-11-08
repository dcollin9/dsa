package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	found := binarySearch(arr, 9)
	fmt.Println("Found 1: ", found)
	found2 := binarySearch(arr, 13)
	fmt.Println("Found 2: ", found2)
}

func binarySearch(arr []int, target int) bool {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}
