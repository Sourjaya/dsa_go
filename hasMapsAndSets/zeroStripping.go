package hashMapsAndSets

import (
	"fmt"

	"github.com/Sourjaya/dsa/dStructures"
)

func ZeroStrippingUsingSets(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	zeroRows, zeroCols := dStructures.NewHashSet(), dStructures.NewHashSet()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = dStructures.Void{}
				zeroCols[j] = dStructures.Void{}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			_, inRow := zeroRows[i]
			_, inCol := zeroCols[j]
			if inRow || inCol {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println("Matrix after setting zeroes : ", matrix)
}

func ZeroStrippingInPlace(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	firstRowHasZero := false
	firstColHasZero := false

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowHasZero = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstRowHasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColHasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	fmt.Println("Matrix after setting zeroes : ", matrix)
}
