package main

import (
	"fmt"

	"./common"
)

func main() {
	graph := common.ReadAdjacencyLists("./data/kargerMinCut0.txt")

	fmt.Printf("Initial graph:\n")
	printGraph(&graph)

	for len(graph) > 2 {
		contractRandomEdge(&graph)
	}

	fmt.Printf("After contraction:\n")
	printGraph(&graph)
}

func contractRandomEdge(graph *[][]int) {
	numVertices := len(*graph)

	//Adjacency list where we will choose vertex for contraction
	vertex1Index := common.RandomInRange(0, numVertices-1)
	vertex1List := (*graph)[vertex1Index]
	//Index of vertex that we will contract
	vertex1 := vertex1List[0]
	vertex2 := vertex1List[common.RandomInRange(1, len(vertex1List)-1)]

	//Find vertex2 Index
	var vertex2Index int
	for i := range *graph {
		if (*graph)[i][0] == vertex2 {
			vertex2Index = i
			break
		}

	}

	//Replace vertex2 with vertex1 in all vertices except for vertex1
	for i := range *graph {
		if (*graph)[i][0] == vertex1 {
			for j := range (*graph)[vertex2Index] {
				if j == 0 || (*graph)[vertex2Index][j] == vertex1 {
					continue
				}
				(*graph)[vertex1Index] = append((*graph)[vertex1Index], (*graph)[vertex2Index][j])
			}
		} else {
			for j := range (*graph)[i] {
				if j == 0 {
					continue
				}
				if (*graph)[i][j] == vertex2 {
					(*graph)[i][j] = vertex1
				}
			}
		}
	}

	//Remove vertex2
	*graph = append((*graph)[:vertex2Index], (*graph)[vertex2Index+1:]...)

	fmt.Printf("Step: %d --> %d\n", vertex2, vertex1)
	printGraph(graph)
}

func printGraph(graph *[][]int) {
	for i := range *graph {
		fmt.Printf("%v\n", (*graph)[i])
	}
}
