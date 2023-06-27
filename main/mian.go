package main

import "fmt"
import "go_learn/leetcode/base_sort"

func main() {

	// 冒泡排序
	x := base_sort.BubbleSort([]int{2, 3, 41, 33, 1})
	fmt.Println(x)
	//直接插入排序
	fmt.Println(base_sort.InsertSort([]int{2, 3, 41, 33, 1}))
	//简单选择排序
	fmt.Println(base_sort.SimpleSelectSort([]int{2, 3, 41, 33, 1}))
}
