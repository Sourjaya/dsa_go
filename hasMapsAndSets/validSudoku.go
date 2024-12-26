package hashMapsAndSets

import (
	"github.com/Sourjaya/dsa/dStructures"
	dstructures "github.com/Sourjaya/dsa/dStructures"
)

func ValidSudoku(board [][]int) bool {
	rowSets := make([]dstructures.HashSet, 9)
	colSets := make([]dstructures.HashSet, 9)
	subgridSets := make([][]dstructures.HashSet, 3)

	for r := 0; r < 9; r++ {
		rowSets[r] = dStructures.NewHashSet()
	}
	for c := 0; c < 9; c++ {
		colSets[c] = dStructures.NewHashSet()
	}

	for r := 0; r < 3; r++ {
		subgridSets[r] = make([]dStructures.HashSet, 3)
		for c := 0; c < 3; c++ {
			subgridSets[r][c] = dStructures.NewHashSet()
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == 0 {
				continue
			}
			_, inRow := rowSets[i][num]
			_, inCol := colSets[j][num]
			_, inSubgrid := subgridSets[i/3][j/3][num]
			if inRow || inCol || inSubgrid {
				return false
			}
			rowSets[i][num] = dstructures.Void{}
			colSets[j][num] = dstructures.Void{}
			subgridSets[i/3][j/3][num] = dstructures.Void{}
		}
	}
	return true
}
