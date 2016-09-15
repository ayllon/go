package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ayllon/go/algo"
	"github.com/tmc/dot"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	source, destination string
	weight              float64
}

func (e edge) GetSource() string {
	return e.source
}

func (e edge) GetDestination() string {
	return e.destination
}

func (e edge) GetWeight() float64 {
	return e.weight
}

func populateGraph(paths ...string) []algo.Edge {
	var edges []algo.Edge

	for _, path := range paths {
		log.Println("Processing", path)
		edge := edge{}

		fd, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()

		reader := bufio.NewReader(fd)
		for line, err := reader.ReadString('\n'); err != io.EOF; line, err = reader.ReadString('\n') {
			if err != nil {
				log.Fatal(err)
			}

			components := strings.Split(line, " ")
			if len(components) < 3 {
				continue
			}
			edge.source, edge.destination = components[0], components[1]
			edge.weight, err = strconv.ParseFloat(strings.Trim(components[2], " \n"), 64)
			if err != nil {
				log.Fatal(err)
			}

			edges = append(edges, edge)
			log.Println("Edge", edge)
		}
	}

	return edges
}

func sumWeight(edges []algo.Edge) float64 {
	sum := 0.0
	for _, e := range edges {
		sum += e.GetWeight()
	}
	return sum
}

func printDot(edges []algo.Edge) {
	g := dot.NewGraph("span")
	nodes := make(map[string]*dot.Node)

	for _, e := range edges {
		src, ok := nodes[e.GetSource()]
		if !ok {
			src = dot.NewNode(e.GetSource())
			nodes[e.GetSource()] = src
			g.AddNode(src)
		}
		dst, ok := nodes[e.GetDestination()]
		if !ok {
			dst = dot.NewNode(e.GetDestination())
			nodes[e.GetDestination()] = dst
			g.AddNode(dst)
		}
		de := dot.NewEdge(src, dst)
		de.Set("label", fmt.Sprint(e.GetWeight()))
		g.AddEdge(de)
	}

	os.Stdout.WriteString(g.String())
}

func main() {
	log.SetOutput(os.Stderr)

	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("Need at least the root name and one file")
	}

	root := flag.Arg(0)

	in := populateGraph(flag.Args()[1:]...)
	inWeight := sumWeight(in)
	var out []algo.Edge

	log.Println("Calling Edmond's algorithm with root", root, "and", len(in), "edges, total weight", inWeight)
	algo.Edmonds(root, in, &out)

	outWeight := sumWeight(out)
	log.Println("Got", len(out), "edges out, total weight", outWeight)

	printDot(out)
}
