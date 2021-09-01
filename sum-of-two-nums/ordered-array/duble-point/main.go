package main

import "fmt"

// 给定一个有序数组，和一个目标值，找出数组中两个数的和等于这个目标值，并以数组形式返回这两个值
// 假设数组中只存在一对值

// 上一个方法使用二分查找，时间复杂度为nlog(n)，还是偏高，下面使用双指针方法实现
// 定义左指针，右指针, 如果连个指针指向的值之和等于target，直接返回
// 如果和大于target, 说明和应该在右指针左边，需要将右指针减一，进行下轮遍历
// 如果和小于target， 说明和应该在左指针右边，需要将左指针加一，进行下轮遍历
func sumOfTwoNumsByDublePoint(list []int, target int) []int {
	// 定义左指针和右指针
	low := 0
	high := len(list) - 1
	// low大于high结束循环
	for low <= high {
		if list[low] + list[high] == target {
			return []int{list[low], list[high]}
		} else if list[low] + list[high] > target {
			high--
		} else {
			low++
		}
	}
	return nil
}

func main() {
	list := []int{1,3,4,4,5,7}
	fmt.Println(sumOfTwoNumsByDublePoint(list, 10))
}
