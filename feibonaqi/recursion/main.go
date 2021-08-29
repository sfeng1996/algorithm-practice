package main

import "fmt"

// 斐波那契数列的特征是，第0和1位的数分别为0,1，从第二位开始，后面每一位的值都等于前一位加上前两位的值，该题求出任意给定位的斐波那契值
// 0,1,1,2,3,5,8,13,21,34

// 除了第0位和第一位是固定值之外，后面每一位的值都可以表示为f(n-1)+f(n-2)，所以可以使用递归求解，递归的条件值为n=0,n=1
func feibonaqiByRecursion(num int) int {
	// 当给定位为0时，直接返回0
	if num == 0 {
		return 0
	}
	// 当给定位为1时，直接返回1
	if num == 1 {
		return 1
	}
	// 递归
	return feibonaqiByRecursion(num-1) + feibonaqiByRecursion(num-2)
}

func main() {
	fmt.Println(feibonaqiByRecursion(8))
}