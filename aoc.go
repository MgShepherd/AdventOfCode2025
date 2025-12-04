package main

import (
	"fmt"
	"michael/aoc/internal/problems"
	"os"
)

const PROB_NUM = 4

func main() {
	result, err := problems.Solve(PROB_NUM)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	} else {
		fmt.Printf("The solution is: %d\n", result)
	}
}
