package graph

type Node struct {
	Name string
}

type Edge struct {
	Node   *Node
	Weight int
}

type Graph struct {
	Nodes []*Node
	Edges map[string][]*Edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: nil,
		Edges: map[string][]*Edge{},
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node, w int) {
	g.Edges[n1.Name] = append(g.Edges[n1.Name], &Edge{
		Node:   n2,
		Weight: w,
	})

	g.Edges[n2.Name] = append(g.Edges[n2.Name], &Edge{
		Node:   n1,
		Weight: w,
	})
}

func (g *Graph) RemoveEdge(n1, n2 string) {
	removeEdge(g, n1, n2)
	removeEdge(g, n2, n1)
}

func (g *Graph) RemoveNode(name string) {
	idx := -1
	for i, n := range g.Nodes {
		if n.Name == name {
			idx = i
			break
		}
	}

	if idx > -1 {
		g.Nodes[idx] = g.Nodes[len(g.Nodes)-1]
		g.Nodes = g.Nodes[:len(g.Nodes)-1]
	}

	delete(g.Edges, name)

	for k := range g.Edges {
		removeEdge(g, k, name)
	}
}

func removeEdge(g *Graph, n, m string) {
	edges := g.Edges[n]

	idx := -1

	for i, e := range edges {
		if e.Node.Name == m {
			idx = i
			break
		}
	}

	if idx > -1 {
		edges[idx] = edges[len(edges)-1]
		g.Edges[n] = edges[:len(edges)-1]
	}
}
