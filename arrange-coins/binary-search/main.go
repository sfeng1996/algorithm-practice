package main

import "fmt"

// 给定个数的硬币，将硬币排列，每行硬币的个数与行数相等，比如第一行只有一个硬币，第二行有两个硬币。求出满足行数与硬币数相等有多少行, 如果硬币数不够，返回前面的行数即可

// 使用二分查找，一般从一个有序数组找出一个数，都可以使用二分查找降低时间复杂度，那么这题如何转换成从有序数组里找数呢
// 假如最后一行行数是硬币数，行数从1到n肯定是有序的，那么我们要找的是1到n之间根据某个数求得的硬币总数等于n，说明这一行就是要找的数
func getArrangeCoinsByBinarySearch(num int) int {
	// 定义low, high
	// low表示一行没有
	low := 0
	// high表示第num行
	high := num
	for low <= high {
		// 求出中间值
		mid := low + (high - low)/2
		// 求出硬币总数,根据num=(x^2+x)/2，这是等差数列求和公式
		costs := ((mid+1)*mid)/2
		// 如果cost等于num，说明mid就是这一行
		if costs == num {
			return mid
		}
		// 如果cost小于num，说明值在mid右边
		if costs < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(getArrangeCoinsByBinarySearch(10))
}
