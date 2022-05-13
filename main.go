package main

import (
	"fmt"
	"wfc/grid"
	"wfc/timer"
)

func main() {
	timer.Start("Main")

	rules := [][]int{
		{0, 1},
		{0, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{4, 5},
	}

	t := grid.NewTerrain(256, rules)
	// g := t.Window(0, 8, 3, 3)
	// t.Grid.Print()
	// g.Print()
	timer.End("Main")
	fmt.Printf("rows:%d cols:%d total:%d\n", t.Grid.Rows, t.Grid.Cols, t.Grid.TotalCells)
}
