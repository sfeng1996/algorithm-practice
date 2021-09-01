package main

import (
	"fmt"
)

// 给一个整数数组，找出平均数最大且长度为k的下标连续的子数组，并输出该最大平均数
// 输入: [1,12,-5,-6,50,3], k=4
// 输出: 12.75
// 最大平均数(12-5-6+50)/4=12.75

// 使用滑动窗口，先计算第一个窗口的值，然后根据第一个窗口的值计算出后面每个窗口的值，再选出最大值
func maxAverage(list []int, k int) float64 {
	sum := 0
	// 计算第一个窗口的值
	for i:=0; i<k; i++ {
		sum += list[i]
	}
	// 初始化最大值，当前最大值为第一个窗口和
	maxNum := sum
	// 遍历求出后面窗口的值，后面的窗口和等于前一个窗口和减去窗口的第一个数，再加上当前遍历的数
	for j:=k; j<len(list); j++ {
		sum = sum - list[j-k] + list[j]
		maxNum = max(maxNum, sum)
	}
	return float64(maxNum) / float64(k)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
func main() {
	list := []int{1,12,-5,-6,50,3}
	k := 4
	fmt.Println(maxAverage(list, k))
}
