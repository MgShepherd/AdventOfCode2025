package problems

import (
	"michael/aoc/internal/utils"
	"strings"
)

type position struct {
	x, y int
}

func solveProblem7() (int, error) {
	data, err := utils.ReadProblemFile(7)
	if err != nil {
		return -1, err
	}

	grid, startPos := readToGrid(strings.Split(strings.TrimSpace(data), "\n"))
	numSplits := traverseGrid(grid, startPos)
	return numSplits, nil
}

func readToGrid(lines []string) ([][]int, position) {
	grid := make([][]int, len(lines))
	startPos := position{x: 0, y: 0}

	for i, line := range lines {
		elements := make([]int, len(lines[i]))
		for j, el := range line {
			switch el {
			case '.':
				elements[j] = 0
			case 'S':
				startPos = position{x: j, y: i}
				fallthrough
			default:
				elements[j] = 1
			}
		}
		grid[i] = elements
	}

	return grid, startPos
}

func traverseGrid(grid [][]int, startPos position) int {
	numSplits := 0
	currentHeads := []position{startPos}
	for len(currentHeads) > 0 {
		nextHeads := []position{}
		for _, head := range currentHeads {
			if head.y+1 >= len(grid) {
				continue
			}

			switch grid[head.y+1][head.x] {
			case 0:
				grid[head.y+1][head.x] = 2
				nextHeads = append(nextHeads, position{x: head.x, y: head.y + 1})
			case 1:
				hasSplit := false
				if head.x > 0 && grid[head.y+1][head.x-1] == 0 {
					grid[head.y+1][head.x-1] = 2
					nextHeads = append(nextHeads, position{x: head.x - 1, y: head.y + 1})
					hasSplit = true
				}
				if head.x < len(grid[0])-1 && grid[head.y+1][head.x+1] == 0 {
					grid[head.y+1][head.x+1] = 2
					nextHeads = append(nextHeads, position{x: head.x + 1, y: head.y + 1})
					hasSplit = true
				}

				if hasSplit {
					numSplits++
				}
			}
		}
		currentHeads = nextHeads
	}
	return numSplits
}
