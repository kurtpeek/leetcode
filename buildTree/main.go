package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	var j int
	for i, val := range inorder {
		if val == preorder[0] {
			j = i
			break
		}
	}

	root.Left = buildTree(preorder[1:j+1], inorder[:j])
	root.Right = buildTree(preorder[j+1:], inorder[j+1:])

	return root
}
