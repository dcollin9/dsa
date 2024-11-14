package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	visit := make([][]int, len(grid))
	for i := range visit {
		visit[i] = make([]int, len(grid[0]))
	}

	paths := dfs(grid, 0, 0, visit)
	fmt.Println("Path Count: ", paths) // should be 2
}

// Count the unique paths from the top left to the bottom right.
// A single path may only move along 0s and can't visit the same cell more than once.
// Can only move down, left, right, or up.

// Base cases:
// 1) Out of bounds
// 2) We visit a place that's already been visited (us hash set to track)
// 3) We reach a blocked position

// Time Complexity: O(4^(n * m)) each node can have 4 children worst case n * m times
// Space complexity: O(n * m) - once for the call stack, once for visited
func dfs(grid [][]int, r, c int, visited [][]int) int {
	// in visited, we'll store the num of unique paths
	rows, cols := len(grid), len(grid[0])

	if r == rows || r < 0 || c == cols || c < 0 || grid[r][c] == 1 || visited[r][c] == 1 {
		return 0
	}

	if r == rows-1 && c == cols-1 {
		return 1
	}

	// mark it as visited so we don't revisit in the same path
	visited[r][c] = 1
	count := 0
	count += dfs(grid, r+1, c, visited)
	count += dfs(grid, r-1, c, visited)
	count += dfs(grid, r, c+1, visited)
	count += dfs(grid, r, c-1, visited)

	// unmark it so we can visit it again in a different path
	visited[r][c] = 0
	return count
}
