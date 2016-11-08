package main

import (
	"fmt"

	"./common"
)

func main() {
	var res = new(common.SortResult)

	res.Numbers = common.ReadIntegers("./data/2sum.txt")

	fmt.Printf("Total numbers: %d\n", len(res.Numbers))

	hashMap := make(map[int]bool)

	//Fill in hashMap
	for _, i := range res.Numbers {
		hashMap[i] = true
	}

	fmt.Printf("Built hash\n")

	targetPairs := 0
	processed := 0

	for t := -10000; t < 10001; t++ {
		for _, x := range res.Numbers {
			y := t - x

			if y == x {
				fmt.Printf("Continue: %d\n", x)
				continue
			}

			if _, ok := hashMap[y]; ok {
				targetPairs++
				break
			}
		}
		processed++

		if processed%1000 == 0 {
			fmt.Printf("Processed so far: %d\n", processed)
		}
	}

	fmt.Printf("Target pairs: %d\n", targetPairs)
}
