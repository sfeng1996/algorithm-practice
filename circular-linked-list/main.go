package main

import "fmt"

// 判断给定链表是否有环，例如1->2->3->4->5->2，这个链表就是环形链表

// 定义链表
type LNode struct {
	Data int
	Next *LNode
}

// 使用双指针，定义两个指针，只要两个指针相等，说明有环
// 如果两个节点的话，第二个节点指向的值为nil，说明无环
// 如果快指针的下一个值为nil，说明无环
func judgeRingLink(head *LNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 慢指针为第一个节点
	slow := head
	// 快指针为第二个节点
	quick := head.Next
	for slow != quick {
		// 如果快指针的下一个节点为空，说明不存在环，返回false,或者只有一个几点，返回false
		if quick == nil || quick.Next == nil {
			return false
		}
		slow = slow.Next
		// quick必须必slow快两步，不然会无限循环，找不到重复值
		quick = quick.Next.Next
	}
	return true
}

func main() {
	node1 := &LNode{Data: 1}
	node2 := &LNode{Data: 2}
	node3 := &LNode{Data: 3}
	node4 := &LNode{Data: 4}
	node5 := &LNode{Data: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	//node5.Next = node2
	fmt.Println(judgeRingLink(node1))

}