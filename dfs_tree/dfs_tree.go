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

	// delete a value from the tree
	deleteFromBST(root, 2)
	found3 := searchTree(root, 2)
	fmt.Println(found3)
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
// Situations to consider: a node has 0 or 1 child. A target node has 2 children
// have to handle all situations
// Time Complexity: O(logn)
// Space Complexity: O(logn) (for the call stack)
func deleteFromBST(root *helpers.TreeNode, val int) *helpers.TreeNode {
	if root == nil {
		return nil
	}

	if val > root.Val {
		root.Right = deleteFromBST(root.Right, val)
	} else if val < root.Val {
		root.Left = deleteFromBST(root.Left, val)
	} else {
		// we're returning the value here. We're deleting no matter what.
		// in the case that we call deleteFromBST, we will traverse until we get to that leaf node,
		// and then set it's parent's root.Left/root.Right to nil, because that's what this current node's
		// children will be
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// The inorder successor is the left-most node in the right subtree of the target node
			minNode := MinValueNode(root.Right)
			root.Val = minNode.Val
			root.Right = deleteFromBST(root.Right, minNode.Val)
		}
	}

	return root
}

func MinValueNode(root *helpers.TreeNode) *helpers.TreeNode {
	curr := root
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}
