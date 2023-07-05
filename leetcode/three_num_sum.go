package leetcode

import (
	"sort"
)

func ThreeNumSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	result := [][]int{}
	// 排序
	sort.Ints(nums)
	//nums = base_sort.BubbleSort(nums)
	for i, num := range nums {
		// 遇到大于0的数据就可以直接返回了， 因为后面的数肯定都是大于0 的，和不可能为0
		if num > 0 {
			return result
		}
		// 如果所有的数字都相同，则直接返回了
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 左指针从i 的下一个开始计数， 防止重复统计
		left := i + 1
		right := len(nums) - 1
		for left < right {
			value := num + nums[left] + nums[right]
			if value == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
				// 这儿两个 while 循环是为了 解决连续相同值，直接移动指针
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
				continue
			}
			// 大于0表示  正数大了， 右指针左移
			if value > 0 {
				right--
			} else {
				// 否则的话  左指针右移
				left++
			}
		}
	}
	return result
}
