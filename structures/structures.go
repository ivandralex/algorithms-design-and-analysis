package structures

// A MaxHeap implements heap.Interface and holds Items.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	i := old[n-1]
	*h = old[0 : n-1]
	return i
}
