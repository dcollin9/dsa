package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 7, 8, 4}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// Time Complexity: O(nlogn)
// Space Complexity: O(logn) call stack of mergeSort + creating O(n) arrays in merge
// So space complexity: O(n + logn) => O(n)

// this function will get called O(logn) times.
// the merge call inside of it is an O(n), on average, time complexity
func mergeSort(arr []int, s, e int) {
	// if s and e are the same or s > e, return
	if e-s+1 <= 1 {
		return
	}

	m := (e + s) / 2
	mergeSort(arr, s, m)
	mergeSort(arr, m+1, e)

	merge(arr, s, e, m)

}

// Time complexity here could technically be O(2n), since it takes O(n) to copy the arrays
// and O(n) to then iterate through both arrays
func merge(arr []int, s, e, m int) {
	// The arrays here HAVE to be created to specific length for the copy function.
	// Otherwise the copy function won't do anything bc it goes only to the lenght of the
	// shorter array, which in this case is size 0
	lArr := make([]int, m-s+1)
	rArr := make([]int, e-m)

	copy(lArr, arr[s:m+1])
	copy(rArr, arr[m+1:e+1])

	l := 0
	r := 0

	for l < len(lArr) && r < len(rArr) {
		if lArr[l] <= rArr[r] { // the equals sign here makes it stable
			arr[s] = lArr[l]
			l++
		} else {
			arr[s] = rArr[r]
			r++
		}
		s++
	}

	for l < len(lArr) {
		arr[s] = lArr[l]
		l++
		s++
	}

	for r < len(rArr) {
		arr[s] = rArr[r]
		r++
		s++
	}
}
