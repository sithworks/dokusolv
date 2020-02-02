package solver

import (
	"errors"
	"fmt"
	"strings"
)

const DIMENSION = 9

type Board struct {
	Cells [][]int
}

type SudokuSolverFn func(*Board) *Board

func Make() *Board {
	var b Board
	b.Cells = make([][]int, DIMENSION)
	for i := range b.Cells {
		b.Cells[i] = make([]int, DIMENSION)
	}
	return &b
}

func (b *Board) String() string {
	var builder strings.Builder

	for x := 0; x < DIMENSION; x++ {
		for y := 0; y < DIMENSION; y++ {
			builder.WriteString(fmt.Sprintf("%2d|", b.Cells[x][y]))
		}
		builder.WriteString("\n")
	}
	return builder.String()

}

func(b *Board) Get(x, y int) (val int, err error) {
	if !(x >= 0 && x < DIMENSION && y >= 0 && y < DIMENSION) {
		return 0, errors.New("Index out of bounds")
	}
	return b.Cells[x][y], nil
}


func Solve(sudoku *Board) *Board {
	return sudoku
}
