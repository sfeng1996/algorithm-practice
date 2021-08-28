package main

import "fmt"

// 求正整数的平方根，查找离它最近的整数即可

// 使用二分查找，二分查找时间复杂度为log(n)，主要思路是将求x的中间值，使得当前值逼近x的中间值
func binarySearch(x int) int {
	// 定义两个指针，左边值，右边值
	index := -1
	l := 0
	r := x
	// 循环查找，l始终小于等于r
	for l <= r {
		// 求中间值
		mid := l + (r - l)/2
		if (mid * mid) <= x {
			index = mid
			// 进入下一次循环
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return index
}

func main() {
	fmt.Println(binarySearch(26))
}
