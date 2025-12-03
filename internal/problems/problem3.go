package problems

import (
	"fmt"
	"michael/aoc/internal/utils"
	"strconv"
	"strings"
)

func solveProblem3() (int, error) {
	data, err := utils.ReadProblemFile(3)
	if err != nil {
		return -1, err
	}

	totalJoltage := 0
	for _, bank := range strings.Split(data, "\n") {
		if len(bank) == 0 {
			continue
		}
		elements, err := convertToIntSlice(strings.TrimSpace(bank))
		if err != nil {
			return -1, err
		}
		totalJoltage += getLargestJoltage(elements)
	}

	return totalJoltage, nil
}

func convertToIntSlice(line string) ([]int, error) {
	intSlice := make([]int, len(line))

	for i, el := range line {
		intEl, err := strconv.Atoi(string(el))
		if err != nil {
			return []int{}, fmt.Errorf("[ERROR] Unable to process element: %c\n", el)
		}
		intSlice[i] = intEl
	}

	return intSlice, nil
}

func getLargestJoltage(bank []int) int {
	largestValues := make([]int, 2)

	for i, el := range bank {
		if el > largestValues[0] && i < len(bank)-1 {
			largestValues[0] = el
			largestValues[1] = 0
		} else if el > largestValues[1] {
			largestValues[1] = el
		}
	}

	largestVal := fmt.Sprintf("%d%d", largestValues[0], largestValues[1])
	result, _ := strconv.Atoi(largestVal)
	return result
}
