package main

import "fmt"

// 选择排序是先将每一趟最前面的数最为最小(或最大)的数，然后一次将该值与后面的值比较，一轮下来选出最小的值，没一轮都会选出最小的值
// 时间复杂度为O(n^2)
// 空间复杂度为O(1)
// 不稳定(相等元素的相对位置在交换顺序后相对位置不变，说明该算法稳定)

func selectSort(list []int) []int {
	for i := range list {
		for j:=i+1; j<len(list); j++ {
			// 选择排序将最前面的值与每一个值比较
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

func main() {
	list := []int{1,3,6,2,5,9,8,10,20,15}
	fmt.Println(selectSort(list))
}
