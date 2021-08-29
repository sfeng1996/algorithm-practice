package main

import "fmt"

// 给定个数的硬币，将硬币排列，每行硬币的个数与行数相等，比如第一行只有一个硬币，第二行有两个硬币。求出满足行数与硬币数相等有多少行, 如果硬币数不够，返回前面的行数即可

// 遍历
func getArrangeCoins(num int) int {
	// i表示行数，也就是表示每行硬币的个数, i从1开始，因为没有第0行，i小于等于num
	for i:=1; i<=num; i++ {
		// 每次遍历减去这一行的硬币数
		num = num - i
		// 如果剩余的硬币数小于等于i，说明硬币不够或者刚好用完硬币，直接返回i
		if num <= i {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(getArrangeCoins(6))
}