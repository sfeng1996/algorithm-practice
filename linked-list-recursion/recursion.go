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
// 使用递归反转, 递归就是先执行最后返回的过程，递归是先压栈，栈是先进后出的数据结构，最后进来的函数最先出栈被执行

func Reverse(head *LNode) *LNode{
	// 传入空链表或者到最后一个接点了，直接返回
	if head == nil || head.Next == nil {
		return head
	} else {
		// 寻找到最后一个接点即返回，此时newHead为最后一个节点，head为最后一个节点的前一个节点
		newHead := Reverse(head.Next)
		// 逆转
		head.Next.Next = head
		// 将head.next设置为nil，防止环形
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