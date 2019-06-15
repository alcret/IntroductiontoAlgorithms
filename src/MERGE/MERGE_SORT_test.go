package MERGE

import (
	"fmt"
	"testing"
)

func TestMERGE(t *testing.T) {
	//var ii []int = []int{2,4,5,7,1,2,3,6}
	ii := []int{5, 2, 4, 7, 1, 3, 2, 6}
	//fmt.Println(ii)
	Merge(ii, 0, 3, len(ii))
	//Merge_Sort(ii,0,len(ii))
	fmt.Println(ii)
}
