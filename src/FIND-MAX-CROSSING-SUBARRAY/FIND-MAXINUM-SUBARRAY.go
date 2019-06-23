package FIND_MAX_CROSSING_SUBARRAY

func Find_Maximum_SubArray(array []int, low, high int) (rlow, rhigh, sum int) {
	if high == low {
		return low, high, array[low]
	} else {
		mid := (low + high) / 2
		left_low, left_high, left_sum := Find_Maximum_SubArray(array, low, mid)
		right_low, right_high, right_sum := Find_Maximum_SubArray(array, mid+1, high)
		cross_low, cross_high, cross_sum := Find_Max_Crossing_SubArray(array, low, mid, high)

		if left_sum >= right_sum && left_sum >= cross_sum {

			return left_low, left_high, left_sum
		} else if right_sum >= left_sum && right_sum >= cross_sum {

			return right_low, right_high, right_sum
		} else {

			return cross_low, cross_high, cross_sum
		}
	}
}
