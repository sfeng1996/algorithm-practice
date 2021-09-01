package main

import (
	"fmt"
	"math"
)

// 给定二叉树，使用深度优先求出最小深度，最小深度表示所有叶子节点到根节点的深度最小值

// 定义二叉树结构体
type BinaryTree struct {
	Data int
	Left *BinaryTree
	Right *BinaryTree
}

func minDepthByDepthFirst(binaryTree *BinaryTree) int {
	// 如果给定树为空，则返回0
	if binaryTree == nil {
		return 0
	}
	// 叶子节点没有左子树和右子树
	if binaryTree.Left == nil && binaryTree.Right == nil {
		return 1
	}
	// 先定义最小值
	min := math.MaxInt64
	// 如果左子树不为空，说明该节点不是叶子节点，递归求最小值
	if binaryTree.Left != nil {
		min = minNum(minDepthByDepthFirst(binaryTree.Left), min)
	}
	if binaryTree.Right != nil {
		min = minNum(minDepthByDepthFirst(binaryTree.Right), min)
	}
	// min没有加上当前的节点，所以加1
	return min + 1
}

func minNum(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}


func main() {
	// 构造二叉树
	node7 := &BinaryTree{Data: 7}
	node6 := &BinaryTree{Data: 6, Left: node7}
	node5 := &BinaryTree{Data: 5}
	node4 := &BinaryTree{Data: 4}
	node3 := &BinaryTree{Data: 3, Left: node6}
	node2 := &BinaryTree{Data: 2, Left: node4, Right: node5}
	node1 := &BinaryTree{Data: 1, Left: node2, Right: node3}
	fmt.Println(minDepthByDepthFirst(node1))
}