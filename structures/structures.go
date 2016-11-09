package structures

// A MaxHeap implements heap.Interface and holds Items.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Pop pop
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	i := old[n-1]
	*h = old[0 : n-1]
	return i
}

//Push push
func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
