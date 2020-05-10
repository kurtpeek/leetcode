package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	assert.Equal(t, root, buildTree(inorder, postorder))
}

var root = &TreeNode{
	Val:   3,
	Left:  &TreeNode{Val: 9},
	Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
}

func TestPreorder(t *testing.T) {
	assert.Exactly(t, []int{3, 9, 20, 15, 7}, preOrder(root))
}

func TestInorder(t *testing.T) {
	assert.Exactly(t, []int{9, 3, 15, 20, 7}, inOrder(root))
}

func TestPostOrder(t *testing.T) {
	assert.Exactly(t, []int{9, 15, 7, 20, 3}, postOrder(root))
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	preorder := []int{root.Val}
	preorder = append(preorder, preOrder(root.Left)...)
	preorder = append(preorder, preOrder(root.Right)...)
	return preorder
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	inorder := []int{root.Val}
	inorder = append(inOrder(root.Left), inorder...)
	inorder = append(inorder, inOrder(root.Right)...)
	return inorder
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	postorder := []int{root.Val}
	postorder = append(postOrder(root.Right), postorder...)
	postorder = append(postOrder(root.Left), postorder...)
	return postorder
}
