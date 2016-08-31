package algo

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	TestEdge struct {
		Source, Destination string
		Weight              float64
	}
)

func (e TestEdge) GetSource() string {
	return e.Source
}

func (e TestEdge) GetDestination() string {
	return e.Destination
}

func (e TestEdge) GetWeight() float64 {
	return e.Weight
}

func (e TestEdge) String() string {
	return fmt.Sprint(e.Source, "->", e.Destination, " ", e.Weight)
}

func TestKruskalEmpty(t *testing.T) {
	edges := []Edge{}
	result := []Edge{}

	Kruskal(edges, &result)
	if len(result) != 0 {
		t.Error("Expecting empty response")
	}
}

func TestKruskalOne(t *testing.T) {
	edges := []Edge{
		TestEdge{Source: "a", Destination: "b"},
	}

	result := []Edge{}
	Kruskal(edges, &result)
	if len(result) != 1 {
		t.Fatal("Expecting one single edge, got ", len(result))
	}
	if !reflect.DeepEqual(edges[0], result[0]) {
		t.Error("Expecting the single response to be the same as the input")
	}
}

func TestKruskalTwo(t *testing.T) {
	edges := []Edge{
		TestEdge{Source: "a", Destination: "b"},
		TestEdge{Source: "b", Destination: "c"},
	}

	result := []Edge{}
	Kruskal(edges, &result)
	if len(result) != 2 {
		t.Fatal("Expecting two edges, got ", len(result))
	}
	for i := range result {
		if !reflect.DeepEqual(edges[i], result[i]) {
			t.Error("Edge mismatch")
		}
	}
}

func TestKruskalTwoLoop(t *testing.T) {
	edges := []Edge{
		TestEdge{Source: "a", Destination: "b", Weight: 10},
		TestEdge{Source: "b", Destination: "a", Weight: 1},
	}

	result := []Edge{}
	Kruskal(edges, &result)
	if len(result) != 1 {
		t.Fatal("Expecting one edge, got ", len(result))
	}

	e := result[0].(TestEdge)
	if !(e.Source == "b" && e.Destination == "a" && e.Weight == 1) {
		t.Error("Unexpected result: ", e)
	}
}

// From https://projecteuler.net/index.php?section=problems&id=107
func TestKruskalNetwork(t *testing.T) {
	edges := []Edge{
		TestEdge{Source: "a", Destination: "b", Weight: 16},
		TestEdge{Source: "b", Destination: "e", Weight: 20},
		TestEdge{Source: "e", Destination: "g", Weight: 11},
		TestEdge{Source: "g", Destination: "f", Weight: 27},
		TestEdge{Source: "f", Destination: "c", Weight: 31},
		TestEdge{Source: "c", Destination: "a", Weight: 12},
		TestEdge{Source: "a", Destination: "d", Weight: 21},
		TestEdge{Source: "b", Destination: "d", Weight: 17},
		TestEdge{Source: "e", Destination: "d", Weight: 18},
		TestEdge{Source: "g", Destination: "d", Weight: 23},
		TestEdge{Source: "f", Destination: "d", Weight: 19},
		TestEdge{Source: "c", Destination: "d", Weight: 28},
	}

	result := []Edge{}
	Kruskal(edges, &result)
	if len(result) != 6 {
		t.Fatal("Expecting 6 edges")
	}

	total := 0.0
	for _, item := range result {
		total += item.(TestEdge).Weight
	}
	if total != 93 {
		t.Error("Was expecting a total weight of 93, got ", total)
	}
}
