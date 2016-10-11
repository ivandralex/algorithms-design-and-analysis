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

//ReadEdges reads adjency list from file and returns (adjency list, vertex index) pair
func ReadEdges(path string) (*[][]int, *[]int) {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	vertexInitMap := make(map[int]bool)
	indices := []int{}
	adjacencyMap := make(map[int][]int)
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

		addVertex(edge[0], &indices, &vertexInitMap, &adjacencyMap, &vertices)
		addVertex(edge[1], &indices, &vertexInitMap, &adjacencyMap, &vertices)

		adjacencyMap[edge[0]] = append(adjacencyMap[edge[0]], edge[1])
		if edge[0] != edge[1] {
			adjacencyMap[edge[1]] = append(adjacencyMap[edge[1]], -edge[0])
		}
	}

	for _, v := range vertices {
		graph = append(graph, adjacencyMap[v])
	}

	return &graph, &indices
}

func addVertex(v int, indices *[]int, vertexInitMap *map[int]bool, adjacencyMap *map[int][]int, vertices *[]int) {
	//Init map for edge tail
	if _, ok := (*vertexInitMap)[v]; !ok {
		//Mark vertex as initialized
		(*vertexInitMap)[v] = true
		//Save index to indices array
		*indices = append(*indices, len(*vertices))
		//Init adjacency list for vertex v
		(*adjacencyMap)[v] = []int{v}
		//Save key of adjaceny list
		*vertices = append(*vertices, v)
	}
}

func main() {
	graph, indices := ReadEdges("./data/SCC.txt")
	fmt.Println("Finished reading edges")

	//Map of explored nodes
	explored := make(map[int]bool)
	//Array of list indices sorted by finishing times of the first pass of dfsLoop
	finishingTimes := []int{}
	//Number of processed nodes
	var numProcessed int

	//First pass on inverted graph
	dfsLoop(graph, indices, -1, &explored, &finishingTimes, &numProcessed)

	//TODO: collect finishing times

	//Second pass on the original graph
	dfsLoop(graph, indices, 1, &explored, &finishingTimes, &numProcessed)
}

func dfsLoop(graph *[][]int, indices *[]int, factor int, explored *map[int]bool, finishingTimes *[]int, numProcessed *int) {
	//Indices graph loop ordering
	for _, index := range *indices {
		i := (*graph)[index][0]
		//if i not yet explored
		if _, ok := (*explored)[i]; !ok {
			//TODO: assign s
			dfs(graph, index, explored, finishingTimes, numProcessed)
		}
	}
}

func dfs(graph *[][]int, index int, explored *map[int]bool, finishingTimes *[]int, numProcessed *int) {
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
