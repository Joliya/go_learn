package base_sort

func TwoSplitSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		left, right := 0, i-1
		temp := nums[i]
		for left <= right {
			mid := (left + right) / 2
			if temp < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		for j := i - 1; j >= left; j-- {
			nums[j+1] = nums[j]
		}
		nums[left] = temp
	}
	return nums
}
