package main

import "fmt"

/*
	104.二叉树的最大深度
	给定一个二叉树，找出其最大深度。
	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
	说明:叶子节点是指没有子节点的节点。
	示例：
	给定二叉树 [3,9,20,null,null,15,7]，
		3
	   / \
	  9  20
		/  \
	   15   7
	返回它的最大深度 3
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main()  {
	aa := TreeNode{}
	fmt.Println(maxDepth(&aa))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
