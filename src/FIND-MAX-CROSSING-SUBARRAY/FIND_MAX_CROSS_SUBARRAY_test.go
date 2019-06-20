package FIND_MAX_CROSSING_SUBARRAY

import (
	"fmt"
	"testing"
)

func TestFind_Max_Crossing_SubArray(t *testing.T) {
	array := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Println(Find_Maximum_SubArray(array, 0, len(array)-1))

}
