package main

import (
	"fmt"
	"michael/aoc/internal/problems"
	"os"
)

const PROB_NUM = 1

func main() {
	_, err := problems.Solve(PROB_NUM)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
