package main

import "fmt"

// Find the nth fibonacci number
// n is the number at position n (O-indexed, so 4 would be 5th number)
// f(n) = f(n - 1) + f(n - 2)
func main() {
	res := fibonacciBruteForce(4) // should be 3
	fmt.Println(res)

	m := make(map[int]int, 4)
	res2 := fibonacciWMemoization(4, m) // should be 3
	fmt.Println(res2)

	res3 := fibonacciDP(4) // should be 3
	fmt.Println(res3)
}

// we can use recursion
// for the recursion version, this assumes I don't have n - 1 or n - 2. I only know n, the number of elements I have
// Time Complexity: O(2^n). In the decision tree, at each level (max n levels), we increase by a factor of 2
// Space Complexity: O(n). The call stack size will never be greater than n (height of decision tree)
func fibonacciBruteForce(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciBruteForce(n-1) + fibonacciBruteForce(n-2)
}

// topdowndp is called that for a reason, we start at the end goal and go back from there.
// so we start at the number we want (n) and work back from it until we get to 0
// Time Complexity: O(n)
// Space complexity: O(n)
func fibonacciWMemoization(n int, m map[int]int) int {
	if n <= 1 {
		return n
	} else if v, ok := m[n]; ok {
		return v
	}

	m[n] = fibonacciWMemoization(n-1, m) + fibonacciWMemoization(n-2, m)
	return m[n]
}

// We could calculate this iteratively... we don't need to use recursion per se, right?
// Turns out, the iterative approach is the dynamic programming approach
// Time Complexity: O(n)
// Space Complexity: O(1)
func fibonacciDP(n int) int {
	if n < 2 {
		return n
	}
	n1, n2 := 0, 1
	for i := 2; i <= n; i++ {
		tmp := n1
		n1 = n2
		n2 = tmp + n2
	}
	return n2
}
