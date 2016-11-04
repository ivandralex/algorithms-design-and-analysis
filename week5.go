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

	//graph := [][]int{}

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

			//fmt.Printf("edgeParts: %v\n", edgeParts)

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
	adjacencyMap, _ := ReadEdges("./data/dijkstra.txt")
	fmt.Printf("Finished reading edges: %v\n", (*adjacencyMap)[1])

	fmt.Printf("10 shortest distances:\n")
}
