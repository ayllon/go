package algo

import (
	"sort"
)

type (
	forest struct {
		setLabel map[string]string
	}

	edgeSlice []Edge
)

func (eh edgeSlice) Len() int           { return len(eh) }
func (eh edgeSlice) Less(i, j int) bool { return eh[i].GetWeight() < eh[j].GetWeight() }
func (eh edgeSlice) Swap(i, j int)      { eh[i], eh[j] = eh[j], eh[i] }

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
func Kruskal(edges []Edge, out *[]Edge) {
	f := forest{
		setLabel: make(map[string]string),
	}

	for _, e := range edges {
		f.addVertex(e.GetSource())
		f.addVertex(e.GetDestination())
	}

	sort.Sort(edgeSlice(edges))

	for _, e := range edges {
		if f.union(e.GetSource(), e.GetDestination()) {
			*out = append(*out, e)
		}
	}
}
