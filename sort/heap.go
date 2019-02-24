package sort

import "math"

type MinHeap struct {
    data []int
}

func NewMinHeap() *MinHeap {
    // 第一个元素仅用于结束insert中的 for 循环
    h := &MinHeap{data: []int{math.MinInt32}}
    return h
}

func (this *MinHeap) Build(data []int) {
    for _, v := range data {
		this.Insert(v)
	}
}

func (this *MinHeap) Insert(v int) {
    this.data = append(this.data, v)
    i := len(this.data) - 1
    // 上浮
    for ; this.data[i/2] > v; i /= 2 {
        this.data[i] = this.data[i/2]
    }

    this.data[i] = v
}

func (this *MinHeap) PopMin() (int, error) {
    if len(this.data) <= 1 {
        return 0, fmt.Errorf("MinHeap is empty")
	}
	minData := this.data[1]
	lastData := this.data[len(this.data)-1]
	var i, child int
	for i = 1; i*2 < len(this.data); i = child {
		child = i * 2
		if child < len(this.data)-1 && this.data[child+1] < this.data[child] {
			child ++
		}
		// 下滤一层
		if lastData > this.data[child] {
			this.data[i] = this.data[child]
		} else {
			break
		}
	}
	this.data[i] = lastData
	this.data = this.data[:len(this.data)-1]
	return minData, nil
}

func (this *MinHeap) Length() int {
	return len(this.data) - 1
}

func (this *MinHeap) Min() (int, error) {
	if len(this.data) > 1 {
		return this.data[1], nil
	}
	return 0, fmt.Errorf("heap is empty")
}

// Min Heap
func HeapSort(data []int) []int {
	h := NewMinHeap()
	h.Build(data)
	res := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		res[i], _ = h.PopMin()
	}
	return res
}