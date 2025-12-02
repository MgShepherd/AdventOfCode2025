package problems

import (
	"fmt"
	"michael/aoc/internal/utils"
	"strconv"
	"strings"
)

func solveProblem2() (int, error) {
	data, err := utils.ReadProblemFile(2)
	if err != nil {
		return -1, err
	}

	ranges := strings.Split(data, ",")
	totalInvalid := 0
	for _, r := range ranges {
		start, end, err := convertToRangeElements(strings.TrimSpace(r))
		if err != nil {
			return -1, err
		}
		totalInvalid += getSumInvalidInRange(start, end)
	}
	return totalInvalid, nil
}

func convertToRangeElements(r string) (int, int, error) {
	splitIdx := strings.Index(r, "-")
	if splitIdx == -1 {
		return -1, -1, fmt.Errorf("[ERROR] Unable to get range components of %s\n", r)
	}

	start, err := strconv.Atoi(r[:splitIdx])
	if err != nil {
		return -1, -1, fmt.Errorf("[ERROR]: Unable to convert %s to integer", r[:splitIdx])
	}
	end, err := strconv.Atoi(r[splitIdx+1:])
	if err != nil {
		return -1, -1, fmt.Errorf("[ERROR]: Unable to convert %s to integer", r[:splitIdx])
	}

	return start, end, nil
}

func getSumInvalidInRange(start, end int) int {
	numInvalid := 0
	for i := start; i <= end; i++ {
		strVal := strconv.Itoa(i)
		if len(strVal)%2 != 0 {
			continue
		}

		midPoint := len(strVal) / 2
		invalid := true
		for j := range midPoint {
			if strVal[j] != strVal[midPoint+j] {
				invalid = false
				break
			}
		}

		if invalid {
			numInvalid += i
		}
	}
	return numInvalid
}
