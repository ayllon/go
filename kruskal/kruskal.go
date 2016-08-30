package kruskal

import (
	"sort"
)

type (
	// Edge is an interface that must be implemented by the types passed to Kruskal
	Edge interface {
		GetSource() string
		GetDestination() string
		Less(Edge) bool
	}

	// Edges is an alias for a slice of Edge
	Edges []Edge

	forest struct {
		setLabel map[string]string
	}
)

func (es Edges) Len() int {
	return len(es)
}

func (es Edges) Less(i, j int) bool {
	return es[i].Less(es[j])
}

func (es Edges) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

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
func Kruskal(edges Edges) Edges {
	f := forest{
		setLabel: make(map[string]string),
	}

	for _, edge := range edges {
		f.addVertex(edge.GetSource())
		f.addVertex(edge.GetDestination())
	}

	edgesCopy := make(Edges, len(edges))
	copy(edgesCopy, edges)
	sort.Sort(edgesCopy)

	minSpanForest := make(Edges, 0)
	for _, edge := range edgesCopy {
		if f.union(edge.GetSource(), edge.GetDestination()) {
			minSpanForest = append(minSpanForest, edge)
		}
	}

	return minSpanForest
}
