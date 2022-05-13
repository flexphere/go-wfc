package main

import (
	"fmt"
	"math"
	"wfc/grid"
	"wfc/timer"
)

func main() {
	timer.Start("Main")

	terrainSize := 3.0
	rules := [][]int{
		{0, 1},
		{0, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{4, 5},
	}

	t := grid.NewTerrain(int(terrainSize), rules)

	targetSize := 256.0
	loopCount := int(math.Ceil(targetSize-terrainSize) / 2)

	for i := 0; i < loopCount; i++ {
		fmt.Printf("%d/%d\r", i, loopCount)
		t.Expand()
	}
	timer.End("Main")

	fmt.Printf("rows:%d cols:%d total:%d\n", t.Grid.Rows, t.Grid.Cols, t.Grid.TotalCells)
}
