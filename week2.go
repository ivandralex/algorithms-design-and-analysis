package main

import (
	"fmt"

	"./common"
)

var testCase1 = []int{4, 9, 11, 8, 10, 7, 15, 12, 5, 13, 6, 3, 2, 14, 1}

//Result of procedure
type Result struct {
	numbers []int
	counter int
}

func main() {
	var res = new(Result)

	if false {
		res.numbers = testCase1
	} else {
		launch("./data/week2_10.txt")
		launch("./data/week2_100.txt")
		launch("./data/week2_1000.txt")
	}
}

func launch(path string) {
	var res = new(Result)

	if false {
		res.numbers = testCase1
	} else {
		res.numbers = common.ReadIntegers(path)
	}

	sortAndCount(res, 0, len(res.numbers)-1, 0)

	//fmt.Printf("Sorted array: %v\n", res.numbers)
	fmt.Printf("%s: %d\n", path, res.counter)
}

func sortAndCount(res *Result, start int, end int, level int) {
	if end-start < 1 {
		return
	}

	res.counter += end - start

	var p = getPivotIndex(start, end, 1)

	//If pivot is not the first element let's swap pivot and first element
	if p != start {
		swap(start, p, res.numbers)
		p = start
	}

	//fmt.Printf("#%d\tBefore: %v %d\n", level, res.numbers[start:end+1], res.numbers[p])

	for j := start + 1; j <= end; j++ {
		//fmt.Printf("#%d\tEl: %d < %d < %d: %d\n", level, start, j, end, res.numbers[j])
		//If there is an inversion
		if res.numbers[p] > res.numbers[j] {
			swap(j, p+1, res.numbers)
			swap(p, p+1, res.numbers)
			p++
		}
	}

	/* Debug logs
	fmt.Printf("#%d\tAfter: %v %d\n", level, res.numbers[start:end+1], res.numbers[p])
	fmt.Printf("#%d\tTemp: %v\n", level, res.numbers)
	if p > start {
		fmt.Printf("#%d\tFork right: %d: %v\n", level, res.numbers[p], res.numbers[start:p])
	}
	if p+1 < end {
		fmt.Printf("#%d\tFork left: %d: %v\n", level, res.numbers[p], res.numbers[p+1:end+1])
	}
	*/

	if p > start {
		sortAndCount(res, start, p-1, level+1)
	}
	if p+1 < end {
		sortAndCount(res, p+1, end, level+1)
	}
}

func getPivotIndex(start int, end int, task int) int {
	return start
}

func swap(i int, j int, numbers []int) {
	var temp = numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp
}
