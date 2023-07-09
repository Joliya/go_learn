package main

import (
	"fmt"
	"go_learn/go_knowledge"
	"go_learn/leetcode"
	"go_learn/leetcode/array"
	"go_learn/leetcode/base_sort"
	"go_learn/leetcode/list"
)

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

	// 合并两个有序数组
	a := []int{2, 3, 66, 88, 102}
	b := []int{4, 55, 65, 99}
	fmt.Println(array.MergeOrderArray(a, b))

	// 三数之和
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(leetcode.ThreeNumSum(nums))
	fmt.Println(leetcode.ThreeNumSum([]int{0, 0, 0, 0}))

	// 验证反转列表
	head := &list.Node{Data: 1}
	node2 := &list.Node{Data: 2}
	node3 := &list.Node{Data: 3}
	node4 := &list.Node{Data: 4}
	node5 := &list.Node{Data: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	fmt.Println("原始链表:")
	list.DisplayLinkedList(head)

	fmt.Println("反转后的链表:")
	newHead := list.ReverseList(head)
	list.DisplayLinkedList(newHead)
}
