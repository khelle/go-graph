package Graph

/**
 *
 */
type MatrixGraph struct {
	Vertices			[]*GraphVertex
	VerticesCount		int
	Edges				[][]*GraphEdge
	EdgesCount			int
	MaxVertexId			int
}

/**
 * @return *MatrixGraph
 */
func NewMatrixGraph() *MatrixGraph {
	graph := &MatrixGraph{}

	graph.Vertices 		= make([]*GraphVertex, 0)
	graph.VerticesCount = 0
	graph.Edges    		= make([][]*GraphEdge, 0)
	graph.EdgesCount    = 0
	graph.MaxVertexId   = -1

	return graph
}

/**
 *
 */
func (graph *MatrixGraph) AddVertex(id int, V *GraphVertex) {
	graph.VerticesCount++

	if graph.MaxVertexId < id {
		lastMax := graph.MaxVertexId
		currentMax := id

		vertices := make([]*GraphVertex, currentMax+1)
		copy(vertices, graph.Vertices)

		graph.Vertices = vertices

		var i int
		var edges [][]*GraphEdge

		edges = make([][]*GraphEdge, currentMax+1)
		for i = 0; i < currentMax+1; i++ {
			edges[i] = make([]*GraphEdge, currentMax+1)
		}
		for i = 0; i < lastMax+1; i++ {
			copy(edges[i], graph.Edges[i])
		}

		graph.Edges = edges
		graph.MaxVertexId = currentMax
	}

	if graph.Vertices[id] != nil {
		return
	}

	graph.Vertices[id] = V
}

/**
 *
 */
func (graph *MatrixGraph) GetVertex(id int) *GraphVertex {
	if id > graph.MaxVertexId {
		return nil
	}

	return graph.Vertices[id]
}

/**
 *
 */
func (graph *MatrixGraph) RemoveVertex(id int) {
	if graph.GetVertex(id) == nil {
		return;
	}

	var i int

	graph.Vertices[id] = nil
	for i = 0; i < graph.MaxVertexId+1; i++ {
		graph.Edges[id][i] = nil
	}
	for i = 0; i < graph.MaxVertexId+1; i++ {
		graph.Edges[i][id] = nil
	}

	if id != graph.MaxVertexId {
		return
	}
}

/**
 *
 */
func (graph *MatrixGraph) GetVertices() []*GraphVertex {
	var vertex *GraphVertex
	var vertices []*GraphVertex

	vertices = make([]*GraphVertex, 0)
	for i := 0; i < graph.MaxVertexId+1; i++ {
		vertex = graph.Vertices[i]

		if vertex != nil {
			vertices = append(vertices, vertex)
		}
	}

	return vertices
}

/**
 *
 */
func (graph *MatrixGraph) GetVerticesCount() int {
	return graph.VerticesCount
}

/**
 *
 */
func (graph *MatrixGraph) GetNeighbours(id int) []*GraphVertex {
	var edge *GraphEdge
	var edges []*GraphVertex

	edges = make([]*GraphVertex, 0)
	for i := 0; i < len(graph.Edges[id]); i++ {
		edge = graph.Edges[id][i]

		if edge != nil {
			edges = append(edges, graph.GetVertex(i))
		}
	}

	return edges
}

/**
 *
 */
func (graph *MatrixGraph) AddEdge(from int, to int, E *GraphEdge) {
	if graph.Edges[from][to] != nil {
		return
	}

	E.From = from
	E.To = to

	graph.Edges[from][to] = E
	graph.EdgesCount++
}

/**
 *
 */
func (graph *MatrixGraph) GetEdge(from int, to int) *GraphEdge {
	return graph.Edges[from][to]
}

/**
 *
 */
func (graph *MatrixGraph) RemoveEdge(from int, to int) {
	graph.Edges[from][to] = nil
	graph.EdgesCount--
}

/**
 *
 */
func (graph *MatrixGraph) GetEdges() []*GraphEdge {
	edges := make([]*GraphEdge, 0)

	for i := 0; i < graph.MaxVertexId+1; i++ {
		for j := 0; j < graph.MaxVertexId+1; j++ {
			if graph.Edges[i][j] != nil {
				edges = append(edges, graph.Edges[i][j])
			}
		}
	}

	return edges
}

/**
 *
 */
func (graph *MatrixGraph) GetEdgesCount() int {
	return graph.EdgesCount
}
