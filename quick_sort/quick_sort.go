package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 7, 8, 4}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// Time Complexity: O(nlogn) (assuming a good pivot, O(n^2) with a bad pivot)
// Space Complexity: O(logn) (assuming a good pivot, O(n) with a bad pivot)
func quickSort(arr []int, s, e int) {
	if e-s+1 <= 1 {
		return
	}

	lastInsert := s
	pivot := arr[e]
	for i := s; i < e; i++ {
		if arr[i] < pivot {
			arr[lastInsert], arr[i] = arr[i], arr[lastInsert]
			lastInsert++
		}
	}

	// switch pivot to correct place
	arr[e] = arr[lastInsert]
	arr[lastInsert] = pivot

	quickSort(arr, s, lastInsert-1)
	quickSort(arr, lastInsert+1, e)
}
