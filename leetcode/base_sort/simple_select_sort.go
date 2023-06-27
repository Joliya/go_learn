package base_sort

func SimpleSelectSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 0; j < length-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
	}
	return nums
}
