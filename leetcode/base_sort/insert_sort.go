package base_sort

func InsertSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		temp := nums[i]
		for j := i - 1; j >= 0; j-- {
			if temp < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}
