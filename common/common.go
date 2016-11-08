package common

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//ReadIntegers reads integers line by line
func ReadIntegers(path string) []int {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		numbers = append(numbers, num)
	}

	return numbers
}

//ReadAdjacencyLists reads adjency list from file
func ReadAdjacencyLists(path string) [][]int {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	graph := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//Init new adjacency list
		list := []int{}
		//Read and split line
		lineStr := scanner.Text()
		numbers := strings.Split(lineStr, "\t")

		for i := range numbers {
			number, err := strconv.Atoi(numbers[i])
			if err != nil {
				continue
			}

			list = append(list, number)
		}

		graph = append(graph, list)
	}

	return graph
}

//RandomInRange generates random in range
func RandomInRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//PrintGraph prints 2d array string by string
func PrintGraph(graph *[][]int) {
	for i := range *graph {
		fmt.Printf("%v\n", (*graph)[i])
	}
}

//SortResult of inversions counter procedure
type SortResult struct {
	Numbers []int
	Counter int
}

//SortAndCount sorts and counts inversed pairs
func SortAndCount(res *SortResult) *SortResult {
	var length = len(res.Numbers)

	if length < 2 {
		return res
	}

	var leftRes = new(SortResult)
	leftRes.Numbers = res.Numbers[0 : length/2]

	var rightRes = new(SortResult)
	rightRes.Numbers = res.Numbers[length/2:]

	//Find inversions in left-half subarray
	leftRes = SortAndCount(leftRes)
	//Find inversions in right-half subarray
	rightRes = SortAndCount(rightRes)
	//Find split inversions for which inverted numbers belong to different subarrays
	var merged = countSplitAndMerge(leftRes, rightRes)

	return merged
}

func countSplitAndMerge(left *SortResult, right *SortResult) *SortResult {
	var merged = new(SortResult)

	//fmt.Printf("Merging %v and %v\n", *left, *right)

	merged.Counter = left.Counter
	merged.Counter += right.Counter

	var i int
	var j int

	//Let's piggyback on merge sort's merge subroutine
	for i < len(left.Numbers) || j < len(right.Numbers) {
		if (i < len(left.Numbers) && j < len(right.Numbers) && left.Numbers[i] < right.Numbers[j]) || j == len(right.Numbers) {
			merged.Numbers = append(merged.Numbers, left.Numbers[i])
			i++
		} else {
			merged.Numbers = append(merged.Numbers, right.Numbers[j])
			j++

			//If we read from right subarray before reached the end of the left
			//Then we have inversions formed by current right number and all remaining left numbers
			if i < len(left.Numbers) {
				merged.Counter += len(left.Numbers) - i
			}
		}
	}

	//fmt.Printf("Merged to %v\n", *merged)

	return merged
}
