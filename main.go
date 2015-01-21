package main

import (
	"fmt"

	"github.com/yosemitebandit/think-complexity/graph"
)

func main() {
	v1, v2, v3, v4, v5 := graph.Vertex{"v1"}, graph.Vertex{"v2"}, graph.Vertex{"v3"}, graph.Vertex{"v4"}, graph.Vertex{"v5"}
	e1, e2 := graph.Edge{v1, v2}, graph.Edge{v3, v4}

	g := graph.NewGraph()

	for _, v := range []graph.Vertex{v1, v2, v3, v4, v5} {
		g.AddVertex(v)
	}
	for _, e := range []graph.Edge{e1, e2} {
		g.AddEdge(e)
	}

	fmt.Println(g)
}
