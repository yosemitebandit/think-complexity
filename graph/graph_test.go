package graph

import "testing"

func TestAddVertex(t *testing.T) {
	v1 := Vertex{"v1"}
	g := NewGraph()
	g.AddVertex(v1)
	if g.vertices[0] != v1 {
		t.Error("expected v1, got ", g.vertices[0])
	}
}

func TestAddEdge(t *testing.T) {
	v1, v2 := Vertex{"v1"}, Vertex{"v2"}
	e := Edge{v1, v2}
	g := NewGraph()
	g.AddEdge(e)
	if !g.IsConnected(v1, v2) {
		t.Error("v1 -/-> v2")
	}
}

func TestIsConnected(t *testing.T) {
	v1, v2, v3 := Vertex{"v1"}, Vertex{"v2"}, Vertex{"v3"}
	g := NewGraph()
	e := Edge{v1, v2}
	g.AddEdge(e)
	if !g.IsConnected(v1, v2) {
		t.Error("expected true, got false")
	}
	if g.IsConnected(v1, v3) {
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
		t.Error("expected %v, got %v", expected_index, i)
	}
	i = g.GetEdgeIndex(v1, v3)
	expected_index = -1
	if i != expected_index {
		t.Error("expected %v, got %v", expected_index, i)
	}

}
