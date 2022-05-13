package grid

import (
	"fmt"
)

// var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

type Cell struct {
	Id              int
	Domain          []int
	Colapsed_domain int
	IsPropagated    bool
	IsCollapsed     bool
}

func NewCell(id int, initialDomains []int) *Cell {
	cell := new(Cell)
	cell.Id = id
	cell.Domain = initialDomains
	cell.Colapsed_domain = 9
	cell.IsPropagated = false
	cell.IsCollapsed = false
	return cell
}

func (cell *Cell) Collapse() {
	if len(cell.Domain) < 1 {
		fmt.Printf("%v\n", cell)
	}
	cell.IsCollapsed = true
	cell.IsPropagated = true
	cell.Colapsed_domain = cell.Domain[rnd.Intn(len(cell.Domain))]
	cell.Domain = []int{cell.Colapsed_domain}
}

func (cell *Cell) String() string {
	return fmt.Sprintf("{\tid:%d,\tdomain:%v,\tcolapsed_domain:%d,\tpropagated:%v,\tcollappsed:%v}\n", cell.Id, cell.Domain, cell.Colapsed_domain, cell.IsPropagated, cell.IsCollapsed)
}
