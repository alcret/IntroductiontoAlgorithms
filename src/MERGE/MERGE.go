package MERGE

func Merge(array []int, p, q, r int) {
	n1 := q - p + 1
	//n2 := r - q

	//arrayL := make([]int, n1+1)
	//arrayR := make([]int, n2+1)
	var arrayL = make([]int, 0)
	var arrayR = make([]int, 0)

	arrayL = append(arrayL, array[:n1]...)
	arrayR = append(arrayR, array[n1:]...)
	arrayL = append(arrayL, 9999)
	arrayR = append(arrayR, 9999)
	i := 0
	j := 0
	//fmt.Println(arrayL,arrayR,arrayL[i])

	for k := p; k < r; k++ {
		if arrayL[i] <= arrayR[j] {
			array[k] = arrayL[i]
			i = i + 1
		} else {
			array[k] = arrayR[j]
			j = j + 1
		}
	}

}
