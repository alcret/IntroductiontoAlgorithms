package STRASSEN

func SQUARE_MATRIX_MULTIPLAY(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i, _ := range C {
		C[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = 0
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}
