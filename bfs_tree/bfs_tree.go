package main

import (
	"dsa/helpers"
	"fmt"
)

// return the level order traversal of a Tree's nodes
func main() {
	// create tree
	arr := []int{0, 3, 9, 20, -1, -1, 15, 7}
	root := helpers.BuildTree(arr)
	want := [][]int{{3}, {9, 20}, {15, 7}}
	got := bfs(root)
	fmt.Println("Got: ", got)
	fmt.Println("Want: ", want)
}

// Time Complexity: O(n) - we have to visit every node
// Space Complexity: O(n) - we store the entire level of the tree in the queue at a time
func bfs(root *helpers.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*helpers.TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		queueLen := len(queue)
		tmpArr := make([]int, queueLen)
		for i := 0; i < queueLen; i++ {
			// pop first element off the queue w/out removing
			// add its children to the queue
			// remove that element off the queue
			tmpArr[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, tmpArr)
		queue = queue[queueLen:] // pop all elements off here
	}

	return res
}
