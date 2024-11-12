package main

import (
	"dsa/helpers"
	"fmt"
)

func main() {
	inputArr := []int{0, 3, 1, 4, -1, 2} // first element we just ignore
	root := helpers.BuildTree(inputArr)
	found := searchTree(root, 5)
	fmt.Println(found)

	// insert a new value and search for it
	insertIntoBST(root, 5)
	found2 := searchTree(root, 5)
	fmt.Println(found2)
}

// search a binary search tree
// Time Complexity: O(logn)
func searchTree(root *helpers.TreeNode, val int) bool {
	if root == nil {
		return false
	}

	if val > root.Val {
		return searchTree(root.Right, val)
	} else if val < root.Val {
		return searchTree(root.Left, val)
	} else {
		return true
	}
}

// insert into a binary search tree - maintain order
// easiest way is to insert into a leaf node
// Time Complexity: O(logn)
func insertIntoBST(root *helpers.TreeNode, val int) *helpers.TreeNode {
	if root == nil {
		return &helpers.TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else if val <= root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

// delete from a binary search tree - maintain order even in a delete
