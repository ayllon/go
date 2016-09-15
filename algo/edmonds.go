// Implemented from
// http://www.cs.tau.ac.il/~zwick/grad-algo-13/directed-mst.pdf


package algo

import (
	"container/heap"
	"github.com/satori/go.uuid"
	"math"
)

type (
	// dummyEdge is used to make the graph strongly connected
	// Its weight is +Inf, to make it very unappealing
	dummyEdge struct {
		Source, Destination string
	}

	// graph stores the state of the algorithm
	graph struct {
		Vertices map[string]bool
		Edges    []Edge

		// In is the chosen incoming edge for a given vertex
		In map[string]Edge
		// Const is used to adjust the weight when a supervertex is created
		Const map[string]float64
		// Prev is the vertex preceding a given vertex
		Prev map[string]*string
		// Parent stores the supervertex in which the vertex was contracted
		Parent map[string]*string
		// Children stores the list of vertices contracted into the given vertex
		Children map[string][]string
		// P is a priority queue with all the incoming edges of a given vertex
		P map[string]*edgeHeap
	}

	// edgeHeap implements the required methods by container/heap
	edgeHeap struct {
		edges []Edge
		g     *graph
	}
)

// Len returns the length of the heap
func (eh edgeHeap) Len() int {
	return len(eh.edges)
}

// Less returns true if the element i is less that the element j
func (eh edgeHeap) Less(i, j int) bool {
	return eh.g.weight(eh.edges[i]) < eh.g.weight(eh.edges[j])
}

// weight calculates the weight of the edge e as required by Edmond's algorithm
func (g *graph) weight(e Edge) float64 {
	w := e.GetWeight()
	v := e.GetDestination()
	for g.Parent[v] != nil {
		w += g.Const[e.GetDestination()]
		v = *g.Parent[v]
	}
	return w
}

// Swap elements i and j
func (eh edgeHeap) Swap(i, j int) {
	eh.edges[i], eh.edges[j] = eh.edges[j], eh.edges[i]
}

// Push a new element into the heap
func (eh *edgeHeap) Push(x interface{}) {
	eh.edges = append(eh.edges, x.(Edge))
}

// Pop the first element from the heap
func (eh *edgeHeap) Pop() interface{} {
	old := eh.edges
	n := len(old)
	x := old[n-1]
	eh.edges = old[0 : n-1]
	return x
}

// meld merges the src heap into the dst one, and clears the source
func meld(dst, src *edgeHeap) {
	dst.edges = append(dst.edges, src.edges...)
	heap.Init(dst)
	src.edges = nil
}

// GetSource returns the source of the edge
func (d dummyEdge) GetSource() string {
	return d.Source
}

// GetDestination returns the destination of the edge
func (d dummyEdge) GetDestination() string {
	return d.Destination
}

// GetWeight always return +Inf for a dummy edge
func (d dummyEdge) GetWeight() float64 {
	return math.Inf(0)
}

// Edmonds return the Directed Minimum Spanning Tree from edges in out
func Edmonds(root string, edges []Edge, out *[]Edge) {
	if len(edges) == 0 {
		return
	}

	g := &graph{
		Edges:    edges,
		Vertices: make(map[string]bool),

		In:       make(map[string]Edge),
		Const:    make(map[string]float64),
		Prev:     make(map[string]*string),
		Parent:   make(map[string]*string),
		Children: make(map[string][]string),
		P:        make(map[string]*edgeHeap),
	}
	for _, e := range edges {
		g.Vertices[e.GetSource()] = true
		g.Vertices[e.GetDestination()] = true
	}

	g.connectAll()
	g.initialize()
	g.contract()
	g.expand(root, out)
}

// A requirement of this implementation is that the graph is strongly connected,
// so we add extra edges with a very high weight where required
func (g *graph) connectAll() {
	type pair struct {
		u, v string
	}
	linked := make(map[pair]bool)

	// Build list of linked vertices
	for _, e := range g.Edges {
		linked[pair{u: e.GetSource(), v: e.GetDestination()}] = true
	}

	for u := range g.Vertices {
		for v := range g.Vertices {
			if u != v && !linked[pair{u: u, v: v}] {
				g.Edges = append(g.Edges, dummyEdge{Source: u, Destination: v})
			}
		}
	}
}

// contract the graph into a single supervertex
func (g *graph) contract() {
	a := g.pickVertex()
	for g.P[a].Len() > 0 {
		e := heap.Pop(g.P[a]).(Edge)
		b := g.find(e.GetSource())
		if a != b {
			g.In[a] = e
			g.Prev[a] = &b
			if g.In[e.GetSource()] == nil {
				a = b
			} else {
				c := uuid.NewV4().String()
				g.initVertex(c)
				for g.Parent[a] == nil {
					g.Parent[a] = &c
					g.Const[a] = -g.In[a].GetWeight()
					g.Children[c] = append(g.Children[c], a)
					meld(g.P[c], g.P[a])
					a = *g.Prev[a]
				}
				a = c
			}
		}
	}
}

// initialize the required data per vertex for the algorithm to work
func (g *graph) initialize() {
	for u := range g.Vertices {
		g.initVertex(u)
	}
	for _, e := range g.Edges {
		heap.Push(g.P[e.GetDestination()], e)
	}
}

// initVertex initialized the data for the given vertex
func (g *graph) initVertex(u string) {
	g.In[u] = nil
	g.Const[u] = 0
	g.Prev[u] = nil
	g.Parent[u] = nil
	g.Children[u] = nil
	g.P[u] = &edgeHeap{g: g}
	heap.Init(g.P[u])
}

// pickVertex just needs to return any vertex from the graph to start the contraction phase
func (g *graph) pickVertex() string {
	for v := range g.Vertices {
		return v
	}
	return ""
}

// find returns the isolated (super)vertex that contains u
func (g *graph) find(u string) string {
	for g.Parent[u] != nil {
		u = *g.Parent[u]
	}
	return u
}

// expand takes a contracted graph, and undo the contraction starting at the root
func (g *graph) expand(root string, out *[]Edge) {
	var R []string
	g.dismantle(&R, root)
	for len(R) > 0 {
		var c string
		c, R = R[0], R[1:]
		e := g.In[c]
		g.In[e.GetDestination()] = e
		g.dismantle(&R, e.GetDestination())
	}
	for u := range g.Vertices {
		if u == root || g.In[u] == nil {
			continue
		}
		*out = append(*out, g.In[u])
	}
}

// dismantle undos the contraction of the (super)vertex u, and puts the children into R
func (g *graph) dismantle(R *[]string, u string) {
	for g.Parent[u] != nil {
		for _, v := range g.Children[*g.Parent[u]] {
			if v != u {
				g.Parent[v] = nil
				if len(g.Children[v]) != 0 {
					*R = append(*R, v)
				}
			}
		}
		u = *g.Parent[u]
	}
}
