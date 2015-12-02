package Graph

/**
 *
 */
type LinkedGraphVertex struct {
	Vertex				*GraphVertex
	Edges				[]*LinkedGraphEdge
}

/**
 *
 */
func NewLinkedGraphVertex(vertex *GraphVertex) *LinkedGraphVertex {
	linked := &LinkedGraphVertex{}

	linked.Vertex = vertex
	linked.Edges = make([]*LinkedGraphEdge, 0)

	return linked
}

/**
 *
 */
type LinkedGraphEdge struct {
	Vertex				*GraphVertex
	Edge				*GraphEdge
}

/**
 *
 */
func NewLinkedGraphEdge(vertex *GraphVertex, edge *GraphEdge) *LinkedGraphEdge {
	linked := &LinkedGraphEdge{}

	linked.Vertex = vertex
	linked.Edge = edge

	return linked
}

/**
 *
 */
type LinkedGraph struct {
	Vertices			[]*LinkedGraphVertex
	VerticesCount		int
	EdgesCount			int
	MaxVertexId			int
}

/**
 * @return *LinkedGraph
 */
func NewLinkedGraph() *LinkedGraph {
	graph := &LinkedGraph{}

	graph.Vertices 		= make([]*LinkedGraphVertex, 0)
	graph.VerticesCount = 0
	graph.EdgesCount    = 0
	graph.MaxVertexId   = -1

	return graph
}

/**
 *
 */
func (graph *LinkedGraph) AddVertex(id int, V *GraphVertex) {
	graph.VerticesCount++

	if graph.MaxVertexId < id {
		currentMax := id

		vertices := make([]*LinkedGraphVertex, currentMax+1)
		copy(vertices, graph.Vertices)

		graph.Vertices = vertices
		graph.MaxVertexId = currentMax
	}

	if graph.Vertices[id] != nil {
		return
	}

	graph.Vertices[id] = NewLinkedGraphVertex(V)
}

/**
 *
 */
func (graph *LinkedGraph) GetVertex(id int) *GraphVertex {
	if id > graph.MaxVertexId {
		return nil
	}

	vertex := graph.Vertices[id]
	if vertex == nil {
		return nil
	}

	return vertex.Vertex
}

/**
 *
 */
func (graph *LinkedGraph) RemoveVertex(id int) {
	if graph.GetVertex(id) == nil {
		return;
	}

	graph.Vertices[id] = nil
}

/**
 *
 */
func (graph *LinkedGraph) GetVertices() []*GraphVertex {
	var vertex *LinkedGraphVertex
	var vertices []*GraphVertex

	vertices = make([]*GraphVertex, 0)
	for i := 0; i < graph.MaxVertexId+1; i++ {
		vertex = graph.Vertices[i]

		if vertex != nil && vertex.Vertex != nil {
			vertices = append(vertices, vertex.Vertex)
		}
	}

	return vertices
}

/**
 *
 */
func (graph *LinkedGraph) GetVerticesCount() int {
	return graph.VerticesCount
}

/**
 *
 */
func (graph *LinkedGraph) GetNeighbours(id int) []*GraphVertex {
	if graph.MaxVertexId < id {
		return nil
	}

	vertices := make([]*GraphVertex, 0)
	for _, edge := range graph.Vertices[id].Edges {
		vertices = append(vertices, edge.Vertex)
	}

	return vertices
}

/**
 *
 */
func (graph *LinkedGraph) AddEdge(from int, to int, E *GraphEdge) {
	if graph.GetEdge(from, to) != nil {
		return
	}

	parent := graph.Vertices[from]
	E.From = from
	E.To = to
	parent.Edges = append(parent.Edges, NewLinkedGraphEdge(graph.GetVertex(to), E))
	graph.EdgesCount++
}

/**
 *
 */
func (graph *LinkedGraph) GetEdge(from int, to int) *GraphEdge {
	parent := graph.Vertices[from]

	for _, edge := range parent.Edges {
		if edge.Vertex.Id == to {
			return edge.Edge
		}
	}

	return nil
}

/**
 *
 */
func (graph *LinkedGraph) RemoveEdge(from int, to int) {
	parent := graph.Vertices[from]

	for key, edge := range parent.Edges {
		if edge.Vertex.Id == to {
			parent.Edges = append(parent.Edges[:key], parent.Edges[key+1:]...)
			graph.EdgesCount--
			return
		}
	}
}

/**
 *
 */
func (graph *LinkedGraph) GetEdges() []*GraphEdge {
	edges := make([]*GraphEdge, 0)

	for _, vertex := range graph.Vertices {
		if vertex != nil {
			for _, edge := range vertex.Edges {
				edges = append(edges, edge.Edge)
			}
		}
	}

	return edges
}

/**
 *
 */
func (graph *LinkedGraph) GetEdgesCount() int {
	return graph.EdgesCount
}
