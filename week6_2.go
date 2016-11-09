package main

import (
	"container/heap"
	"fmt"

	"./structures"
)

func main() {
	//MinHeap
	hMin := &structures.MinHeap{2, 1, 5}
	heap.Init(hMin)
	heap.Push(hMin, 3)
	fmt.Printf("minimum: %d\n", (*hMin)[0])
	for hMin.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(hMin))
	}

	hMax := &structures.MaxHeap{2, 1, 5}
	heap.Init(hMax)
	heap.Push(hMax, 3)
	fmt.Printf("maximum: %d\n", (*hMax)[0])
	for hMax.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(hMax))
	}
}
