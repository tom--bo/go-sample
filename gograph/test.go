package main

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
)

func main() {
	graphAst, _ := gographviz.Parse([]byte(`digraph G {}`))
	graph := gographviz.NewGraph()
	gographviz.Analyse(graphAst, graph)
	graph.AddEdge("\"a\"", "\"b\"", true, map[string]string{"label": "test"})
	graph.AddEdge("\"a\"", "\"c\"", true, map[string]string{"label": "test2"})
	graph.AddNode("G", "\"a\"", nil)
	graph.AddNode("G", "\"b\"", nil)
	output := graph.String()
	fmt.Println(output)

	var q []string
	q = append(q, "a")
	q = append(q, "b")
	q = append(q, "c")
}
