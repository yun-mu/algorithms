package sort

// Each time you select the smallest.
func SelectSort(data []int) {
	l := len(data)
    for i := 0; i < l-1; i++ {
		minI := i
        for j := i+1; j < l ; j++ {
			if data[j] < data[minI] {
				minI = j
			}
		}
		if minI != i {
            data[i], data[minI] = data[minI], data[i]
		}
	}
}