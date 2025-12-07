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
	numTimelines := getTimelinesFromLocation(grid, startPos, make(map[position]int))
	return numTimelines, nil
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

func getTimelinesFromLocation(grid [][]int, current position, timelinesFromPos map[position]int) int {
	if val, ok := timelinesFromPos[current]; ok {
		return val
	}

	if current.y+1 >= len(grid) {
		return 1
	}

	total := 0
	switch grid[current.y+1][current.x] {
	case 0:
		total = getTimelinesFromLocation(grid, position{x: current.x, y: current.y + 1}, timelinesFromPos)
	case 1:
		if current.x > 0 {
			total += getTimelinesFromLocation(grid, position{x: current.x - 1, y: current.y + 1}, timelinesFromPos)
		}
		if current.x < len(grid[0])-1 {
			total += getTimelinesFromLocation(grid, position{x: current.x + 1, y: current.y + 1}, timelinesFromPos)
		}
	}

	timelinesFromPos[current] = total
	return total
}
