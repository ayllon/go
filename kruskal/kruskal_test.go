package kruskal

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	TestEdge struct {
		Source, Destination string
		Weight              int
	}

	TestEdgeList struct {
		items []TestEdge
	}
)

func (e TestEdge) GetSource() string {
	return e.Source
}

func (e TestEdge) GetDestination() string {
	return e.Destination
}

func (e TestEdge) String() string {
	return fmt.Sprint(e.Source, "->", e.Destination, " ", e.Weight)
}

func (l *TestEdgeList) Len() int {
	return len(l.items)
}

func (l *TestEdgeList) Less(i, j int) bool {
	return l.items[i].Weight < l.items[j].Weight
}

func (l *TestEdgeList) Swap(i, j int) {
	l.items[i], l.items[j] = l.items[j], l.items[i]
}

func (l *TestEdgeList) Get(i int) Edge {
	return l.items[i]
}

func (l *TestEdgeList) Append(e Edge) {
	l.items = append(l.items, e.(TestEdge))
}

func TestEmpty(t *testing.T) {
	edges := &TestEdgeList{}
	result := &TestEdgeList{}

	Kruskal(edges, result)
	if result.Len() != 0 {
		t.Error("Expecting empty response")
	}
}

func TestOne(t *testing.T) {
	edges := &TestEdgeList{}
	result := &TestEdgeList{}

	edges.Append(TestEdge{Source: "a", Destination: "b"})

	Kruskal(edges, result)
	if result.Len() != 1 {
		t.Fatal("Expecting one single edge, got ", result.Len())
	}
	if !reflect.DeepEqual(edges.Get(0), result.Get(0)) {
		t.Error("Expecting the single response to be the same as the input")
	}
}

func TestTwo(t *testing.T) {
	edges := &TestEdgeList{}
	result := &TestEdgeList{}

	edges.Append(TestEdge{Source: "a", Destination: "b"})
	edges.Append(TestEdge{Source: "b", Destination: "c"})

	Kruskal(edges, result)
	if result.Len() != 2 {
		t.Fatal("Expecting two edges, got ", result.Len())
	}
	for i := range result.items {
		if !reflect.DeepEqual(edges.Get(i), result.Get(i)) {
			t.Error("Edge mismatch")
		}
	}
}

func TestTwoLoop(t *testing.T) {
	edges := &TestEdgeList{}
	result := &TestEdgeList{}

	edges.Append(TestEdge{Source: "a", Destination: "b", Weight: 10})
	edges.Append(TestEdge{Source: "b", Destination: "a", Weight: 1})

	Kruskal(edges, result)
	if result.Len() != 1 {
		t.Fatal("Expecting one edge, got ", result.Len())
	}

	e := result.items[0]
	if !(e.Source == "b" && e.Destination == "a" && e.Weight == 1) {
		t.Error("Unexpected result: ", e)
	}
}

// From https://projecteuler.net/index.php?section=problems&id=107
func TestNetwork(t *testing.T) {
	edges := &TestEdgeList{}
	edges.Append(TestEdge{Source: "a", Destination: "b", Weight: 16})
	edges.Append(TestEdge{Source: "b", Destination: "e", Weight: 20})
	edges.Append(TestEdge{Source: "e", Destination: "g", Weight: 11})
	edges.Append(TestEdge{Source: "g", Destination: "f", Weight: 27})
	edges.Append(TestEdge{Source: "f", Destination: "c", Weight: 31})
	edges.Append(TestEdge{Source: "c", Destination: "a", Weight: 12})
	edges.Append(TestEdge{Source: "a", Destination: "d", Weight: 21})
	edges.Append(TestEdge{Source: "b", Destination: "d", Weight: 17})
	edges.Append(TestEdge{Source: "e", Destination: "d", Weight: 18})
	edges.Append(TestEdge{Source: "g", Destination: "d", Weight: 23})
	edges.Append(TestEdge{Source: "f", Destination: "d", Weight: 19})
	edges.Append(TestEdge{Source: "c", Destination: "d", Weight: 28})

	result := &TestEdgeList{}
	Kruskal(edges, result)
	if result.Len() != 6 {
		t.Fatal("Expecting 6 edges")
	}

	total := 0
	for _, item := range result.items {
		total += item.Weight
	}
	if total != 93 {
		t.Error("Was expecting a total weight of 93, got ", total)
	}
}
