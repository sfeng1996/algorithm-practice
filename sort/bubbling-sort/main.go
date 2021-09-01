package main

import "fmt"

// 冒泡排序每次遍历比较相邻的值，进行排序
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// 稳定性: 稳定

func bubblingSort(list []int) []int {
	for i := range list {
		for j:=i+1; j<len(list); j++ {
			// 冒泡排序是比较相邻值，不能使用list[i]与list[j]比较
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
	return list
}

func main() {
	list := []int{1,3,6,2,5,9,8,10,20,15}
	fmt.Println(bubblingSort(list))
}
