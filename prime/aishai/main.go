package main

import "fmt"

// 给定一个数，求这个数以内的素数，素数(除0,1)表示出了1和本身能被整除，没有其它数能被整除的数

// 埃塞法将合数标记，使得合数不会被遍历
func eratosthenes(n int) int {
	// 首先所有数都是素数，不过该数组默认值是false，false表示素数
	isPrimeList := make([]bool, 100)
	var count int
	for i:=2; i<n; i++ {
		// 是素数count加一
		 if !isPrimeList[i] {
		 	count ++
		 	// 从i的i倍开始，因为1到i被标记过了，每次递增i，都是合数，将切片的值标记为合数
			for j:=i*i; j<n; j+=i {
				isPrimeList[j] = true
			}
		 }
	}
	return count
}

func main() {
	fmt.Printf("%d 以内素数的个数: %d\n", 100, eratosthenes(100))
}
