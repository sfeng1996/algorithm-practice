package main

import "fmt"

// 找出数组的中心下标，数组的中心表示当前值的左边所有值的和等于当前值右边所有值的和，该题是求出当前值的下标

// 先求出数组所有值的和，循环遍历数组，并求和，在求和的过程中，并将初始的和减掉当前值，当两个和相等，这个值的下标就是中心下标

func pivotIndex(list []int) int {
	sum := sum(list)
	total := 0
	for i:=0; i<len(list); i++ {
		total += list[i] // 遍历求和
		if total == sum { // 为什么是先比较，而不是现将sum减去当前值再比较，因为如果两边相等的话，是带上当前值的，如果将sum当前值再比较，那永远也找不到中心下标
			return i
		}
		sum -= list[i]
	}
	return -1 // 表示没找到
}

func sum(list []int) int {
	sum := 0
	for i:=0; i<len(list); i++ {
		sum += list[i]
	}
	return sum
}

func main() {
	list := []int{1,2,3,5,6,8,3} // sum = 28
	fmt.Println(pivotIndex(list))
}