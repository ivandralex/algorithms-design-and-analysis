package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
func ReadEdges(path string) (*[][]int, *[]int, *map[int]int) {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	vertexIndexMap := make(map[int]int)
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

		addVertex(edge[0], &indices, &vertexIndexMap, &adjacencyMap, &vertices)
		addVertex(edge[1], &indices, &vertexIndexMap, &adjacencyMap, &vertices)

		adjacencyMap[edge[0]] = append(adjacencyMap[edge[0]], edge[1])
		if edge[0] != edge[1] {
			adjacencyMap[edge[1]] = append(adjacencyMap[edge[1]], -edge[0])
		}
	}

	for _, v := range vertices {
		graph = append(graph, adjacencyMap[v])
	}

	return &graph, &indices, &vertexIndexMap
}

func addVertex(v int, indices *[]int, vertexIndexMap *map[int]int, adjacencyMap *map[int][]int, vertices *[]int) {
	//Init map for edge tail
	if _, ok := (*vertexIndexMap)[v]; !ok {
		//Save vertex index to the map thus mark it as initialized
		currentIndex := len(*vertices)
		(*vertexIndexMap)[v] = currentIndex
		//Save index to indices array
		*indices = append(*indices, currentIndex)
		//Init adjacency list for vertex v
		(*adjacencyMap)[v] = []int{v}
		//Save key of adjaceny list
		*vertices = append(*vertices, v)
	}
}

func main() {
	graph, indices, vertexIndexMap := ReadEdges("./data/SCC.txt")
	fmt.Println("Finished reading edges")

	//Map of explored nodes
	explored := make(map[int]bool)
	//Array of list indices sorted by finishing times of the first pass of dfsLoop
	finishingTimes := []int{}

	//First pass on inverted graph`
	dfsLoop(graph, indices, vertexIndexMap, -1, &explored, &finishingTimes, nil)

	fmt.Println("Finished inverted pass")

	//Reset explored vertices
	explored = make(map[int]bool)
	leaders := make(map[int]int)

	//Second pass on the original graph in order defined by reversed finishing times
	dfsLoop(graph, &finishingTimes, vertexIndexMap, 1, &explored, nil, &leaders)

	sizes := []int{}

	for _, size := range leaders {
		sizes = append(sizes, size)
	}

	sort.Ints(sizes)

	fmt.Printf("Top 5 SCCs sizes: %v\n", sizes[len(sizes)-5:])
}

func dfsLoop(graph *[][]int, indices *[]int, vertexIndexMap *map[int]int, factor int, explored *map[int]bool, finishingTimes *[]int, leaders *map[int]int) {
	var numProcessed int
	//Indices graph loop ordering
	for k := len(*indices) - 1; k >= 0; k-- {
		index := (*indices)[k]
		i := (*graph)[index][0]
		//if i not yet explored
		if _, ok := (*explored)[i]; !ok {
			fmt.Printf("%d$ Vertex %d not explored\n", factor, i)
			numProcessed = 0
			dfs(graph, vertexIndexMap, index, factor, explored, finishingTimes, &numProcessed)
			if factor == 1 {
				//Save number of vertices in SSC with leader i
				(*leaders)[i] = numProcessed
			}
		}
	}
}

func dfs(graph *[][]int, vertexIndexMap *map[int]int, index int, factor int, explored *map[int]bool, finishingTimes *[]int, numProcessed *int) {
	//We are executing DFS for vertex defined by adjacency list graph[index]

	//Get adjacency list
	list := (*graph)[index]
	//Vertex for which we run DFS
	i := list[0]

	fmt.Printf("%d$ Started DFS on vertex %d\n", factor, i)

	//Mark i as explored
	(*explored)[i] = true

	//for each (i, j) from G ...
	for _, j := range list {
		//We only follow edges in direction defined by factor
		if j == i || (factor == -1 && j > 0) || (factor == 1 && j < 0) {
			continue
		}

		absJ := factor * j

		//if j not yet explored
		if _, ok := (*explored)[absJ]; !ok {
			fmt.Printf("%d$ %d DFS recurses on vertex %d\n", factor, i, j)
			dfs(graph, vertexIndexMap, (*vertexIndexMap)[absJ], factor, explored, finishingTimes, numProcessed)
		}
	}
	(*numProcessed)++
	if factor == -1 {
		*finishingTimes = append(*finishingTimes, index)
	}
	fmt.Printf("%d$ Finished DFS on vertex %d\n", factor, i)
}
