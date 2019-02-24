package sort

func ShellSort(data []int) {
	l := len(data)
	if l < 2 {
		return
	}
	
	for gap := l>>1; gap > 0; gap >>=1 {
        for i := gap; i < l; i++ {
			for j := i; j >= gap && data[j] < data[j-gap]; j -= gap {
                data[j], data[j-gap] = data[j-gap], data[j]
			}
		}
	}
}