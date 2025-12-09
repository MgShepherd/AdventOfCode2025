package problems

import (
	"fmt"
	"math"
	"michael/aoc/internal/utils"
	"strconv"
	"strings"
)

const POSITION_COMPONENTS = 2

func solveProblem9() (int, error) {
	data, err := utils.ReadProblemFile(9)
	if err != nil {
		return -1, err
	}

	positions, err := readToPositions(strings.Split(strings.TrimSpace(data), "\n"))
	if err != nil {
		return -1, err
	}

	return getMaxArea(positions), nil
}

func readToPositions(lines []string) ([]position, error) {
	positions := make([]position, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		elements := strings.Split(line, ",")
		if len(elements) != POSITION_COMPONENTS {
			return []position{}, fmt.Errorf("[ERROR] Unable to convert %s into position\n", line)
		}

		xPos, err := strconv.Atoi(elements[0])
		if err != nil {
			return []position{}, fmt.Errorf("[ERROR] Unable to convert %s into x coordinate\n", elements[0])
		}

		yPos, err := strconv.Atoi(elements[1])
		if err != nil {
			return []position{}, fmt.Errorf("[ERROR] Unable to convert %s into y coordinate\n", elements[0])
		}

		positions[i] = position{x: xPos, y: yPos}
	}

	return positions, nil
}

func getMaxArea(positions []position) int {
	maxArea := 0
	for i, pos1 := range positions {
		for j := i + 1; j < len(positions); j++ {
			xDistance := math.Abs(float64(positions[j].x-pos1.x)) + 1
			yDistance := math.Abs(float64(positions[j].y-pos1.y)) + 1
			area := int(xDistance * yDistance)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
