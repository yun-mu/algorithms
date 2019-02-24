package sort

// Every time insert forward, front all is already sorted.
func InsertSort(data []int) {
	l := len(data)
    for i := 1; i < l; i++ {
        for j := i-1; j >= 0 && data[j] > data[j+1]; j-- {
			data[j], data[j+1] = data[j+1], data[j]
		}
	}
}