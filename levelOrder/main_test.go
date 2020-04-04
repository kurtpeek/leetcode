package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	assert.Exactly(t, [][]int{
		[]int{3},
		[]int{9, 20},
		[]int{15, 7},
	}, levelOrder(root))
}

func TestLevelOrderNilRoot(t *testing.T) {
	assert.Exactly(t, [][]int{}, levelOrder(nil))
}
