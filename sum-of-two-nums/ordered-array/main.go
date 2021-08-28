package main

import "fmt"

// 给定一个有序数组，和一个目标值，找出数组中两个数的和等于这个目标值，并以数组形式返回这两个值
// 假设数组中只存在一对值

// 该题的目的，实际上是寻找数组中的两个数，先遍历数组中的值，假定这个值就是其中一个值，寻找下一个值可以使用二分查找，降低时间复杂度，对于有序数组，二分查找是最节省时间的算法
func sumOfTwoNumsByBinarySearch(list []int, target int) []int {
	// 假定i就是其中一个值
	for i := range list {  // 遍历数组，会返回两个值，前面值是下标，后面是值，如果只写一个的话，默认是下标
		low := i // 最左边值下标
		high := len(list) - 1  // 最右边值下标
		for low <= high {
			mid := low + (high - low) / 2 // 计算中间点
			// 如果中间值等于目标值减去当前遍历值，直接返回list[i]，list[mid]
			if list[mid] == target - list[i] {
				return []int{list[i], list[mid]}
				// 如果中间值大于目标值减去当前值, 说明另一个值在中间值的左边，high减一，进行下轮遍历
			} else if list[mid] > target - list[i] {
				high = mid - 1
			} else {
				// 如果中间值小于目标值减去当前值, 说明另一个值在中间值的右边，low加一，进行下轮遍历
				low = mid + 1
			}
		}

	}
	return nil
}

func main() {
	list := []int{1,3,4,4,5,7}
	fmt.Println(sumOfTwoNumsByBinarySearch(list, 10))
}
