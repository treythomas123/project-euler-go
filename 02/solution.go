package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max := 4000000
	if len(os.Args) > 1 {
		maxArg, err := strconv.Atoi(os.Args[1])
		if err == nil {
			max = maxArg
		}
	}
	fmt.Println(solution(max))
}

func solution(max int) (sum int) {
	a := 0
	b := 1
	for b <= max {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return
}
