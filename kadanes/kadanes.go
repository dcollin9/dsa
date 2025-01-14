package main

import "fmt"

func main() {
	arr := []int{2, -3, 4, -2, 2, 1, -1, 4}
	maxSubarraySum := kadanesAlg(arr)
	fmt.Println(maxSubarraySum) // should be 8
}

// be greedy
// calculate max subarray sum in O(n) time
func kadanesAlg(arr []int) int {
	maxSum := arr[0]
	currSum := 0

	for i := 0; i < len(arr); i++ {
		currSum = Max(currSum, 0) // if we have a negative value here, we will never be able to add that to our currSum and get a greater value, so we take 0 instead
		currSum += arr[i]
		maxSum = Max(currSum, maxSum)
	}

	return maxSum
}

func Max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
