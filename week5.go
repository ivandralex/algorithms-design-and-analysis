package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Vertex vertex
type Vertex struct {
	head   int
	length int
}

//ReadEdges reads adjency list from file
func ReadEdges(path string) (*map[int][]*Vertex, *[]int) {
	var file, err = os.Open(path)

	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	adjacencyMap := make(map[int][]*Vertex)
	vertices := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//Read and split line
		lineStr := scanner.Text()
		numbers := strings.Split(lineStr, "\t")
		var list []*Vertex
		var vertex int

		for i, str := range numbers {
			if i == 0 {
				vertex, _ = strconv.Atoi(str)

				if err != nil {
					fmt.Printf("Failed to parse \"%s\"\n", str)
					continue
				}

				list = []*Vertex{}
				continue
			}

			if str == "" {
				continue
			}

			edgeParts := strings.Split(str, ",")

			edge := new(Vertex)
			edge.head, _ = strconv.Atoi(edgeParts[0])
			edge.length, _ = strconv.Atoi(edgeParts[1])

			list = append(list, edge)
		}

		vertices = append(vertices, vertex)
		adjacencyMap[vertex] = list
	}

	return &adjacencyMap, &vertices
}

func main() {
	adjacencyMap, vertices := ReadEdges("./data/dijkstra.txt")
	fmt.Printf("Finished reading edges")

	minPaths := findMinPaths(1, adjacencyMap, vertices)

	targetVertices := []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}

	var length int
	var targetPaths []string

	for _, v := range targetVertices {
		if _, ok := (*minPaths)[v]; ok {
			length = (*minPaths)[v]
		} else {
			length = 1000000
		}
		targetPaths = append(targetPaths, strconv.Itoa(length))
	}

	pathsStr := strings.Join(targetPaths, ",")

	fmt.Printf("10 shortest distances: %s\n", pathsStr)
}

//findMinPaths fins min paths for all vertices reachable from source vertex
func findMinPaths(source int, adjacencyMap *map[int][]*Vertex, vertices *[]int) *map[int]int {
	//Maps vertex to the lendth of the shortest path to this vertex
	minPaths := make(map[int]int)
	minPaths[source] = 0

	//Vertices with assigned shortest path lengths
	g := []int{source}

	//Vertices that do not belong to the frontier
	innerVertices := make(map[int]bool)

	var minPath int
	var minVertex int
	var onFrontier bool

	for len(g) < len(*vertices) {
		//Assign minPath with max possible length
		minPath = 1000000

		//Add new vertex to g
		for _, v := range g {
			//If all v neighbours already added to g
			if _, ok := innerVertices[v]; ok {
				continue
			}

			onFrontier = false

			//Griddy vertex selction
			for _, vertex := range (*adjacencyMap)[v] {
				if _, ok := minPaths[vertex.head]; ok {
					continue
				}

				onFrontier = true

				if minPaths[v]+vertex.length < minPath {
					minPath = minPaths[v] + vertex.length
					minVertex = vertex.head
				}
			}

			//v is no longer on frontier
			if !onFrontier {
				innerVertices[v] = true
			}
		}

		minPaths[minVertex] = minPath
		g = append(g, minVertex)
	}

	return &minPaths
}
