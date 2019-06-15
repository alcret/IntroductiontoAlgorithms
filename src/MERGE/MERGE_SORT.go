package MERGE

func Merge_Sort(array []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		Merge_Sort(array, p, q)
		Merge_Sort(array, q+1, r)
		Merge(array, p, q, r)
	}
}
