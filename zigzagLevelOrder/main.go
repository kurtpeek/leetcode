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

func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := list.New()
	if root != nil {
		queue.PushBack(&Node{TreeNode: root, Depth: 0})
	}

	depth := 0
	result := make([][]int, 0)
	for queue.Len() > 0 {
		var e *list.Element
		if depth%2 == 0 {
			e = queue.Front()
		} else {
			e = queue.Back()
		}
		node := e.Value.(*Node)
		if node.Depth > depth {
			depth = node.Depth
			continue
		}

		queue.Remove(e)

		if node.Depth > len(result)-1 {
			result = append(result, make([]int, 0))
		}
		result[len(result)-1] = append(result[len(result)-1], node.TreeNode.Val)

		if node.Depth%2 == 0 {
			if node.TreeNode.Left != nil {
				queue.PushBack(&Node{TreeNode: node.TreeNode.Left, Depth: node.Depth + 1})
			}
			if node.TreeNode.Right != nil {
				queue.PushBack(&Node{TreeNode: node.TreeNode.Right, Depth: node.Depth + 1})
			}
		} else {
			if node.TreeNode.Right != nil {
				queue.PushFront(&Node{TreeNode: node.TreeNode.Right, Depth: node.Depth + 1})
			}
			if node.TreeNode.Left != nil {
				queue.PushFront(&Node{TreeNode: node.TreeNode.Left, Depth: node.Depth + 1})
			}
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

	fmt.Println(zigzagLevelOrder(root))
}
