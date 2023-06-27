package main

import "fmt"
import "go_learn/leetcode/base_sort"
import "go_learn/go_knowledge"

func CopySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func main() {
	nums := []int{23, 11, 7, 29, 33, 59, 8, 20, 9, 3, 2, 6, 10, 44, 83, 28, 5, 1, 0, 36}
	// 冒泡排序

	x := base_sort.BubbleSort(CopySlice(nums))
	fmt.Println(x)
	//直接插入排序
	fmt.Println(base_sort.InsertSort(CopySlice(nums)))
	//简单选择排序
	fmt.Println(base_sort.SimpleSelectSort(CopySlice(nums)))
	// 二分排序算法
	fmt.Println(base_sort.TwoSplitSort(CopySlice(nums)))

	// go_struct
	go_knowledge.GoStructInit()
}
