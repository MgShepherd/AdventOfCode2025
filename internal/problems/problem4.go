package problems

import (
	"michael/aoc/internal/utils"
	"strings"
)

func solveProblem4() (int, error) {
	data, err := utils.ReadProblemFile(4)
	if err != nil {
		return -1, err
	}

	grid := buildGrid(strings.Split(data, "\n"))
	return getNumValid(grid), nil
}

func buildGrid(lines []string) [][]int {
	var grid [][]int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, readLineToSlice(strings.TrimSpace(line)))
	}
	return grid
}

func readLineToSlice(line string) []int {
	intSlice := make([]int, len(line))
	for i, el := range line {
		if el == '@' {
			intSlice[i] = 1
		} else {
			intSlice[i] = 0
		}
	}
	return intSlice
}

func getNumValid(grid [][]int) int {
	numValid := 0
	numRemoved := 0
	for keepChecking := true; keepChecking; keepChecking = numRemoved > 0 {
		numRemoved = 0

		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == 1 && isValid(grid, x, y) {
					grid[y][x] = 0
					numRemoved++
					numValid++
				}
			}
		}
	}
	return numValid
}

func isValid(grid [][]int, x, y int) bool {
	numSurrounding := 0
	for currY := y - 1; currY <= y+1; currY++ {
		for currX := x - 1; currX <= x+1; currX++ {
			if (currX == x && currY == y) || currY < 0 || currY >= len(grid) || currX < 0 || currX >= len(grid[currY]) {
				continue
			}

			if grid[currY][currX] == 1 {
				numSurrounding++
			}
		}
	}
	return numSurrounding < 4
}
