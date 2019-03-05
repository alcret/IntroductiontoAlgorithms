package INSERTION_SORT

//插入排序
// 5,6,3,7,9,4,6
//时间复杂度n2
func Insertion_Sort(array []int) {
	for j := 1; j < len(array); j++ {
		key := array[j]
		i := j - 1
		for i >= 0 && array[i] > key {
			array[i+1] = array[i]
			i = i - 1
		}
		array[i+1] = key
	}
}
