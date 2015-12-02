# go-graph

Graph implementation written in Go

```
type Graph interface {
	AddVertex(id int, V *GraphVertex)
	GetVertex(id int)						*GraphVertex
	RemoveVertex(id int)
	GetVertices()							[]*GraphVertex
	GetVerticesCount()						int
	GetNeighbours(id int)					[]*GraphVertex
	AddEdge(from int, to int, E *GraphEdge)
	GetEdge(from int, to int)				*GraphEdge
	RemoveEdge(from int, to int)
	GetEdges()								[]*GraphEdge
	GetEdgesCount()							int
}
```