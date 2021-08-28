package main

import "fmt"

// 给定一个无序数组，和一个目标值，找出数组中两个数的和等于这个目标值，并以数组形式返回这两个值
// 假设数组中只存在一对值

func sumOfTwoNums(list []int, target int) []int {
	// 初始化一个map，key为数组的值，value为数组下标
	numMap := make(map[int]int)
	for _, i := range list {
		// target减去如果遍历的值，能够在map找到，就可以直接返回
		if j, ok := numMap[target - i]; ok {
			return []int{i, j}
		}
		// 如果找不到，就将这个值及下标保存在map里
		numMap[list[i]] = i
	}
	return nil
}

func main() {
	list := []int{1,3,2,4,5,7}
	fmt.Println(sumOfTwoNums(list, 10))
}