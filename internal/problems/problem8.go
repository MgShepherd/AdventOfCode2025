package problems

import (
	"fmt"
	"maps"
	"michael/aoc/internal/utils"
	"slices"
	"strconv"
	"strings"
)

const COORD_SIZE = 3
const NUM_LARGEST_CIRCUITS = 3

type coord struct {
	x, y, z int
}

func solveProblem8() (int, error) {
	data, err := utils.ReadProblemFile(8)
	if err != nil {
		return -1, err
	}

	coords, err := readToCoords(strings.Split(strings.TrimSpace(data), "\n"))
	if err != nil {
		return -1, err
	}

	finalConnection, err := getCompletingConnection(getDistances(coords), len(coords))
	if err != nil {
		return -1, err
	}

	return finalConnection[0].x * finalConnection[1].x, nil
}

func readToCoords(lines []string) ([]coord, error) {
	coords := make([]coord, len(lines))

	for i, line := range lines {
		if len(line) == 0 {
			break
		}
		line = strings.TrimSpace(line)
		elements := strings.Split(line, ",")
		if len(elements) != COORD_SIZE {
			return []coord{}, fmt.Errorf("[ERROR] Unable to process coordinate %s\n", line)
		}

		parsedInts := make([]int, COORD_SIZE)
		for i, el := range elements {
			intEl, err := strconv.Atoi(el)
			if err != nil {
				return []coord{}, fmt.Errorf("[ERROR] Unable to convert coordinate element %s to integer\n", el)
			}
			parsedInts[i] = intEl
		}

		coords[i] = coord{x: parsedInts[0], y: parsedInts[1], z: parsedInts[2]}
	}

	return coords, nil
}

func getDistances(coords []coord) map[int][]coord {
	distances := make(map[int][]coord)
	for _, el1 := range coords {
		for _, el2 := range coords {
			if el1 == el2 {
				continue
			}

			distance := (el1.x-el2.x)*(el1.x-el2.x) + (el1.y-el2.y)*(el1.y-el2.y) + (el1.z-el2.z)*(el1.z-el2.z)
			distances[distance] = []coord{el1, el2}
		}
	}

	return distances
}

func getCompletingConnection(distances map[int][]coord, numCoords int) ([]coord, error) {
	keys := slices.Collect(maps.Keys(distances))
	slices.Sort(keys)

	circuits := [][]coord{}
	for i := range len(keys) {
		el1Idx, el2Idx := -1, -1
		for j, c := range circuits {
			for _, el := range c {
				if el == distances[keys[i]][0] {
					el1Idx = j
				} else if el == distances[keys[i]][1] {
					el2Idx = j
				}

				if el1Idx != -1 && el2Idx != -1 {
					break
				}
			}
			if el1Idx != -1 && el2Idx != -1 {
				break
			}
		}

		if el1Idx != -1 && el1Idx == el2Idx {
			continue
		} else if el1Idx != -1 && el2Idx == -1 {
			circuits[el1Idx] = append(circuits[el1Idx], distances[keys[i]][1])
			if len(circuits[el1Idx]) == numCoords {
				return distances[keys[i]], nil
			}
		} else if el1Idx == -1 && el2Idx != -1 {
			circuits[el2Idx] = append(circuits[el2Idx], distances[keys[i]][0])
			if len(circuits[el2Idx]) == numCoords {
				return distances[keys[i]], nil
			}
		} else if el1Idx != -1 && el2Idx != -1 {
			circuits[el1Idx] = append(circuits[el1Idx], circuits[el2Idx]...)
			circuits[el2Idx] = []coord{}
			if len(circuits[el1Idx]) == numCoords {
				return distances[keys[i]], nil
			}
		} else {
			circuits = append(circuits, []coord{distances[keys[i]][0], distances[keys[i]][1]})
		}
	}

	return []coord{}, fmt.Errorf("[ERROR] Unable to find completing circuit\n")
}
