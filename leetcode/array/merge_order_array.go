package array

/*
合并两个有序数组
*/

func MergeOrderArray(array1 []int, array2 []int) []int {
	len1 := len(array1)
	len2 := len(array2)
	first, second := 0, 0
	var array3 []int
	for first < len1 && second < len2 {
		if array1[first] > array2[second] {
			array3 = append(array3, array2[second])
			second += 1
		} else {
			array3 = append(array3, array1[first])
			first += 1
		}
	}
	if first < len1 {
		array3 = append(array3, array1[first:]...)
	}
	if second < len2 {
		array3 = append(array3, array2[second:]...)
	}
	return array3
}
