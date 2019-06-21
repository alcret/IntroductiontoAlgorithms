package STRASSEN

import (
	"fmt"
	"testing"
)

func TestStrassen(t *testing.T) {
	A := [][]int{[]int{1, 1}, []int{1, 1}}
	B := [][]int{[]int{1, 1}, []int{1, 1}}

	C := SQUARE_MATRIX_MULTIPLAY(A, B)
	fmt.Println(C)

}
