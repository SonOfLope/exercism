package tree

import (
	"errors"
	"sort"
)

// Record represents the parent child relationship of records
type Record struct {
	ID     int
	Parent int
}

// Node represents the list of child records of each record
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree from the list of parent child relationship of records
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	recMap := make(map[int]*Node)

	for i, r := range records {

		// check for id continuity
		if r.ID != i {
			return nil, errors.New("Non continous record id")
		}

		if _, ok := recMap[r.ID]; ok {
			return nil, errors.New("Duplications are not allowed")
		}

		if r.ID != 0 && r.ID <= r.Parent {
			return nil, errors.New("Invalid ID number")
		}

		if r.ID == 0 && r.Parent != 0 {
			return nil, errors.New("Root node cannot have a parent")
		}

		n := &Node{ID: r.ID}
		recMap[r.ID] = n

		if _, ok := recMap[r.Parent]; ok && r.ID != 0 {
			recMap[r.Parent].Children = append(recMap[r.Parent].Children, n)
		}
	}

	return recMap[0], nil
}
