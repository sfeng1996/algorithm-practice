package main


import "fmt"

// 链表是一种数据结构，里面的每个元素都包含下一个元素的位置信息，和数组做个对比，数组在内存中存放需要一段连续的位置，而数组则不用，可以分开存储在内存的任意位置。
// 这样做的好处是插入和删除速度快，步骤少，如果要在头部插入一个新的元素，链表只需要将第一个元素的位置信息添加进新的元素里即可，操作步骤为O(1)，而数组则需要将里面所有的元素都往后移一位，步骤为O(n)。
// 坏处在于查找很慢，在链表里如果要找到某个元素，必须从第一个开始，顺藤摸瓜式地往下查找。 所以，链表通常用在插入和删除比较多的场景，比如记账软件和代办事项等。

type LNode struct {
	Data int
	Next *LNode
}

func PrintLNode(err string, head *LNode) {
	fmt.Println(err)
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}

func Reverse(head *LNode) *LNode {
	// 递归返回条件，当前节点的下一个节点不存在时，停止递归，返回上层调用
	if head.Next == nil {
		return head
	} else {
		// 递归调用，知道链表到达末尾，然后返回新链表的表头
		newHead := Reverse(head.Next)
		// 当前节点的下一个节点的下一个节点为当前节点
		head.Next.Next = head
		head.Next = nil
		return newHead
	}
}

func main() {
	node1 := &LNode{}
	node1.Data = 1
	node2 := &LNode{}
	node2.Data = 2
	node3 := &LNode{}
	node3.Data = 3
	node4 := &LNode{}
	node4.Data = 4
	node5 := &LNode{}
	node5.Data = 5
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	PrintLNode("反转前:", node1)
	node := Reverse(node1)
	PrintLNode("反转后:", node)
}
