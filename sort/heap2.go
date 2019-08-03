package sort

// Max Heap
func HeapSort2(data []int) {
	l := len(data)
	mid := l/2
	for i := mid; i >= 0; i-- {
        heap(data, i, l-1)
	}

	for i := l-1; i >= 0; i-- {
		data[i], data[0] = data[0], data[i]
		heap(data, 0, i-1)
	}
}

func heap(data []int, i, end int) {
	l, r := 2*i+1, 2*i+2
	if l > end {
		return
	}

	n := l
	if r <= end && data[r] > data[l] {
        n = r
	}

	if data[i] > data[n] {
        return
	}
	data[i], data[n] = data[n], data[i]
	heap(data, n, end)
}
