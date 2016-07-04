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
		//launch("./data/week2_10.txt", 3)
		//launch("./data/week2_100.txt", 3)
		//launch("./data/week2_1000.txt")
		launch("./data/week2.txt", 1)
	}
}

func launch(path string, task int) {
	var res = new(Result)

	if false {
		res.numbers = testCase1
	} else {
		res.numbers = common.ReadIntegers(path)
	}

	sortAndCount(res, 0, len(res.numbers)-1, 0, task)

	//fmt.Printf("Sorted array: %v\n", res.numbers)
	fmt.Printf("%s: %d\n", path, res.counter)
}

func sortAndCount(res *Result, start int, end int, level int, task int) {
	if end-start < 1 {
		return
	}

	res.counter += end - start

	var p = getPivotIndex(res, start, end, task)

	//fmt.Printf("#%dPivot: %d %d %d\n", level, start, end, p)

	//If pivot is not the first element let's swap pivot and first element
	if p != start {
		swap(start, p, res.numbers)
		p = start
	}

	//fmt.Printf("#%d\tBefore: %v %d\n", level, res.numbers[start:end+1], res.numbers[p])

	for j := start + 1; j <= end; j++ {
		//If there is an inversion swap elements
		if res.numbers[start] > res.numbers[j] {
			swap(j, p+1, res.numbers)
			p++
		}
	}

	swap(start, p, res.numbers)

	/*
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
		sortAndCount(res, start, p-1, level+1, task)
	}
	if p+1 < end {
		sortAndCount(res, p+1, end, level+1, task)
	}
}

func getPivotIndex(res *Result, start int, end int, task int) int {
	if task == 3 {
		var middle = getMiddle(start, end)

		var temp = new(Result)
		temp.numbers = []int{res.numbers[start], res.numbers[end], res.numbers[middle]}

		sortAndCount(temp, 0, 2, 0, 1)

		var median = temp.numbers[1]

		switch median {
		case res.numbers[start]:
			return start
		case res.numbers[end]:
			return end
		case res.numbers[middle]:
			return middle
		}

		return middle
	} else if task == 2 {
		return end
	} else if task == 1 {
		return start
	}

	return -1
}

func getMiddle(start int, end int) int {
	if end-start%2 == 0 {
		return (end + start - 1) / 2
	}
	return (end + start) / 2
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a < b {
		return b
	}

	return a
}

func swap(i int, j int, numbers []int) {
	var temp = numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp
}
