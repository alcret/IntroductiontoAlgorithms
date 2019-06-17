package MERGE

func Merge(array []int, p, q, r int) {
	// 左右两个组，每个都要有自己的顺序
	n1 := q - p + 1
	n2 := r - q

	arrayL := make([]int, n1+1)
	copy(arrayL, array[p:n1+p])
	arrayL[n1] = 9999

	arrayR := make([]int, n2+1)
	copy(arrayR, array[q+1:q+n2+1])
	arrayR[n2] = 9999

	i, j := 0, 0
	for k := p; k <= r; k++ {
		if arrayL[i] <= arrayR[j] {
			array[k] = arrayL[i]
			i++
		} else {
			array[k] = arrayR[j]
			j++
		}
	}
}
