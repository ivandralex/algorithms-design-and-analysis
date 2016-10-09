package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

//ReadEdges reads adjency list from file
func ReadEdges(path string) *[][]int {
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

	return &graph
}

func main() {
	graph := ReadEdges("./data/SCC.txt")
	fmt.Println("Finished reading edges")

	//common.PrintGraph(graph)

	//Map of explored nodes
	explored := make(map[int]bool)
	//Finishing times
	finishingTimes := make(map[int]int)
	//Number of processed nodes
	var numProcessed int

	//First pass on inverted graph
	dfsLoop(graph, true, &explored, &finishingTimes, &numProcessed)

	//TODO: collect finishing times

	//Second pass on the original graph
	dfsLoop(graph, false, &explored, &finishingTimes, &numProcessed)

	//common.PrintGraph(&graph)
}

func dfsLoop(graph *[][]int, inverted bool, explored *map[int]bool, finishingTimes *map[int]int, numProcessed *int) {
	for index, list := range *graph {
		i := list[0]
		//if i not yet explored
		if _, ok := (*explored)[i]; !ok {
			//TODO: assign s
			dfs(graph, index, explored, finishingTimes, numProcessed)
		}
	}
}

func dfs(graph *[][]int, index int, explored *map[int]bool, finishingTimes *map[int]int, numProcessed *int) {
	//Restore list from dfsLoop
	list := (*graph)[index]
	i := list[0]

	return

	//Mark i as explored
	(*explored)[i] = true

	//for each (i, j) from G ...
	for _, j := range list {
		if j == i {
			continue
		}

		//if j not yet explored
		if _, ok := (*explored)[j]; !ok {
			indexOfJsVertex := j
			dfs(graph, indexOfJsVertex, explored, finishingTimes, numProcessed)
		}
	}
	(*numProcessed)++
	(*finishingTimes)[i] = *numProcessed
}
