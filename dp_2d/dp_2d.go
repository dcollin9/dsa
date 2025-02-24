package main

import "fmt"

// Problem:
// Given a matrix m * n, count the number of unique paths from the top left to the bottom right.
// You are only allowed to move down or to the right.
// M and N represent the rows and columns size, respectively.

func main() {
	m, n := 4, 4
	uniquePaths := recursion(m, n, 0, 0)
	fmt.Println("Recursion Brute Force: ", uniquePaths) // should be 20
	memo := make(map[string]int, m*n)
	uniquePaths = recursionWithMemo(m, n, 0, 0, memo)
	fmt.Println("Recursion with Memo: ", uniquePaths)
}

// brute force solution
// Time complexity: O(2^(n + m)) bc at every node, we can double the nodes we will visit from that node
// Space Complexity: O(n + m) bc the call stack/height of decision tree is at most n + m since we don't
// visit every cell, bc we just move down and right
func recursion(m, n, r, c int) int {
	if r == m || c == m {
		return 0
	}
	if r == m-1 && c == n-1 {
		return 1
	}

	return recursion(m, n, r+1, c) + recursion(m, n, r, c+1)
}

// recursion with memoization (top down dp solution)
func recursionWithMemo(m, n, r, c int, memo map[string]int) int {
	if r == m || c == n {
		return 0
	}
	if r == m-1 || c == n-1 {
		return 1
	}
	str := fmt.Sprint(r, c)
	if v, ok := memo[str]; ok {
		return v
	}

	memo[str] = recursionWithMemo(m, n, r+1, c, memo) + recursionWithMemo(m, n, r, c+1, memo)
	return memo[str]
}

// bottom up dp solution
