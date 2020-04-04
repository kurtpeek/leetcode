package main

import (
	"container/list"
	"fmt"
)

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node ...
type Node struct {
	TreeNode *TreeNode
	Depth    int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	queue := list.New()
	queue.PushBack(&Node{TreeNode: root, Depth: 0})

	result := [][]int{make([]int, 0)}
	for queue.Len() > 0 {
		e := queue.Front()
		node := e.Value.(*Node)
		queue.Remove(e)

		if node.Depth > len(result)-1 {
			result = append(result, make([]int, 0))
		}
		result[len(result)-1] = append(result[len(result)-1], node.TreeNode.Val)

		if node.TreeNode.Left != nil {
			queue.PushBack(&Node{TreeNode: node.TreeNode.Left, Depth: node.Depth + 1})
		}
		if node.TreeNode.Right != nil {
			queue.PushBack(&Node{TreeNode: node.TreeNode.Right, Depth: node.Depth + 1})
		}
	}
	return result
}

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(levelOrder(root))
}
