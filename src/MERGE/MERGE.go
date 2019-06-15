package MERGE

func Merge(array []int, p, q, r int) {
	// 左右两个组，每个都要有自己的顺序
	n1 := q - p + 1
	n2 := r - q

	arrayL := make([]int, n1+1)
	copy(arrayL, array[:n1])
	arrayR := make([]int, n2+1)
	copy(arrayR, array[n1:])

	arrayL[n1] = 9999
	arrayR[n2] = 9999

	i := 0
	j := 0

	for k := p; k <= r; k++ {
		if arrayL[i] <= arrayR[j] {
			array[k] = arrayL[i]
			i = i + 1
		} else {
			array[k] = arrayR[j]
			j = j + 1
		}
	}
}
