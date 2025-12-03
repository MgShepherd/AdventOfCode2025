package problems

import (
	"fmt"
	"michael/aoc/internal/utils"
	"strconv"
	"strings"
)

const LARGEST_LEN = 12

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
	largestValues := make([]int, LARGEST_LEN)

	for i, el := range bank {
		for currentIdx := range largestValues {
			if el > largestValues[currentIdx] && i < len(bank)-len(largestValues)+currentIdx+1 {
				largestValues[currentIdx] = el
				for j := currentIdx + 1; j < len(largestValues); j++ {
					largestValues[j] = 0
				}
				break
			}
		}
	}

	strVal := convertIntSliceToString(largestValues)
	result, _ := strconv.Atoi(strVal)
	return result
}

func convertIntSliceToString(elements []int) string {
	var b strings.Builder
	for _, el := range elements {
		b.WriteString(strconv.Itoa(el))
	}

	return b.String()
}
