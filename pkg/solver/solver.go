package solver

import (
	"fmt"
	"strings"
)

const DIMENSION = 9

type board struct {
	Cells [][]int
}

func Make() *board {
	var b board
	b.Cells = make([][]int, DIMENSION)
	for i := range b.Cells {
		b.Cells[i] = make([]int, DIMENSION)
	}
	return &b
}

func (b *board) String() string {
	var builder strings.Builder

	for x := 0; x < DIMENSION; x++ {
		for y := 0; y < DIMENSION; y++ {
			builder.WriteString(fmt.Sprintf("%2d|", b.Cells[x][y]))
		}
		builder.WriteString("\n")
	}
	return builder.String()

}

func Solve(sudoku [][]int) [][]int {
	return sudoku
}
