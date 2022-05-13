package grid

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

type Grid struct {
	Cells          []*Cell
	MaxDomains     int
	InitialDomains []int
	DomainRules    [][]int
	Rows           int
	Cols           int
	TotalCells     int
}

func New(rows int, cols int, domain_rules [][]int) *Grid {
	grid := new(Grid)
	grid.Rows = rows
	grid.Cols = cols
	grid.TotalCells = rows * cols

	grid.DomainRules = domain_rules
	grid.MaxDomains = len(domain_rules)

	for i := 0; i < grid.MaxDomains; i++ {
		grid.InitialDomains = append(grid.InitialDomains, i)
	}

	for i := 0; i < grid.TotalCells; i++ {
		grid.Cells = append(grid.Cells, NewCell(i, grid.InitialDomains))
	}

	return grid
}

func (grid *Grid) Build(index int) {
	if grid.allCollapsed() {
		return
	}

	c := grid.Cells[index]
	c.Collapse()

	neighbor_ids := []int{}
	neighbor_ids_uncollapsed := []int{}
	for _, neighbor := range grid.getNeighbors(c.Id) {
		neighbor_ids = append(neighbor_ids, neighbor.Id)
		if !neighbor.IsCollapsed {
			neighbor_ids_uncollapsed = append(neighbor_ids_uncollapsed, neighbor.Id)
		}
	}

	grid.propagateNeighbors(neighbor_ids)
	if len(neighbor_ids_uncollapsed) > 0 {
		randomNeighbor := neighbor_ids_uncollapsed[rnd.Intn(len(neighbor_ids_uncollapsed))]
		grid.Build(randomNeighbor)
	}
}

func (grid *Grid) Propagate(index int) []int {
	c := grid.Cells[index]
	if c.IsCollapsed {
		return nil
	}

	neighbor_rules := [][]int{}
	neighbor_domains := [][]int{}
	unpropagated_neighbors := []int{}
	for _, neighbor := range grid.getNeighbors(index) {
		neighbor_rules = append(neighbor_rules, grid.getRulesFromDomain(neighbor.Domain))
		neighbor_domains = append(neighbor_domains, neighbor.Domain)

		if !neighbor.IsPropagated {
			unpropagated_neighbors = append(unpropagated_neighbors, neighbor.Id)
		}
	}

	c.IsPropagated = true
	c.Domain = FindAllDuplicates(neighbor_rules)

	return unpropagated_neighbors
}

func (grid *Grid) Merge(subgrid *Grid, gridRow int, gridCol int) {

	for i := range subgrid.Cells {
		if i >= subgrid.TotalCells {
			return
		}

		cell := subgrid.Cells[i]
		subRow := int(i / subgrid.Cols)
		subCol := i % subgrid.Cols

		targetRow := gridRow + subRow
		targetCol := gridCol + subCol
		targetIndex := targetRow*grid.Cols + targetCol

		if targetIndex >= grid.TotalCells {
			return
		}

		grid.Cells[targetIndex] = cell
		grid.Cells[targetIndex].Id = targetIndex
	}
}

func (grid *Grid) Range(fromRow int, fromCol int, rowSize int, colSize int) *Grid {
	if fromRow+rowSize > grid.Rows || fromCol+colSize > grid.Cols {
		panic("Range out of bounds")
	}

	g := New(rowSize, colSize, grid.DomainRules)
	for i := 0; i < rowSize; i++ {
		startIndex := fromRow*grid.Cols + fromCol
		endIndex := startIndex + colSize
		loop := endIndex - startIndex
		for j := 0; j < loop; j++ {
			g.Cells[i*g.Cols+j] = grid.Cells[i*grid.Cols+startIndex+j]
		}
	}

	return g
}

func (grid *Grid) AsArray() []int {
	var result []int
	for _, cell := range grid.Cells {
		result = append(result, cell.Colapsed_domain)
	}
	return result
}

func (grid *Grid) PrintDebug(index int) {
	for i, c := range grid.Cells {
		if i%grid.Cols == 0 {
			fmt.Printf("\n")
		}

		if c.Id == index {
			fmt.Printf("\x1b[31m%d\x1b[0m", c.Colapsed_domain)
			continue
		}

		if c.IsCollapsed {
			fmt.Printf("\x1b[32m%d\x1b[0m", c.Colapsed_domain)
			continue
		}

		if c.IsPropagated {
			fmt.Printf("\x1b[33m%d\x1b[0m ", c.Colapsed_domain)
			continue
		}

		fmt.Printf("%d", c.Colapsed_domain)
	}
	fmt.Printf("\n")
}

func (grid *Grid) PrintDebugWithDomain(index int) {
	for i, c := range grid.Cells {
		if i%grid.Cols == 0 {
			fmt.Printf("\n")
		}

		var domains []string
		for _, domain := range c.Domain {
			domains = append(domains, fmt.Sprintf("%d", domain))
		}

		pad := grid.MaxDomains - len(domains)
		for i := 0; i < pad; i++ {
			domains = append(domains, " ")
		}

		domain := strings.Join(domains, ",")

		if c.Id == index {
			fmt.Printf(" \x1b[31m%d | %v\x1b[0m ", c.Colapsed_domain, domain)
			continue
		}

		if c.IsCollapsed {
			fmt.Printf(" \x1b[32m%d | %v\x1b[0m ", c.Colapsed_domain, domain)
			continue
		}

		if c.IsPropagated {
			fmt.Printf(" \x1b[33m%d | %v\x1b[0m ", c.Colapsed_domain, domain)
			continue
		}

		fmt.Printf(" %d | %v ", c.Colapsed_domain, domain)
	}
	fmt.Printf("\n")
}

func (grid *Grid) Print() {
	for i, c := range grid.Cells {
		if i%grid.Cols == 0 {
			fmt.Printf("\n")
		}

		if c.Colapsed_domain == 0 {
			fmt.Printf("\x1b[48;2;%d;%d;%dm%d\x1b[0m", 112, 181, 255, c.Colapsed_domain)
			continue
		}

		if c.Colapsed_domain == 1 {
			fmt.Printf("\x1b[48;2;%d;%d;%dm%d\x1b[0m", 255, 205, 112, c.Colapsed_domain)
			continue
		}

		if c.Colapsed_domain == 2 {
			fmt.Printf("\x1b[48;2;%d;%d;%dm%d\x1b[0m", 93, 184, 79, c.Colapsed_domain)
			continue
		}

		if c.Colapsed_domain == 3 {
			fmt.Printf("\x1b[48;2;%d;%d;%dm%d\x1b[0m", 57, 105, 49, c.Colapsed_domain)
			continue
		}

		if c.Colapsed_domain == 4 {
			fmt.Printf("\x1b[48;2;%d;%d;%dm%d\x1b[0m", 84, 59, 19, c.Colapsed_domain)
			continue
		}

		if c.Colapsed_domain == 5 {
			fmt.Printf("\x1b[48;2;%d;%d;%dm%d\x1b[0m", 43, 41, 36, c.Colapsed_domain)
			continue
		}

		fmt.Printf("\x1b[48;2;%d;%d;%dm%d\x1b[0m", 255, 255, 255, c.Colapsed_domain)
	}
	fmt.Printf("\n")
}

func (grid *Grid) propagateNeighbors(indexes []int) {
	nextPropagation := []int{}
	for _, index := range indexes {
		neighbors := grid.Propagate(index)
		nextPropagation = append(nextPropagation, neighbors...)
	}

	if len(nextPropagation) > 0 {
		grid.propagateNeighbors(nextPropagation)
	}
}

func (grid *Grid) getNeighbors(index int) []*Cell {
	neighbors := make([]*Cell, 0)

	topIndex := index - grid.Cols
	bottomIndex := index + grid.Cols
	rightIndex := index + 1
	leftIndex := index - 1

	if topIndex >= 0 {
		neighbors = append(neighbors, grid.Cells[topIndex])
	}

	if bottomIndex < grid.TotalCells {
		neighbors = append(neighbors, grid.Cells[bottomIndex])
	}

	if rightIndex < grid.TotalCells && rightIndex%grid.Cols != 0 {
		neighbors = append(neighbors, grid.Cells[rightIndex])
	}

	if leftIndex >= 0 && leftIndex%grid.Cols < grid.Cols-1 {
		neighbors = append(neighbors, grid.Cells[leftIndex])
	}

	return neighbors
}

func (grid *Grid) getRulesFromDomain(domains []int) []int {
	var rules []int
	for _, domain := range domains {
		rules = append(rules, grid.DomainRules[domain]...)
	}

	return FindUnique(rules)
}

func (grid *Grid) allCollapsed() bool {
	for _, cell := range grid.Cells {
		if cell.IsCollapsed == false {
			return false
		}
	}
	return true
}
