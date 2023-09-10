package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(values []int, node *TreeNode) []int {
	if node == nil {
		return values
	}

	if node.Left == nil && node.Right == nil {
		values = append(values, node.Val)
		return values
	}

	values = append(values, traverse([]int{}, node.Left)...)
	values = append(values, node.Val)
	values = append(values, traverse([]int{}, node.Right)...)

	return values
}

// for leetcode
// returns the inorder traversal of its node values
// inorder traversal looks at Left then Node then Right values
func inorderTraversal(root *TreeNode) []int {
	values := []int{}

	if root == nil {
		return values
	}

	values = append(values, traverse([]int{}, root.Left)...)
	values = append(values, root.Val)
	values = append(values, traverse([]int{}, root.Right)...)

	return values
}
