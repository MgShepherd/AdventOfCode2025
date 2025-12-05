package problems

import (
	"fmt"
	"michael/aoc/internal/utils"
	"strconv"
	"strings"
)

const RANGE_COMPONENTS = 2

func solveProblem5() (int, error) {
	data, err := utils.ReadProblemFile(5)
	if err != nil {
		return -1, err
	}

	return processData(data)
}

func processData(data string) (int, error) {
	rangeMap := make(map[int]int)

	for _, line := range strings.Split(data, "\n") {
		element := strings.TrimSpace(line)
		if len(element) == 0 {
			break
		}

		err := addToRangeMap(rangeMap, element)
		if err != nil {
			return -1, err
		}
	}

	performFinalPass(rangeMap)

	return getTotalValuesWithinRanges(rangeMap), nil
}

func addToRangeMap(rangeMap map[int]int, element string) error {
	values := strings.Split(element, "-")
	if len(values) != RANGE_COMPONENTS {
		return fmt.Errorf("[ERROR] Unable to process range element: %s\n", element)
	}

	start, err := strconv.Atoi(values[0])
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to process first component of range %s\n", element)
	}
	end, err := strconv.Atoi(values[1])
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to process end component of range %s\n", element)
	}

	updated := combineKeys(rangeMap, start, end)

	if !updated {
		rangeMap[start] = end
	}
	return nil
}

func performFinalPass(rangeMap map[int]int) {
	for key, val := range rangeMap {
		delete(rangeMap, key)
		combined := combineKeys(rangeMap, key, val)
		if !combined {
			rangeMap[key] = val
		}
	}
}

func combineKeys(rangeMap map[int]int, start, end int) bool {
	updated := false
	for key, val := range rangeMap {
		if start <= key && end >= key {
			delete(rangeMap, key)
			if end > val {
				rangeMap[start] = end
			} else {
				rangeMap[start] = val
			}
			updated = true
		} else if end >= val && start <= val {
			rangeMap[key] = end
			updated = true
		}

		if updated {
			return true
		}
	}

	return false
}

func getTotalValuesWithinRanges(rangeMap map[int]int) int {
	numValues := 0
	for key, val := range rangeMap {
		numValues += val - key + 1
	}
	return numValues
}
