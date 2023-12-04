package main

import graph2 "playground/gocookbook/graph"

func main() {
	graph := graph2.NewGraph()
	c := &graph2.Node{Name: "C"}
	graph.AddNode(c)
	k := &graph2.Node{Name: "K"}
	graph.AddNode(k)
	g := &graph2.Node{Name: "G"}
	graph.AddNode(g)
	f := &graph2.Node{Name: "F"}
	graph.AddNode(f)
	graph.AddEdge(c, k, 1)
	graph.AddEdge(c, g, 2)
	graph.AddEdge(c, f, 5)
}
