package main

import (
	"dokusolv/pkg/solver"
	"fmt"
)

func main() {
	fmt.Println("Dokusolv")
	b := solver.Make()
	fmt.Print(b)

	arr := solver.Solve(b)
	fmt.Println(arr)
}
