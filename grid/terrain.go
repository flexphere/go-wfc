package grid

import (
	"math"
)

type Terrain struct {
	Size  int
	Rules [][]int
	Grid  *Grid
}

func NewTerrain(size int, rules [][]int) *Terrain {
	t := &Terrain{
		Size:  3,
		Rules: rules,
	}
	buildStartingIndex := t.Size * t.Size / 2
	t.Grid = New(t.Size, t.Size, t.Rules)
	t.Grid.Build(buildStartingIndex)

	loopCount := int(math.Ceil(float64(size)-float64(t.Size)) / 2)
	for i := 0; i < loopCount; i++ {
		t.Expand()
	}
	return t
}

func (terrain *Terrain) Expand() {
	grid := New(terrain.Size+2, terrain.Size+2, terrain.Rules)
	grid.Merge(terrain.Grid, 1, 1)
	grid.Propagate(1)
	grid.Build(1)
	terrain.Size += 2
	terrain.Grid = grid
}

func (terrain *Terrain) Window(fromRow int, fromCol int, rowSize int, colSize int) *Grid {
	diffRow := fromRow + rowSize - terrain.Grid.Rows
	diffCol := fromCol + colSize - terrain.Grid.Cols

	if diffRow > 0 || diffCol > 0 {
		max := math.Max(float64(diffRow), float64(diffCol))
		for i := 0; i < int(max); i++ {
			terrain.Expand()
		}
	}

	return terrain.Grid.Range(fromRow, fromCol, rowSize, colSize)
}
