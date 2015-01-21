package graph

type Vertex struct {
	Label string
}

type Edge struct {
	Start Vertex
	End   Vertex
}

type Graph struct {
	vertices    []Vertex
	edges       []Edge
	connections map[Vertex][]Vertex
}

// Graph constructor.
func NewGraph() *Graph {
	var g Graph
	g.connections = make(map[Vertex][]Vertex)
	return &g
}

func (g *Graph) AddVertex(v Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) AddEdge(v1, v2 Vertex) {
	g.connections[v1] = append(g.connections[v1], v2)
	g.connections[v2] = append(g.connections[v2], v1)
	e := Edge{v1, v2}
	g.edges = append(g.edges, e)
}

// Determine if a vertex is connected to another vertex by checking connections.
func (g *Graph) IsConnected(v1, v2 Vertex) bool {
	for _, v := range g.connections[v1] {
		if v2 == v {
			return true
		}
	}
	return false
}
