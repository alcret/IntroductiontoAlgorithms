package FIND_MAX_CROSSING_SUBARRAY

func Find_Max_Crossing_SubArray(array []int, low, mid, high int) (max_left, max_right, left_right_sum int) {
	left_sum := -999
	sum := 0
	for i := mid; i >= low; i-- {
		sum += array[i]
		if sum > left_sum {
			left_sum = sum
			max_left = i
		}
	}

	right_sum := -999
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += array[i]
		if sum > right_sum {
			right_sum = sum
			max_right = i
		}
	}

	return max_left, max_right, left_sum + right_sum
}
