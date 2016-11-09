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
		numbers = common.ReadIntegers("./data/median.txt")
	} else {
		numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	}

	//MinHeap
	hMin := &structures.MinHeap{}
	heap.Init(hMin)
	/*
		heap.Push(hMin, 3)
		fmt.Printf("minimum: %d\n", (*hMin)[0])
		for hMin.Len() > 0 {
			fmt.Printf("%d\n", heap.Pop(hMin))
		}
	*/

	hMax := &structures.MaxHeap{}
	heap.Init(hMax)
	/*
		heap.Push(hMax, 3)
		fmt.Printf("maximum: %d\n", (*hMax)[0])
		for hMax.Len() > 0 {
			fmt.Printf("%d\n", heap.Pop(hMax))
		}
	*/

	for _, n := range numbers {
		median := maintainMedian(n, hMin, hMax)
		fmt.Printf("%v %v => %d\n", (*hMax), (*hMin), median)
	}
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

	if hMin.Len() >= hMax.Len() {
		median = (*hMin)[0]
	} else {
		median = (*hMax)[0]
	}

	return median
}
