package functions

import (
	"encoding/json"
)

func MatrixToString(matrix [][]int) (string, error) {
	data, err := json.Marshal(matrix)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// StringToMatrix converts a JSON string to a [][]int
func StringToMatrix(data string) ([][]int, error) {
	var matrix [][]int
	err := json.Unmarshal([]byte(data), &matrix)
	if err != nil {
		return nil, err
	}
	return matrix, nil
}
