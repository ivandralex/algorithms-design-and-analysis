package main

import "./common"

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func main() {
	graph := common.ReadEdges("./data/SCC.txt")

	common.PrintGraph(&graph)

	//common.PrintGraph(&graph)
}
