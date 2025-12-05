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
	processingRanges := true
	rangeMap := make(map[int]int)
	numFresh := 0

	for _, line := range strings.Split(data, "\n") {
		element := strings.TrimSpace(line)
		if len(element) == 0 {
			processingRanges = false
			continue
		}

		var err error
		if processingRanges {
			err = addRangeToMap(rangeMap, element)
		} else {
			var fresh bool
			fresh, err = checkIfFresh(rangeMap, element)
			if fresh {
				numFresh++
			}
		}

		if err != nil {
			return -1, err
		}
	}

	return numFresh, nil
}

func addRangeToMap(rangeMap map[int]int, element string) error {
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

	if rangeMap[start] == 0 || rangeMap[start] < end {
		rangeMap[start] = end
	}
	return nil
}

func checkIfFresh(rangeMap map[int]int, element string) (bool, error) {
	intEl, err := strconv.Atoi(element)
	fmt.Printf("Checking int element: %d\n", intEl)
	if err != nil {
		return false, fmt.Errorf("[ERROR] Unable to process element %s\n", element)
	}

	for key, val := range rangeMap {
		if key <= intEl && val >= intEl {
			return true, nil
		}
	}

	return false, nil
}
