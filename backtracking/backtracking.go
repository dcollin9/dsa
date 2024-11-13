package main

import (
	"dsa/helpers"
	"fmt"
)

func main() {
	arr := []int{0, 4, 0, 1, -1, 7, 2, 0}
	root := helpers.BuildTree(arr)
	res := CanReachLeaf(root)
	fmt.Println(res) // should be true here
}

// return true if the leaf can be reached from the root
// caveat - we can't have any 0s in our path

// Time Complexity: O(n) - we will visit in the worst case every node
// Space Complexity: O(h) - the call stack will never be more than height of tree
func CanReachLeaf(root *helpers.TreeNode) bool {
	if root == nil || root.Val == 0 {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	if CanReachLeaf(root.Left) {
		return true
	}
	if CanReachLeaf(root.Right) {
		return true
	}

	return false
}
