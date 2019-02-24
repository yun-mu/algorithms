package sort

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data)/2
	pre := MergeSort(data[:mid])
	next := MergeSort(data[mid:])
	return merge(pre, next)
}

func merge(first, second []int) []int {
	data := make([]int, len(first)+len(second))
	i, fI, sI := 0, 0, 0
	for fI < len(first) && sI < len(second) {
		if first[fI] < second[sI] {
			data[i] = first[fI]
			fI++
		} else {
			data[i] = second[sI]
			sI++			
		}
		i++
	}

	for fI < len(first) {
		data[i] = first[fI]
		fI++
		i++
	}

	for sI < len(second) {
		data[i] = second[sI]
		sI++
		i++
	}
	return data
}