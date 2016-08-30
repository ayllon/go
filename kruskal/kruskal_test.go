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
)

func (e *TestEdge) GetSource() string {
	return e.Source
}

func (e *TestEdge) GetDestination() string {
	return e.Destination
}

func (e *TestEdge) Less(e2 Edge) bool {
	return e.Weight < e2.(*TestEdge).Weight
}

func (e *TestEdge) String() string {
	return fmt.Sprint(e.Source, "->", e.Destination, " ", e.Weight)
}

func TestEmpty(t *testing.T) {
	edges := make(Edges, 0)

	result := Kruskal(edges)
	if len(result) != 0 {
		t.Error("Expecting empty response")
	}
}

func TestOne(t *testing.T) {
	edges := make(Edges, 1)
	edges[0] = &TestEdge{Source: "a", Destination: "b"}

	result := Kruskal(edges)
	if len(result) != 1 {
		t.Fatal("Expecting one single edge")
	}
	if !reflect.DeepEqual(edges[0], result[0]) {
		t.Error("Expecting the single response to be the same as the input")
	}
}

func TestTwo(t *testing.T) {
	edges := make(Edges, 2)
	edges[0] = &TestEdge{Source: "a", Destination: "b"}
	edges[1] = &TestEdge{Source: "b", Destination: "c"}

	result := Kruskal(edges)
	if len(result) != 2 {
		t.Fatal("Expecting two edge")
	}
	for i := range result {
		if !reflect.DeepEqual(edges[i], result[i]) {
			t.Error("Edge mismatch")
		}
	}
}

func TestTwoLoop(t *testing.T) {
	edges := make(Edges, 2)
	edges[0] = &TestEdge{Source: "a", Destination: "b", Weight: 10}
	edges[1] = &TestEdge{Source: "b", Destination: "a", Weight: 1}

	result := Kruskal(edges)
	if len(result) != 1 {
		t.Fatal("Expecting one edge")
	}
	if !reflect.DeepEqual(result[0], edges[1]) {
		t.Error("Unexpected result: ", result[0])
	}
}

// From https://projecteuler.net/index.php?section=problems&id=107
func TestNetwork(t *testing.T) {
	edges := make(Edges, 12)
	edges[0] = &TestEdge{Source: "a", Destination: "b", Weight: 16}
	edges[1] = &TestEdge{Source: "b", Destination: "e", Weight: 20}
	edges[2] = &TestEdge{Source: "e", Destination: "g", Weight: 11}
	edges[3] = &TestEdge{Source: "g", Destination: "f", Weight: 27}
	edges[4] = &TestEdge{Source: "f", Destination: "c", Weight: 31}
	edges[5] = &TestEdge{Source: "c", Destination: "a", Weight: 12}
	edges[6] = &TestEdge{Source: "a", Destination: "d", Weight: 21}
	edges[7] = &TestEdge{Source: "b", Destination: "d", Weight: 17}
	edges[8] = &TestEdge{Source: "e", Destination: "d", Weight: 18}
	edges[9] = &TestEdge{Source: "g", Destination: "d", Weight: 23}
	edges[10] = &TestEdge{Source: "f", Destination: "d", Weight: 19}
	edges[11] = &TestEdge{Source: "c", Destination: "d", Weight: 28}

	result := Kruskal(edges)
	if len(result) != 6 {
		t.Fatal("Expecting 6 edges")
	}

	total := 0
	for i := range result {
		total += result[i].(*TestEdge).Weight
	}
	if total != 93 {
		t.Error("Was expecting a total weight of 93, got ", total)
	}
}
