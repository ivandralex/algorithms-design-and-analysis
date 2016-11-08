package main

import (
	"fmt"

	"./common"
)

//var testCase1 = []int{8, 7, 6, 5, 4, 3, 2, 1}
var testCase1 = []int{4, 9, 8, 7, 5, 6, 3, 2, 1}

func main() {
	var res = new(common.SortResult)

	if false {
		res.Numbers = testCase1
	} else {
		res.Numbers = common.ReadIntegers("./data/2sum.txt")
	}

	var result = common.SortAndCount(res)

	fmt.Printf("Inversions found: %d\n", result.Counter)
}
