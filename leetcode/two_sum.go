package leetcode

func twoSum(nums []int, target int) []int {
	indexMap := map[int]int{}
	for i, num := range nums {
		value := target - num
		if j, ok := indexMap[value]; ok {
			return []int{j, i}
		}
		indexMap[num] = i
	}
	return []int{}
}
