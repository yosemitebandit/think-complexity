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
	g := NewGraph()
	g.AddEdge(v1, v2)
	if !g.IsConnected(v1, v2) {
		t.Error("v1 -/-> v2")
	}
}

func TestIsConnected(t *testing.T) {
	v1, v2, v3 := Vertex{"v1"}, Vertex{"v2"}, Vertex{"v3"}
	g := NewGraph()
	g.AddEdge(v1, v2)
	if !g.IsConnected(v1, v2) {
		t.Error("expected true, got false")
	}
	if g.IsConnected(v1, v3) {
		t.Error("expected false, got true")
	}
}

func TestGetEdgeWhenEdgeExists(t *testing.T) {
	v1, v2 := Vertex{"v1"}, Vertex{"v2"}
	g := NewGraph()
	g.AddEdge(v1, v2)
}
