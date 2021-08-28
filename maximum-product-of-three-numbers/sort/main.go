package main

import (
	"fmt"
	"math"
)

// 求数组中三个数的最大乘积，这里分以下几种情况
// 当数组中全是正数的话，那么只需要找出数组中最大的三个数，乘积就是最大
// 当数组中全是负数的话，也只需找出数组中最大的三个数，乘积也是最大
// 当数组中既有正数也有负数，又分以下情况
// 当负数个数为单数时，还是获取数组中最大的三个数
// 当负数个数为复数时，获取绝对值最大的三个数
// 总结以上几种情况，只有两种情况, 对于已经按照顺序排序好的数组，只要获取最大的三个值，数组的前两个值，和最后一个值，再进行比较这个两个值的大小
// max(list[len-1]*list[len-2]*list[len-3], list[0]*list[1]*list[len-1])

// 先将数组排序
func getMaxProduct(list []int) int {
	// 先将数组排序
	list =  sort(list)
	// 获取数组长度
	len := len(list)
	// 当全为正数或是负数的话，找出最大的三个值 list[len-1]*list[len-2]*list[len-3]
	// 当既有正数既有负数的话，找出绝对值最大的三个值
	result := math.Max(float64(list[len-1]*list[len-2]*list[len-3]), float64(list[0]*list[1]*list[len-1]))
	return int(result)
}

func sort(list []int) []int {
	// 冒泡排序
	for i := range list {
		for j:=1; j<len(list)-i; j++ {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	return list
}

func main() {
	list1 := []int{1,22,33,4,5,6}
	list2 := []int{-2,-3,-66,-5}
	list3 := []int{-56,0,34,5,6}
	fmt.Println(getMaxProduct(list1))
	fmt.Println(getMaxProduct(list2))
	fmt.Println(getMaxProduct(list3))
	//fmt.Println(sort(list1))

}