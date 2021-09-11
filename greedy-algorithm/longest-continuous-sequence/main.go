package main

import "fmt"

// 给定一个乱序的数组，求出该数组中最长连续子序列的长度
// 输入[1,2,3,2,3,4,3,4,5,6,7]
// 输出 5

// 使用贪心算法，遍历数组，依次求出每个连续子序列的长度，最后算出最长长度
func LongestContinuousSequence(list []int) int {
	// 定义子序列遍历初始位置
	var start int
	// 定义最大长度
	var max int
	// i从1开始，不然list[i-1]会越界
	for i:=1; i<len(list); i++ {
		// 如果当前值小于前一个数，说明该子序列断了
		if list[i] < list[i-1] {
			start = i
		}
		max = maxNum(max, i-start+1)
	}
	return max
}

func maxNum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	list := []int{1,2,3,2,3,4,3,4,5,6,7}
	//LongestContinuousSequence(list)
	fmt.Println(LongestContinuousSequence(list))
}