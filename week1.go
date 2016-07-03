package main

import (
	"fmt"

	"./common"
)

//var testCase1 = []int{8, 7, 6, 5, 4, 3, 2, 1}
var testCase1 = []int{4, 9, 8, 7, 5, 6, 3, 2, 1}

//Result of inversions counter procedure
type Result struct {
	numbers []int
	counter int
}

func main() {
	var res = new(Result)

	if false {
		res.numbers = testCase1
	} else {
		res.numbers = common.ReadIntegers("./data/week1.txt")
	}

	var result = sortAndCount(res)

	fmt.Printf("Inversions found: %d\n", result.counter)
}

func sortAndCount(res *Result) *Result {
	var length = len(res.numbers)

	if length < 2 {
		return res
	}

	var leftRes = new(Result)
	leftRes.numbers = res.numbers[0 : length/2]

	var rightRes = new(Result)
	rightRes.numbers = res.numbers[length/2:]

	//Find inversions in left-half subarray
	leftRes = sortAndCount(leftRes)
	//Find inversions in right-half subarray
	rightRes = sortAndCount(rightRes)
	//Find split inversions for which inverted numbers belong to different subarrays
	var merged = countSplitAndMerge(leftRes, rightRes)

	return merged
}

func countSplitAndMerge(left *Result, right *Result) *Result {
	var merged = new(Result)

	//fmt.Printf("Merging %v and %v\n", *left, *right)

	merged.counter = left.counter
	merged.counter += right.counter

	var i int
	var j int

	//Let's piggyback on merge sort's merge subroutine
	for i < len(left.numbers) || j < len(right.numbers) {
		if (i < len(left.numbers) && j < len(right.numbers) && left.numbers[i] < right.numbers[j]) || j == len(right.numbers) {
			merged.numbers = append(merged.numbers, left.numbers[i])
			i++
		} else {
			merged.numbers = append(merged.numbers, right.numbers[j])
			j++

			//If we read from right subarray before reached the end of the left
			//Then we have inversions formed by current right number and all remaining left numbers
			if i < len(left.numbers) {
				merged.counter += len(left.numbers) - i
			}
		}
	}

	//fmt.Printf("Merged to %v\n", *merged)

	return merged
}
