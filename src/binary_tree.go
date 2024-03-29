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

// returns is two trees have the same structure and values
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// checks whether the tree is a mirror of itself i.e. symmetric around it's center
/*

Example 1: Returns true

     1
  2     2
3   4 4   3


Example 2: Returns false

       1
    2     2
	   3     3


*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)

}

func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}
