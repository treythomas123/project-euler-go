package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max := 1000
	if len(os.Args) > 1 {
		maxArg, err := strconv.Atoi(os.Args[1])
		if err == nil {
			max = maxArg
		}
	}
	fmt.Println(solution(max))
}

func solution(max int) int {
	sum := 0
	for i := 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
