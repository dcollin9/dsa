package helpers

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// buildTree builds the binary search tree from the array representation.
func BuildTree(arr []int) *TreeNode {
	// If the array is empty or has only one element, return nil as there are no valid nodes.
	if len(arr) < 2 {
		return nil
	}

	// Helper function to recursively build the tree.
	var helper func(index int) *TreeNode
	helper = func(index int) *TreeNode {
		// If index is out of bounds or the current value is -1, return nil.
		if index >= len(arr) || arr[index] == -1 {
			return nil
		}

		// Create a new TreeNode with the current value.
		node := &TreeNode{Val: arr[index]}
		// Recursively set the left and right children.
		node.Left = helper(index * 2)    // Left child is at index * 2
		node.Right = helper(index*2 + 1) // Right child is at index * 2 + 1

		return node
	}

	// Start building the tree from index 1, as 0th index is ignored.
	return helper(1)
}

// printTree performs an in-order traversal of the tree to print values.
func PrintTree(node *TreeNode) {
	if node == nil {
		return
	}
	PrintTree(node.Left)
	fmt.Println(node.Val, " ")
	PrintTree(node.Right)
}
