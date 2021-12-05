package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Node struct {
	children []*Node
}

func buildTreeRec(treeStr string, badName string) []*Node {
	sepCount := strings.Count(treeStr, " ")
	if sepCount == 0 {
		return make([]*Node, 0)
	} else if sepCount == 1 {
		rawLeafs := strings.Split(treeStr, " ")
		leafs := make([]*Node, 0)
		for i := 0; i < len(rawLeafs); i++ {
			if rawLeafs[i] != badName {
				leafs = append(leafs, &Node{})
			}
		}
		return leafs
	} else {
		pOpen := -1
		treeSepIdx := -1
		for i := 0; i < len(treeStr); i++ {
			if treeStr[i] == '(' {
				if pOpen == -1 {
					pOpen = 1
				} else {
					pOpen++
				}
			} else if treeStr[i] == ')' {
				pOpen--
				if pOpen == 0 {
					treeSepIdx = i + 1
					break
				}
			}
		}
		leafs := make([]*Node, 0)
		leftPart := treeStr[:treeSepIdx]
		leftName := strings.Split(leftPart, "(")[0]
		if leftName == badName {
			leafs = append(leafs, buildTree(leftPart, badName).children...)
		} else {
			leafs = append(leafs, buildTree(leftPart, badName))
		}
		if len(treeStr) > treeSepIdx {
			rightPart := treeStr[treeSepIdx+1:]
			rightName := strings.Split(rightPart, "(")[0]
			if rightName == badName {
				leafs = append(leafs, buildTree(rightPart, badName).children...)
			} else {
				leafs = append(leafs, buildTree(rightPart, badName))
			}
		}
		return leafs
	}
}

func buildTree(treeStr string, badName string) *Node {
	pIdx := strings.Index(treeStr, "(")
	treeSubStr := treeStr
	if pIdx > -1 {
		pStart := pIdx + 1
		pEnd := len(treeStr) - 1
		treeSubStr = treeStr[pStart:pEnd]
	}
	return &Node{children: buildTreeRec(treeSubStr, badName)}
}

func calcTreeHeight(tree *Node) int {
	height := 0
	if len(tree.children) > 0 {
		for i := 0; i < len(tree.children); i++ {
			height = int(math.Max(float64(height), float64(calcTreeHeight(tree.children[i]))))
		}
		height++
	}
	return height
}

func main() {
	data, _ := ioutil.ReadFile("tree.txt")

	// Bygg tre
	tree := buildTree(string(data), "Grinch")

	// Finn høyde på tre
	treeHeight := calcTreeHeight(tree)

	fmt.Println(treeHeight)
}
