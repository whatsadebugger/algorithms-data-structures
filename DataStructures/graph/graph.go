package graph

// Graph holds all the vertices.
type Graph struct {
	Vertices []*Vertex
	Count    int
}

// AddVertex adds vertex to graph based on label
func (g *Graph) AddVertex(label string) {
	g.Vertices = append(g.Vertices, NewVertex(label, g.Count))
	g.Count++
}

// FindVertexByLabel is a helper method to find vertex object by label
func (g *Graph) FindVertexByLabel(label string) *Vertex {
	for _, v := range g.Vertices {
		if v.Label == label {
			return v
		}
	}

	return nil
}

// AddEdge connects two vertices
func (g *Graph) AddEdge(from string, to string) {
	f := g.FindVertexByLabel(from)
	t := g.FindVertexByLabel(to)

	f.AddEdge(t)
}
