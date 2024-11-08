package main

import "fmt"

func main() {
	low := 1
	high := 100
	val := binarySearchRange(low, high)
	fmt.Println(val)

	low2 := 1
	high2 := 9
	val2 := binarySearchRange(low2, high2)
	fmt.Println(val2)
}

func binarySearchRange(low, high int) int {
	for low <= high {
		mid := (low + high) / 2
		if isCorrect(mid) > 0 {
			high = mid - 1
		} else if isCorrect(mid) < 0 {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func isCorrect(n int) int {
	if n > 10 {
		return 1
	} else if n < 10 {
		return -1
	} else {
		return 0
	}
}
