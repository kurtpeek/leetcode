package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigZagLevelOrder(t *testing.T) {
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
		[]int{20, 9},
		[]int{15, 7},
	}, zigzagLevelOrder(root))
}

func TestZigZagLevelOrderNilRoot(t *testing.T) {
	assert.Exactly(t, [][]int{}, zigzagLevelOrder(nil))
}
