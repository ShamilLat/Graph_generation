package main

type EdgeId int
type VertexId int

type Vertex struct {
	Depth    int
	ID_      int
	EdgeIds_ []EdgeId
}

type Edge struct {
	ID_        int
	Vertex1Id_ VertexId
	Vertex2Id_ VertexId
	// Color_ Color
}

type Graph struct {
	VertNum_  VertexId
	EdgeNum_  EdgeId
	Vertices_ []Vertex
	Edges_    []Edge
	DepthMap_ [][]VertexId
	// ColoredEdges_ map[Edge.Color][]EdgeId
}

func (v Vertex) GetId() int {
	return v.ID_
}

func (v Vertex) GetEdgesIds() []EdgeId {
	return v.EdgeIds_
}

func (v Vertex) AddEdgeId(id EdgeId) {
	v.EdgeIds_ = append(v.EdgeIds_, id)
}

func (v Vertex) HasVertexId(id EdgeId) bool {

}
