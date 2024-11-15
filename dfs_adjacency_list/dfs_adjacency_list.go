package main

import (
	"fmt"
)

func main() {
	edges := [][]string{{"A", "B"}, {"B", "C"}, {"B", "E"}, {"C", "E"}, {"E", "D"}}
	adjList := buildAdjList(edges)
	visited := make(map[string]bool, len(adjList))
	numPaths := dfs(adjList, "A", "D", visited)
	fmt.Println(numPaths) // should be 2
}

// count the number of paths that lead from a source to destination
func dfs(adjList map[string][]string, node, target string, visited map[string]bool) int {
	if active, ok := visited[node]; ok && active {
		return 0
	}
	if node == target {
		return 1
	}

	// iterate through children and call dfs on them.
	// aggregate counts
	count := 0
	visited[node] = true
	for _, v := range adjList[node] {
		count += dfs(adjList, v, target, visited)
	}

	// now remove from visited
	visited[node] = false
	return count
}

// are edges directional or bidirectional? assume directional
func buildAdjList(edges [][]string) map[string][]string {
	adjList := make(map[string][]string)

	for _, v := range edges {
		adjList[v[0]] = append(adjList[v[0]], v[1]) // append safely handles nil slices
	}

	return adjList
}
