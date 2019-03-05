package INSERTION_SORT

import (
	"fmt"
	"testing"
)

func TestInsertion_Sort(t *testing.T) {
	array := []int{6, 5, 3, 7, 9, 4, 6}
	Insertion_Sort(array)
	fmt.Println(array)
}
