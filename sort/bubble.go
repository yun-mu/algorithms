package sort

// Each time you select the biggest of all.
func BubbleSort(data []int) {
	l := len(data)
    for i := 0; i < l-1; i++ {
        for j := 0; j < l-1-i; j++ {
			if data[j] > data[j+1] {
			    data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}