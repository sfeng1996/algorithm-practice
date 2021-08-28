package main

import "fmt"

// 给定一个有序数组，将其重复项删除，并求出删除后数组长度

// 循环遍历数组值，定义两个指针，i表示为慢指针，j表示为快指针，遍历第一个元素，此时i指向1，j指向2，i!=j，将i加1，再将j值赋予i，这个操作不影响{1,2,3}
// 如果遍历到4，i指向4，j也指向4，那么i=j，直接进入下一个循环，j+1，后面的操作就和上面一致
func removeDuplicates(list []int) int {
	i := 0 // 将i表示慢指针
	for j:=0; j<len(list); j++ { // j表示为快指针
		if list[j] != list[i] {
			i ++
			list[i] = list[j]
		}
	}
	return i + 1 // 因为i从0开始，计算数组长度，需要加1
}


func main() {
	list := []int{1,2,8,4,4,5,5,6,7,7,3}
	fmt.Println(removeDuplicates(list))
}