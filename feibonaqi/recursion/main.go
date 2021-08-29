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

// 上面递归效率很差，因为会重复计算，例如f(8)=f(7)+f(6)，需要计算f(7)和f(6)的值，f(7)又需要计算f(6)和f(5)的值，这样都会计算f(6)的值，所以可以使用数组将重复计算的值保留起来
func feibonaqiByRecursionYouHua(num int) int {
	// 初始化一个切片，切片长度为斐波那契数列长度，因为斐波那契数列从0开始，所以加1
	duplicateList := make([]int, num+1)
	// 当给定位为0时，直接返回0
	if num == 0 {
		return 0
	}
	// 当给定位为1时，直接返回1
	if num == 1 {
		return 1
	}
	// 保存重复值，切片初始化未赋值，默认值为0，所以如果值不等于0的话，说明这位已经在数组里有保存了，直接返回
	if duplicateList[num] != 0 {
		return duplicateList[num]
	}
	// 递归
	return feibonaqiByRecursionYouHua(num-1) + feibonaqiByRecursionYouHua(num-2)
}

func main() {
	fmt.Println(feibonaqiByRecursion(8))
	// 优化后
	fmt.Println(feibonaqiByRecursionYouHua(8))
}