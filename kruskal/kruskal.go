package kruskal

import (
	"sort"
)

type (
	// Edge is an interface that must be implemented by the types passed to Kruskal
	Edge interface {
		GetSource() string
		GetDestination() string
	}

	// Edges is a container for Edge types that must be sortable
	Edges interface {
		sort.Interface
		Get(i int) Edge
		Append(e Edge)
	}

	forest struct {
		setLabel map[string]string
	}
)

func (f *forest) addVertex(id string) {
	f.setLabel[id] = id
}

func (f *forest) union(a, b string) bool {
	setForA := f.setLabel[a]
	setForB := f.setLabel[b]
	if setForA == setForB {
		return false
	}

	f.setLabel[b] = setForA
	for child := range f.setLabel {
		if f.setLabel[child] == setForB {
			f.setLabel[child] = setForA
		}
	}

	return true
}

// Kruskal returns the min spanning tree (or forest, if disconnected) of the input graph
// modeled as its list of edges.
// Beware: the algorithm runs in place, so edges will be sorted.
func Kruskal(edges Edges, out Edges) {
	f := forest{
		setLabel: make(map[string]string),
	}

	for i := 0; i < edges.Len(); i++ {
		e := edges.Get(i)
		f.addVertex(e.GetSource())
		f.addVertex(e.GetDestination())
	}

	sort.Sort(edges)

	for i := 0; i < edges.Len(); i++ {
		e := edges.Get(i)
		if f.union(e.GetSource(), e.GetDestination()) {
			out.Append(e)
		}
	}
}
