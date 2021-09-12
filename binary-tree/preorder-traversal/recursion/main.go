package main

import "fmt"

// 二叉树前序遍历表示按照：根-左-右的顺序进行遍历
//    1
//  2   3
// 4 5 6
//    7
// 该树按照前序遍历，打印结果为：1,2,4,5,3,6,7

// 定义二叉树结构体
type BinaryTree struct {
	Data int
	Left *BinaryTree
	Right *BinaryTree
}

// 打印第一次成为栈顶的元素
func preorderTraversal(root *BinaryTree) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
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
	preorderTraversal(node1)
}