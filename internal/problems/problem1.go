package problems

import (
	"fmt"
	"michael/aoc/internal/utils"
)

func solveProblem1() (int, error) {
	data, err := utils.ReadProblemFile(1)
	if err != nil {
		return -1, err
	}

	fmt.Printf("Successfully read problem 1 data: %s\n", data)
	return 0, nil
}
