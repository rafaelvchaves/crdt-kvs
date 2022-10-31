package delta

import (
	"kvs/crdt"
	pb "kvs/proto"
	"kvs/util"
)

type Graph struct{}

func NewGraph(r util.Replica) *Graph {
	return &Graph{}
}

func (g *Graph) AddVertex(v crdt.Vertex)         {}
func (g *Graph) RemoveVertex(v crdt.Vertex)      {}
func (g *Graph) AddEdge(e crdt.Edge)             {}
func (g *Graph) RemoveEdge(e crdt.Edge)          {}
func (g *Graph) LookupEdge(e crdt.Edge) bool     { return false }
func (g *Graph) LookupVertex(v crdt.Vertex) bool { return false }
func (g *Graph) GetEvent() *pb.Event             { return &pb.Event{} }
func (g *Graph) PersistEvent(event *pb.Event)    {}
