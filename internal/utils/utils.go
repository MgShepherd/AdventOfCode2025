package utils

import (
	"fmt"
	"os"
)

func ReadProblemFile(probNum int) (string, error) {
	fileName := fmt.Sprintf("data/problem%d.txt", probNum)
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("[ERROR]: Unable to read file %s\n", fileName)
	}
	return string(data), err
}
