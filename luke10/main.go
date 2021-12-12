package main

import "fmt"

type Node struct {
	name        string
	passingName string
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func solve(adjGraph map[string][]Node, start string, length int, visited []string) int {
	if len(visited) == length {
		return 1
	}
	numPatterns := 0
	ns := adjGraph[start]
	for i := 0; i < len(ns); i++ {
		node := ns[i]
		isVisited := contains(visited, node.name)
		failedToPass := len(node.passingName) > 0 && !contains(visited, node.passingName)
		if isVisited || failedToPass {
			continue
		}
		numPatterns += solve(adjGraph, node.name, length, append(visited, node.name))
	}
	return numPatterns
}

func main() {
	adjGraph := map[string][]Node{
		"A": {Node{"B", ""}, Node{"C", "B"}, Node{"D", ""}, Node{"E", ""}, Node{"F", ""}, Node{"G", "D"}, Node{"H", ""}, Node{"I", "E"}},
		"B": {Node{"A", ""}, Node{"C", ""}, Node{"D", ""}, Node{"E", ""}, Node{"F", ""}, Node{"G", ""}, Node{"H", "E"}, Node{"I", ""}},
		"C": {Node{"A", "B"}, Node{"B", ""}, Node{"D", ""}, Node{"E", ""}, Node{"F", ""}, Node{"G", "E"}, Node{"H", ""}, Node{"I", "F"}},
		"D": {Node{"A", ""}, Node{"B", ""}, Node{"C", ""}, Node{"E", ""}, Node{"F", "E"}, Node{"G", ""}, Node{"H", ""}, Node{"I", ""}},
		"E": {Node{"A", ""}, Node{"B", ""}, Node{"C", ""}, Node{"D", ""}, Node{"F", ""}, Node{"G", ""}, Node{"H", ""}, Node{"I", ""}},
		"F": {Node{"A", ""}, Node{"B", ""}, Node{"C", ""}, Node{"D", "E"}, Node{"E", ""}, Node{"G", ""}, Node{"H", ""}, Node{"I", ""}},
		"G": {Node{"A", "D"}, Node{"B", ""}, Node{"C", "E"}, Node{"D", ""}, Node{"E", ""}, Node{"F", ""}, Node{"H", ""}, Node{"I", "H"}},
		"H": {Node{"A", ""}, Node{"B", "E"}, Node{"C", ""}, Node{"D", ""}, Node{"E", ""}, Node{"F", ""}, Node{"G", ""}, Node{"I", ""}},
		"I": {Node{"A", "E"}, Node{"B", ""}, Node{"C", "F"}, Node{"D", ""}, Node{"E", ""}, Node{"F", ""}, Node{"G", "H"}, Node{"H", ""}},
	}
	numPatterns := solve(adjGraph, "D", 8, []string{"D"})
	fmt.Println(numPatterns)
}
