package main

import (
	"fmt"
	"math"
)

// 上一个方法实现需要先进行排序，时间复杂度主要在排序算法，如果排序算法时间复杂度高的话，那么整个程序性能不太好，所以下面使用线性扫描的方式实现
// 根据上一个方法的分析, 其实只要在数组中找出五个值即可，最大值，第二大值，第三大值，最小值，第二小值

func getMaxMin(list []int) int {
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MaxInt64, math.MinInt64
	for _, i := range list {
		if i < min1 {
			min2 = min1
			min1 = i
		} else if i < min2 {
			min2 = i
		}
		if i > max1 {
			max3 = max2
			max2 = max1
			max1 = i
		} else if i > max2 {
			max3 = max2
			max2 = i
		} else if i > max3 {
			max3 = i
		}
	}
	return max(min1*min2*max1, max1*max2*max3)
}

func max(a,b int) int {
	if a >b {
		return a
	}
	return b
}

func main() {
	list1 := []int{1,22,33,4,5,6}
	list2 := []int{-2,-3,-66,-5}
	list3 := []int{-56,0,34,5,6}
	fmt.Println(getMaxMin(list1))
	fmt.Println(getMaxMin(list2))
	fmt.Println(getMaxMin(list3))

}