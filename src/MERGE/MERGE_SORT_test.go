package MERGE

import (
	"fmt"
	"testing"
)

func TestMERGE(t *testing.T) {
	var ii []int = []int{2, 4, 5, 7, 1, 2, 3}

	Merge(ii, 0, 3, len(ii)-1)

	//Merge_Sort(ii,0,len(ii))
	fmt.Println(ii)
}
