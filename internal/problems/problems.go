package problems

import "fmt"

func Solve(probNum int) (int, error) {
	switch probNum {
	case 1:
		return solveProblem1()
	default:
		return -1, fmt.Errorf("Unrecognised problem number %d\n", probNum)
	}
}
