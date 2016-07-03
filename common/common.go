package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
