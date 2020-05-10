package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert.Nil(t, sortedArrayToBST([]int{}))
}

func TestSortedArrayToBST(t *testing.T) {
	root := &TreeNode{Val: 0, Left: &TreeNode{Val: -1}, Right: &TreeNode{Val: 1}}
	assert.Exactly(t, root, sortedArrayToBST([]int{-1, 0, 1}))
}

var root = &TreeNode{
	Val: 0,
	Left: &TreeNode{
		Val:  -3,
		Left: &TreeNode{Val: -10},
	},
	Right: &TreeNode{
		Val:  9,
		Left: &TreeNode{Val: 5},
	},
}

func TestSortedArrayToBST2(t *testing.T) {
	assert.Exactly(t, root, sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}
