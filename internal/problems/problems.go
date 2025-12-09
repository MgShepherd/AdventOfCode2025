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
	case 5:
		return solveProblem5()
	case 6:
		return solveProblem6()
	case 7:
		return solveProblem7()
	case 8:
		return solveProblem8()
	case 9:
		return solveProblem9()
	default:
		return -1, fmt.Errorf("Unrecognised problem number %d\n", probNum)
	}
}
