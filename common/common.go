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

//Graph represented by adjacency list
type Graph [][]int

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

//ReadEdges reads adjency list from file
func ReadEdges(path string) [][]int {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	verticesMap := make(map[int]map[int]bool)
	adjancencyMap := make(map[int][]int)
	vertices := []int{}

	graph := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//Read and split line
		lineStr := scanner.Text()
		numbers := strings.Split(lineStr, " ")

		edge := []int{}

		for _, str := range numbers {
			number, err := strconv.Atoi(str)
			if err != nil {
				continue
			}

			edge = append(edge, number)
		}

		//Init map
		if _, ok := verticesMap[edge[0]]; !ok {
			verticesMap[edge[0]] = make(map[int]bool)
			adjancencyMap[edge[0]] = []int{edge[0]}
			vertices = append(vertices, edge[0])
		}

		if _, ok := verticesMap[edge[0]][edge[1]]; !ok {
			verticesMap[edge[0]][edge[1]] = true
			adjancencyMap[edge[0]] = append(adjancencyMap[edge[0]], edge[1])
		}
	}

	for _, v := range vertices {
		graph = append(graph, adjancencyMap[v])
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
