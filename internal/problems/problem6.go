package problems

import (
	"fmt"
	"michael/aoc/internal/utils"
	"strconv"
	"strings"
)

const MAX_COLS = 6

func solveProblem6() (int, error) {
	data, err := utils.ReadProblemFile(6)
	if err != nil {
		return -1, err
	}

	lines := strings.Split(data, "\n")
	var endIdx int
	for endIdx = len(lines) - 1; len(lines[endIdx]) == 0; endIdx-- {
	}

	operations, err := getOperations(strings.TrimSpace(lines[endIdx]))
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
	problems := make([][]string, len(operations))

	idxs := getSeperatorIndexes(lines)

	currentSepIdx := 0
	var current strings.Builder
	for i := range lines {
		currentSepIdx = 0
		for j := range lines[i] {
			if currentSepIdx < len(idxs) && j == idxs[currentSepIdx] {
				problems[currentSepIdx] = append(problems[currentSepIdx], current.String())
				currentSepIdx++
				current.Reset()
			} else if lines[i][j] == ' ' {
				current.WriteRune('0')
			} else {
				current.WriteByte(lines[i][j])
			}
		}

		if current.Len() > 0 {
			problems[currentSepIdx] = append(problems[currentSepIdx], current.String())
			current.Reset()
		}
	}

	for i, problem := range problems {
		operands, err := getOperands(problem)
		if err != nil {
			return []int{}, err
		}

		if operations[i] == "+" {
			results[i] = getAdditionResult(operands)
		} else {
			results[i] = getMultiplicationResult(operands)
		}
	}
	return results, nil
}

func getSeperatorIndexes(lines []string) []int {
	var indexes []int
	for i, el := range lines[0] {
		if el == ' ' {
			fullEmpty := true
			for _, compare := range lines {
				if compare[i] != ' ' {
					fullEmpty = false
					break
				}
			}
			if fullEmpty {
				indexes = append(indexes, i)
			}
		}
	}
	return indexes
}

func getAdditionResult(operands []int) int {
	result := 0
	for _, el := range operands {
		result += el
	}
	return result
}

func getMultiplicationResult(operands []int) int {
	result := 1
	for _, el := range operands {
		result *= el
	}
	return result
}

func getOperands(elements []string) ([]int, error) {
	columnVals := make([]strings.Builder, MAX_COLS)
	var operands []int

	for _, element := range elements {
		cols := strings.Split(element, "")
		for i, el := range cols {
			if el != "0" {
				columnVals[len(cols)-1-i].WriteString(el)
			}
		}
	}

	for _, v := range columnVals {
		if len(v.String()) == 0 {
			break
		}
		intVal, err := strconv.Atoi(v.String())
		if err != nil {
			return []int{}, fmt.Errorf("[ERROR] Unable to convert %s to integer", v.String())
		}
		operands = append(operands, intVal)
	}
	return operands, nil
}

func sumResults(results []int) int {
	total := 0
	for _, el := range results {
		total += el
	}
	return total
}
