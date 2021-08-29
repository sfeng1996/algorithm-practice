package main

import "fmt"

// 斐波那契数列的特征是，第0和1位的数分别为0,1，从第二位开始，后面每一位的值都等于前一位加上前两位的值，该题求出任意给定位的斐波那契值
// 0,1,1,2,3,5,8,13,21,34

// 使用双指针迭代，每次只需求出f(n-1)和f(n-2)即可
// 时间复杂度为n，空间复杂度为1
func feibonaqiByIteration(num int) int {
	if num == 0 {
		return 0
	}
	// 当给定位为1时，直接返回1
	if num == 1 {
		return 1
	}
	// 定义n-1和n-2两个指针
	// low从0开始
	low := 0
	// high从1开始
	high := 1
	// i从2开始即可，前两个值是固定的
	for i:=2; i<=num; i++ {
		// 计算和
		sum := low + high
		// 将高的值赋给低值，
		low = high
		// 将和赋给高值
		high = sum
	}
	return high
}

func main() {
	fmt.Println(feibonaqiByIteration(8))
}
