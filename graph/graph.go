package graph

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

// Determine if a vertex is connected to another vertex by checking connections.
func (g *Graph) IsConnected(v1, v2 Vertex) bool {
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
