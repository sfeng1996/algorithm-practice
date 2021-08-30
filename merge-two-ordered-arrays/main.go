package main

import (
	"fmt"
)

// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
// 说明:
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
// 使用两个指针，分别指向两个数组的值，比较这两个值，进行排序
func mergeTwoOrderedArrays(nums1 []int, m int, nums2 []int, n int) []int{
	// max 指向 nums1 与 nums2 合并之后的最后一个元素
	max := m + n - 1

	// 指向 num1 最后一个元素
	i := m - 1

	// 指向 num2 最后一个元素
	j := n -1


	for i >= 0 && j >= 0 {

		// 从右向左比较值的大小
		if nums1[i] > nums2[j] {
			nums1[max] = nums1[i]

			// i 向左移动
			i--
		} else {
			nums1[max] = nums2[j]

			// j 向左移动
			j--
		}

		// max 向左移动
		max --




	}

	// 如果 i 越界了，将 nums2 剩余的元素赋值到 num1 的 [0,m] 之间
	for j >= 0 {
		nums1[max] = nums2[j]
		max--
		j--
	}

	// 如果 j 越界了，其实下面这个循环可以省略。
	for i >= 0 {
		nums1[max] = nums1[i]
		max--
		i--
	}
	return nums1
}

func main() {
	list1 := []int{1,2,3,5,6,11,0,0,0,0,0,0,0,0}
	list2 := []int{4,5,5,7,8,10}
	fmt.Println(mergeTwoOrderedArrays(list1,6,list2,6))
}