package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	var j int
	for i, val := range inorder {
		if val == postorder[len(postorder)-1] {
			j = i
			break
		}
	}

	root.Left = buildTree(inorder[:j], postorder[:j])
	root.Right = buildTree(inorder[j+1:], postorder[j:len(postorder)-1])

	return root
}
