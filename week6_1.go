package main

import (
	"fmt"

	"./common"
)

func main() {
	numbers := common.ReadIntegers("./data/2sum.txt")

	fmt.Printf("Total numbers: %d\n", len(numbers))

	hashMap := make(map[int]bool)

	//Fill in hashMap
	for _, i := range numbers {
		hashMap[i] = true
	}

	fmt.Printf("Built hash\n")

	targetPairs := make(chan int, 20001)

	for t := -10000; t < 10001; t++ {
		go checkT(t, &hashMap, &numbers, targetPairs)
	}

	sum := 0
	cloe(targetPairs)

	for ok := range targetPairs {
		sum += ok
	}

	fmt.Printf("Target pairs: %d\n", sum)
}

func checkT(t int, hashMap *map[int]bool, numbers *[]int, c chan int) {
	for _, x := range *numbers {
		y := t - x

		if y == x {
			continue
		}

		if _, ok := (*hashMap)[y]; ok {
			c <- 1
			return
		}
	}

	c <- 0
}
