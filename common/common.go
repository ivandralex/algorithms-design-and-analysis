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
