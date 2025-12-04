package problems

import "fmt"

func Solve(probNum int) (int, error) {
	switch probNum {
	case 1:
		return solveProblem1()
	case 2:
		return solveProblem2()
	case 3:
		return solveProblem3()
	case 4:
		return solveProblem4()
	default:
		return -1, fmt.Errorf("Unrecognised problem number %d\n", probNum)
	}
}
