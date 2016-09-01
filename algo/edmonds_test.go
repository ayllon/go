package algo

import (
	"reflect"
	"testing"
)

func TestEdmondsEmpty(t *testing.T) {
	edges := []Edge{}
	result := []Edge{}

	Edmonds("", edges, &result)
	if len(result) != 0 {
		t.Error("Expecting empty response")
	}
}

func TestEdmondsOne(t *testing.T) {
	edges := []Edge{
		TestEdge{Source: "a", Destination: "b"},
	}

	result := []Edge{}
	Edmonds("a", edges, &result)
	if len(result) != 1 {
		t.Fatal("Expecting one single edge, got ", len(result))
	}
	if !reflect.DeepEqual(edges[0], result[0]) {
		t.Error("Expecting the single response to be the same as the input")
	}
}

func TestEdmondsTwo(t *testing.T) {
	edges := []Edge{
		TestEdge{Source: "a", Destination: "b", Weight: 1},
		TestEdge{Source: "b", Destination: "c", Weight: 10},
	}

	result := []Edge{}
	Edmonds("a", edges, &result)
	if len(result) != 2 {
		t.Fatal("Expecting two edges, got ", len(result))
	}
	visited := make(map[string]bool)
	for _, e := range result {
		visited[e.GetSource()] = true
		visited[e.GetDestination()] = true
	}
	t.Log(result)
	if !(visited["a"] && visited["b"] && visited["c"]) {
		t.Error("Not all vertices visited")
	}
}

func TestEdmondsTwoLoop(t *testing.T) {
	edges := []Edge{
		TestEdge{Source: "a", Destination: "b", Weight: 10},
		TestEdge{Source: "b", Destination: "a", Weight: 1},
	}

	result := []Edge{}
	Edmonds("a", edges, &result)
	if len(result) != 1 {
		t.Fatal("Expecting one edge, got ", len(result))
	}

	e := result[0].(TestEdge)
	if !reflect.DeepEqual(e, edges[0]) {
		t.Error("Unexpected result: ", e)
	}

	// Other way around
	result = []Edge{}
	Edmonds("b", edges, &result)
	if len(result) != 1 {
		t.Fatal("Expecting one edge, got ", len(result))
	}

	e = result[0].(TestEdge)
	if !reflect.DeepEqual(e, edges[1]) {
		t.Error("Unexpected result: ", e)
	}
}

// From https://projecteuler.net/index.php?section=problems&id=107
func TestEdmondsNetwork(t *testing.T) {
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
	Edmonds("a", edges, &result)
	t.Log(result)

	visited := make(map[string]bool)
	for _, e := range result {
		visited[e.GetSource()] = true
		visited[e.GetDestination()] = true
	}

	for _, e := range edges {
		if !visited[e.GetSource()] {
			t.Error("Not visited: ", e.GetSource())
		} else if !visited[e.GetDestination()] {
			t.Error("Not visited: ", e.GetDestination())
		}
	}
}
