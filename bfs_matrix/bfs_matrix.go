package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	shortestPath := bfs(grid)
	fmt.Println(shortestPath) // should be 7

}

// Find the length of the shortest path from the top left of the grid to the bottom right
// Can only traverse along 0s, not 1s

// Time Complexity: O(n * m) // we go to every element in the grid once in worst case
// Space Complexity: O(n * m) // visited grid, also the queue could have n * m elements in it
func bfs(grid [][]int) int {

	rows, cols := len(grid), len(grid[0])

	queue := [][]int{{0, 0}}
	visited := make([][]int, rows)
	for i := range rows {
		visited[i] = make([]int, cols)
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	pathLen := 1
	for len(queue) > 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			r := queue[i][0]
			c := queue[i][1]

			if r == rows-1 && c == cols-1 {
				return pathLen
			}

			// add next layer to the queue
			for _, v := range directions {
				newR := r + v[0]
				newC := c + v[1]
				if newR == rows || newR < 0 || newC == cols || newC < 0 || visited[newR][newC] == 1 || grid[newR][newC] == 1 {
					continue
				}

				queue = append(queue, []int{newR, newC})
				visited[r][c] = 1
			}

		}
		queue = queue[queueLen:]
		pathLen++
	}

	return pathLen
}
