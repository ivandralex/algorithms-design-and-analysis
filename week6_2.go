package main

import (
	"container/heap"
	"fmt"

	"./common"
	"./structures"
)

func main() {
	var numbers []int
	if false {
		numbers = common.ReadIntegers("data/median.txt")
	} else {
		numbers = []int{1, 12, 3, 14, 5, 16, 7, 18, 9, 10}
	}

	//MinHeap
	hMin := &structures.MinHeap{}
	heap.Init(hMin)

	hMax := &structures.MaxHeap{}
	heap.Init(hMax)

	var median int

	total := 0

	for _, n := range numbers {
		median = maintainMedian(n, hMin, hMax)
		total += median
		fmt.Printf("%d: %v %v => %d\n", n, (*hMax), (*hMin), median)
	}

	fmt.Printf("Modulo: %d\n", total%10000)
}

/*
[ hMax hMin]

*/
func maintainMedian(n int, hMin *structures.MinHeap, hMax *structures.MaxHeap) int {
	min := 0
	max := 0
	if hMin.Len() > 0 {
		min = (*hMin)[0]
	}
	if hMax.Len() > 0 {
		max = (*hMax)[0]
	}

	//If it lies in higher end
	if n > min {
		heap.Push(hMin, n)
	}

	//If it lies in lower end
	if n < max {
		heap.Push(hMax, n)
	}

	//Re-balance heaps to achieve 50/50 split:

	//if median in higher end
	if hMin.Len()-hMax.Len() > 1 {
		excess := heap.Pop(hMin)
		heap.Push(hMax, excess)
	}
	if hMax.Len()-hMin.Len() > 1 {
		excess := heap.Pop(hMax)
		heap.Push(hMin, excess)
	}

	var median int

	if hMax.Len() >= hMin.Len() {
		median = (*hMax)[0]
	} else {
		median = (*hMin)[0]
	}

	return median
}
