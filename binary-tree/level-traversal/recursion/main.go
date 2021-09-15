package main

import "fmt"

// 二叉树层序遍历表示按照一层一层遍历
//    1
//  2   3
// 4 5 6
//    7
// 该树按照层序遍历，打印结果为：1,2,3,4,5,6,7

// 定义二叉树结构体

type BinaryTree struct {
	Data int
	Left *BinaryTree
	Right *BinaryTree
}

// 定义队列，先进先出


var result [][]int
func levelOrder(root *BinaryTree) {
	result = make([][]int, 0)
	if root == nil {
		return
	}
	dfsHelper(root, 0)
	for _, v := range result{
		fmt.Println(v)
	}
}

func dfsHelper(node *BinaryTree, level int) {
	if node == nil {
		return
	}
	if len(result) == level  {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], node.Data)
	dfsHelper(node.Left, level + 1)
	dfsHelper(node.Right, level + 1)
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
	levelOrder(node1)
}

