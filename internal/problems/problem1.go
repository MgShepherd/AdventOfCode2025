package problems

import (
	"fmt"
	"michael/aoc/internal/utils"
	"strconv"
	"strings"
)

const MAX_VALUE = 100

func solveProblem1() (int, error) {
	data, err := utils.ReadProblemFile(1)
	if err != nil {
		return -1, err
	}

	instructions := strings.Split(data, "\n")
	currentLocation, times0 := 50, 0

	for _, instruction := range instructions {
		if len(instruction) == 0 {
			continue
		}

		currentLocation, times0, err = processInstruction(instruction, currentLocation, times0)
		if err != nil {
			return -1, err
		}

		if currentLocation == 0 {
			times0++
		}
	}

	return times0, nil
}

func processInstruction(instruction string, currentLocation, times0 int) (int, int, error) {
	direction := instruction[0]
	amount, err := strconv.Atoi(strings.TrimSpace(instruction)[1:])
	if err != nil || (direction != 'L' && direction != 'R') {
		return -1, -1, fmt.Errorf("Unable to process instruction: %s\n", instruction)
	}

	if direction == 'L' {
		currentLocation -= amount
	} else {
		currentLocation += amount
	}

	currentLocation = currentLocation % MAX_VALUE
	if currentLocation < 0 {
		currentLocation = MAX_VALUE + currentLocation
	}

	return currentLocation, times0, nil
}
