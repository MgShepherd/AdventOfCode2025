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
const NUM_CONNECTIONS_TO_MAKE = 1000
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

	circuits := getCircuits(getDistances(coords))

	return getProductLargestCircuits(circuits), nil
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

func getCircuits(distances map[int][]coord) [][]coord {
	keys := slices.Collect(maps.Keys(distances))
	slices.Sort(keys)

	circuits := [][]coord{}
	for i := range NUM_CONNECTIONS_TO_MAKE {
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
		} else if el1Idx == -1 && el2Idx != -1 {
			circuits[el2Idx] = append(circuits[el2Idx], distances[keys[i]][0])
		} else if el1Idx != -1 && el2Idx != -1 {
			circuits[el1Idx] = append(circuits[el1Idx], circuits[el2Idx]...)
			circuits[el2Idx] = []coord{}
		} else {
			circuits = append(circuits, []coord{distances[keys[i]][0], distances[keys[i]][1]})
		}
	}

	return circuits
}

func getProductLargestCircuits(circuits [][]coord) int {
	largestCircuits := make([]int, NUM_LARGEST_CIRCUITS)

	for _, c := range circuits {
		cLen := len(c)
		if cLen == 0 {
			continue
		}

		for i := range largestCircuits {
			if cLen > largestCircuits[i] {
				for j := len(largestCircuits) - 1; j > i; j-- {
					largestCircuits[j] = largestCircuits[j-1]
				}
				largestCircuits[i] = cLen
				break
			}
		}
	}

	product := 1
	for _, el := range largestCircuits {
		product *= el
	}
	return product
}
