package main

import (
	"dokusolv/pkg/solver"
	"fmt"
)

func main() {
	fmt.Println("Dokusolv")
	arr := solver.Solve([][]int{{1, 2}, {3, 4}})
	fmt.Println(arr)

	b := solver.Make()

	fmt.Print(b)

}
