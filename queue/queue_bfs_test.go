package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemGraph_Bfs(t *testing.T) {
	type TestCase struct {
		target      *Node
		expecteFind bool
		expecteStep int
	}

	graph := NewItemGraph()

	a := &Node{Value: "A"}
	b := &Node{Value: "B"}
	c := &Node{Value: "C"}
	d := &Node{Value: "D"}
	e := &Node{Value: "E"}
	f := &Node{Value: "F"}
	g := &Node{Value: "G"}

	graph.AddNode(a)
	graph.AddNode(b)
	graph.AddNode(c)
	graph.AddNode(d)
	graph.AddNode(e)
	graph.AddNode(f)
	graph.AddNode(g)

	graph.AddEdge(a, b)
	graph.AddEdge(a, c)
	graph.AddEdge(a, d)
	graph.AddEdge(b, e)
	graph.AddEdge(c, e)
	graph.AddEdge(c, f)
	graph.AddEdge(d, g)
	graph.AddEdge(f, g)

	var tcs = []TestCase{
		{
			target:      g,
			expecteFind: true,
			expecteStep: 2,
		},
		{
			target:      b,
			expecteFind: true,
			expecteStep: 1,
		},
	}

	for _, tc := range tcs {
		find, step := graph.Bfs(tc.target)
		assert.Equal(t, tc.expecteFind, find)
		assert.Equal(t, tc.expecteStep, step)
	}

}
