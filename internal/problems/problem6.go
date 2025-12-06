package problems

import (
	"fmt"
	"michael/aoc/internal/utils"
	"strconv"
	"strings"
)

func solveProblem6() (int, error) {
	data, err := utils.ReadProblemFile(6)
	if err != nil {
		return -1, err
	}

	lines := strings.Split(data, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	var endIdx int
	for endIdx = len(lines) - 1; len(lines[endIdx]) == 0; endIdx-- {
	}

	operations, err := getOperations(lines[endIdx])
	if err != nil {
		return -1, err
	}

	results, err := performArithmetic(operations, lines[:endIdx])
	if err != nil {
		return -1, err
	}

	return sumResults(results), nil
}

func getOperations(line string) ([]string, error) {
	elements := strings.Fields(line)
	for _, element := range elements {
		if element != "*" && element != "+" {
			return []string{}, fmt.Errorf("[ERROR] Unable to process operations line")
		}
	}

	return elements, nil
}

func performArithmetic(operations, lines []string) ([]int, error) {
	results := make([]int, len(operations))

	for _, line := range lines {
		for i, el := range strings.Fields(line) {
			intEl, err := strconv.Atoi(el)
			if err != nil {
				return []int{}, fmt.Errorf("[ERROR] Unable to convert %s to integer", el)
			}

			if operations[i] == "+" {
				results[i] += intEl
			} else {
				if results[i] == 0 {
					results[i] = 1
				}
				results[i] *= intEl
			}
		}
	}
	return results, nil
}

func sumResults(results []int) int {
	total := 0
	for _, el := range results {
		total += el
	}
	return total
}
