package main

import "fmt"

// 给定一个整数数组，求出数组中能构成三角形的最大周长

// 先排序, 然后从最大值开始遍历，得出第一个周长，即为最大周长
func MaxPerimeterOfTriangle(list []int) int {
	list = sort(list)
	for i:=len(list)-1; i>=2; i-- {
		if list[i-1] + list[i-2] > list[i] {
			return list[i-1] + list[i-2] + list[i]
		}
	}
	return 0
}

// 使用冒泡排序
func sort(list []int) []int {
	for i := range list{
		for j:=1; j<len(list)-i; j++ {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	return list
}

func main() {
	list := []int{2,3,3,6,7,2,1}
	fmt.Println(MaxPerimeterOfTriangle(list))
}