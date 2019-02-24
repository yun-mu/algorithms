package sort

func QuickSort(data []int) {
	l := len(data)
	if l < 2 {
		return
	}
	i := 1
	for i < l && data[i] < data[0] {
        i++
	}

	if i == l {
		data[0], data[i-1] = data[i-1], data[0]
		QuickSort(data[:l-1])
		return
	}

	for j := i+1; j < l; j++ {
		if data[j] < data[0] {
			data[j], data[i] = data[i], data[j]
			i++
		}
	}
	data[0], data[i-1] = data[i-1], data[0]
	QuickSort(data[:i-1])
	QuickSort(data[i:])
	return
}