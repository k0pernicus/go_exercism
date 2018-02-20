package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Records []Record

func (r Records) Len() int { return len(r) }

func (r Records) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func (r Records) Less(i, j int) bool { return r[i].ID < r[j].ID }

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	l := len(records)
	if l == 0 {
		return nil, nil
	}
	// Sort all the records
	sort.Sort(Records(records))
	nodes := make([]Node, l)
	exists := make(map[int]bool)
	for i, r := range records {
		if exists[r.ID] {
			return nil, errors.New("Detected duplicate nodes")
		}
		if r.ID >= l {
			return nil, errors.New("The ID canno't be upper than the length of the nodes list")
		}
		if i == 0 && (r.ID != 0 || r.Parent != 0) {
			return nil, errors.New("Root has not been found")
		}
		if i != 0 && r.ID <= r.Parent {
			return nil, errors.New("The ID must be < than the Parents node")
		}
		nodes[i] = Node{ID: r.ID}
		exists[r.ID] = true
		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, &nodes[i])
		}
	}
	return &nodes[0], nil
}
