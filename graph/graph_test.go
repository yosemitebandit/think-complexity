package graph

import "testing"

func TestAddVertex(t *testing.T) {
	v1 := Vertex{"v1"}
	g := NewGraph()
	g.AddVertex(v1)
	if g.Vertices[0] != v1 {
		t.Errorf("expected v1, got ", g.Vertices[0])
	}
}

func TestAddEdge(t *testing.T) {
	v1, v2 := Vertex{"v1"}, Vertex{"v2"}
	e := Edge{v1, v2}
	g := NewGraph()
	g.AddEdge(e)
	if !g.VertexIsConnected(v1, v2) {
		t.Error("v1 -/-> v2")
	}
}

func TestVertexIsConnected(t *testing.T) {
	v1, v2, v3 := Vertex{"v1"}, Vertex{"v2"}, Vertex{"v3"}
	g := NewGraph()
	e := Edge{v1, v2}
	g.AddEdge(e)
	if !g.VertexIsConnected(v1, v2) {
		t.Error("expected true, got false")
	}
	if g.VertexIsConnected(v1, v3) {
		t.Error("expected false, got true")
	}
}

func TestGetEdge(t *testing.T) {
	v1, v2, v3 := Vertex{"v1"}, Vertex{"v2"}, Vertex{"v3"}
	g := NewGraph()
	e := Edge{v1, v2}
	g.AddEdge(e)
	i := g.GetEdgeIndex(v1, v2)
	expected_index := 0
	if i != expected_index {
		t.Errorf("expected %v, got %v", expected_index, i)
	}
	i = g.GetEdgeIndex(v1, v3)
	expected_index = -1
	if i != expected_index {
		t.Errorf("expected %v, got %v", expected_index, i)
	}
}

func TestRemoveEdge(t *testing.T) {
	v1, v2 := Vertex{"v1"}, Vertex{"v2"}
	g := NewGraph()
	e := Edge{v1, v2}
	g.AddEdge(e)
	g.RemoveEdge(e)
	if g.VertexIsConnected(v1, v2) {
		t.Error("expected false, got true")
	}
	expected_index := -1
	if i := g.GetEdgeIndex(v1, v2); i != expected_index {
		t.Errorf("expected %v, got %v", expected_index, i)
	}
}

func TestOutEdges(t *testing.T) {
	v1, v2 := Vertex{"v1"}, Vertex{"v2"}
	g := NewGraph()
	e := Edge{v1, v2}
	g.AddEdge(e)
	edges := g.OutEdges(v1)
	if 1 != len(edges) {
		t.Errorf("expected len of 1, got %v", len(edges))
	}
	if e != edges[0] {
		t.Errorf("expected %v, got %v", e, edges[0])
	}
}

func TestAddAllEdges(t *testing.T) {
	v1, v2, v3, v4, v5 := Vertex{"v1"}, Vertex{"v2"}, Vertex{"v3"}, Vertex{"v4"}, Vertex{"v5"}
	g := NewGraph()
	for _, v := range []Vertex{v1, v2, v3, v4, v5} {
		g.AddVertex(v)
	}
	g.AddAllEdges()
	if 10 != len(g.Edges) {
		t.Errorf("expected %v, got %v", 10, len(g.Edges))
	}
}

func TestIsConnected(t *testing.T) {
	v1, v2, v3, v4, v5 := Vertex{"v1"}, Vertex{"v2"}, Vertex{"v3"}, Vertex{"v4"}, Vertex{"v5"}
	// Connect all vertices to one another -- a thoroughly connected graph
	g1 := NewGraph()
	for _, v := range []Vertex{v1, v2, v3, v4, v5} {
		g1.AddVertex(v)
	}
	g1.AddAllEdges()
	if !g1.IsConnected() {
		t.Errorf("graph g1 should be connected.")
	}
	// Connect some vertices but not all.
	g2 := NewGraph()
	for _, v := range []Vertex{v1, v2, v3, v4, v5} {
		g1.AddVertex(v)
	}
	e1 := Edge{v1, v2}
	g2.AddEdge(e1)
	e2 := Edge{v1, v3}
	g2.AddEdge(e2)
	if g2.IsConnected() {
		t.Errorf("graph g2 should not be connected.")
	}
}
