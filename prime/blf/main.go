package main

import "fmt"

// 给定一个数，求这个数以内的素数，素数(除0,1)表示出了1和本身能被整除，没有其它数能被整除的数

func blf(n int) int {
	// 定义一个变量，保存素数的个数
	var count int

	// 从2开始，因为0,1除外，不是素数
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count ++
		}
	}
	return count
}

func isPrime(x int) bool {
	// 从2开始，因为1本身就能被整除，只需要遍历到根号x即可，所以使用j*J<=x表示
	for j := 2; j * j <= x; j ++ {
		// x是否能够被j, 能够被整除说明不是素数
		if x % j == 0{
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%d 以内素数的个数: %d\n", 100, blf(100))
}