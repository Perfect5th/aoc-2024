package algos

import (
	"math"
	"slices"
)

type Graph[T any] struct {
	vertices []*Vertex[T]
	edges map[Edge[T]]float64
}

func NewGraph[T any](vertices []*Vertex[T], edges map[Edge[T]]float64) *Graph[T] {
	return &Graph[T]{
		vertices: vertices,
		edges: edges,
	}
}

func (g *Graph[T]) Vertices() []*Vertex[T] {
	return g.vertices
}

func (g *Graph[T]) Edges(u *Vertex[T], v *Vertex[T]) (float64, bool) {
	edge, ok := g.edges[Edge[T]{u, v}]
	return edge, ok
}

type Vertex[T any] struct {
	neighbors []*Vertex[T]
	value T
}

func NewVertex[T any](value T) *Vertex[T] {
	return &Vertex[T]{
		neighbors: make([]*Vertex[T], 0),
		value: value,
	}
}

func (v *Vertex[T]) Neighbors() []*Vertex[T] {
	return v.neighbors
}

func (v *Vertex[T]) AddNeighbor(n *Vertex[T]) {
	v.neighbors = append(v.neighbors, n)
}

func (v *Vertex[T]) Value() T {
	return v.value
}

type Edge[T any] struct {U, V *Vertex[T]}

// An implementation of Dijkstra's algorithm for finding the shortest path between nodes.
// It returns the total distance of the shortest path from `source` to `target`.
func DijkstraCost[T any](graph *Graph[T], source *Vertex[T], target *Vertex[T]) (cost float64) {
	dist := make(map[*Vertex[T]]float64)
	q := make([]*Vertex[T], len(graph.Vertices()))
	prev := make(map[*Vertex[T]]*Vertex[T])

	for i, v := range graph.Vertices() {
		dist[v] = math.Inf(0)
		prev[v] = nil
		q[i] = v
	}

	dist[source] = 0

	var u *Vertex[T]
	for len(q) > 0 {
		u = q[0]
		remove := 0
		for i, v := range q[1:] {
			if dist[v] < dist[u] {
				u = v
				remove = i+1
			}
		}

		if u == target {
			break
		}
		q = slices.Delete(q, remove, remove+1)

		for _, v := range u.Neighbors() {
			if !slices.Contains(q, v) {
				continue
			}

			edge, ok := graph.Edges(u, v)
			if !ok {
				panic("No edge between vertices")
			}
			alt := dist[u] + edge

			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	return dist[target]
}
