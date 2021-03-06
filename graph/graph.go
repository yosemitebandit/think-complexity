package graph

import (
	"math/rand"
	"time"
)

type Vertex struct {
	Label string
}

type Edge struct {
	Start Vertex
	End   Vertex
}

type Graph struct {
	Vertices    []Vertex
	Edges       []Edge
	connections map[Vertex][]Vertex
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Graph constructor.
func NewGraph() *Graph {
	var g Graph
	g.connections = make(map[Vertex][]Vertex)
	return &g
}

func (g *Graph) AddVertex(v Vertex) {
	g.Vertices = append(g.Vertices, v)
}

func (g *Graph) AddEdge(e Edge) {
	g.connections[e.Start] = append(g.connections[e.Start], e.End)
	g.connections[e.End] = append(g.connections[e.End], e.Start)
	g.Edges = append(g.Edges, e)
}

func vertexInSlice(v Vertex, slice []Vertex) bool {
	for _, vertex := range slice {
		if vertex == v {
			return true
		}
	}
	return false
}

// Determine if a vertex is connected to another vertex by checking connections.
func (g *Graph) VertexIsConnected(v1, v2 Vertex) bool {
	return vertexInSlice(v1, g.connections[v2]) || vertexInSlice(v2, g.connections[v1])
	/*
		for _, v := range g.connections[v1] {
			if v2 == v {
				return true
			}
		}
		for _, v := range g.connections[v2] {
			if v1 == v {
				return true
			}
		}
		return false
	*/
}

// Finds the edge that connects two vertices.
// Returns -1 if not found.
func (g *Graph) GetEdgeIndex(v1, v2 Vertex) int {
	for i, e := range g.Edges {
		if (e.Start == v1 && e.End == v2) || (e.Start == v2 && e.End == v1) {
			return i
		}
	}
	return -1
}

func (g *Graph) RemoveEdge(e Edge) {
	if i := g.GetEdgeIndex(e.Start, e.End); i != -1 {
		g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)
	}
	for i, v := range g.connections[e.Start] {
		if e.End == v {
			g.connections[e.Start] = append(g.connections[e.Start][:i], g.connections[e.Start][i+1:]...)
		}
	}
	for i, v := range g.connections[e.End] {
		if e.Start == v {
			g.connections[e.End] = append(g.connections[e.End][:i], g.connections[e.End][i+1:]...)
		}
	}
}

func (g *Graph) OutEdges(v Vertex) []Edge {
	edges := []Edge{}
	for _, e := range g.Edges {
		if e.Start == v || e.End == v {
			edges = append(edges, e)
		}
	}
	return edges
}

// From an edgeless graph, connect every pair of vertices.
func (g *Graph) AddAllEdges() {
	for i, v1 := range g.Vertices {
		for _, v2 := range g.Vertices[i+1:] {
			e := Edge{v1, v2}
			g.AddEdge(e)
		}
	}
}

// From an edgeless graph, add edges at random.
func (g *Graph) AddRandomEdges(p float32) {
	for i, v1 := range g.Vertices {
		for _, v2 := range g.Vertices[i+1:] {
			if p >= rand.Float32() {
				e := Edge{v1, v2}
				g.AddEdge(e)
			}
		}
	}
}

func (g *Graph) IsConnected() bool {
	//visited_vertices := []Vertex{}
	queue := []Vertex{}
	queue = append(queue, g.Vertices[0])
	for len(queue) > 0 {
		// Pop a vertex off the queue.
		v1 := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// Get all connected vertices.
		//connected := []Vertex{}
		for _, v2 := range g.Vertices {
			if g.VertexIsConnected(v1, v2) {
			}
		}

	}

	return false
}
