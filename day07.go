package main

// template by LFJ

import (
	"flag"
	"fmt"
	"io/ioutil"

	// "io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day07.input", "Relative file path")
var partB = flag.Bool("partB", false, "Using part B logic")

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

type Node struct {
	value string
}

func (g *Graph) getAccessibleNodes(node *Node) map[Node]bool {
	visitedNodes := make(map[Node]bool)
	var nodesToVisit []*Node

	nodesToVisit = addElements(nodesToVisit, g.edges[*node])

	for len(nodesToVisit) > 0 {
		nodeToVisit := nodesToVisit[0]

		visitedNodes[*nodeToVisit] = true
		for _, node := range g.edges[*nodeToVisit] {
			if _, ok := visitedNodes[*node]; !ok {
				nodesToVisit = append(nodesToVisit, node)
			}
		}

		nodesToVisit = nodesToVisit[1:]
	}
	return visitedNodes
}

func (g *Graph) addNode(n *Node) {
	for _, node := range g.nodes {
		if node == n {
			return
		}
	}
	g.nodes = append(g.nodes, n)
}

func (g *Graph) addEdge(startNode, endNode *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*startNode] = append(g.edges[*startNode], endNode)
}

func (g *Graph) addNodesAndEdges(startNode *Node, endNodes []*Node) {
	g.addNode(startNode)
	for _, endNode := range endNodes {
		g.addNode(endNode)
		g.addEdge(endNode, startNode)
	}
}

func parseRule(rule string, existingNodes map[string]*Node, graph *Graph) {
	originNode, targetNodes := getNodes(rule, existingNodes)
	graph.addNodesAndEdges(originNode, targetNodes)
}

func getNodes(rule string, existingNodes map[string]*Node) (originNode *Node, targetNodes []*Node) {
	originColor := getOriginColor(rule)
	destinationColors := getDestinationColors(rule)

	originNode = getNode(originColor, existingNodes)
	targetNodes = make([]*Node, len(destinationColors))
	for i, color := range destinationColors {
		targetNodes[i] = getNode(color, existingNodes)
	}
	return originNode, targetNodes
}

func getNode(color string, existingNodes map[string]*Node) *Node {
	if _, ok := existingNodes[color]; !ok {
		existingNodes[color] = &Node{value: color}
	}
	return existingNodes[color]
}

func getOriginColor(rule string) (color string) {
	split := strings.Split(rule, " ")
	color = split[0] + " " + split[1]
	return color
}

func getDestinationColors(rule string) []string {
	split := strings.Split(rule, " ")
	nDestinationColor := 0
	for _, word := range split {
		if _, err := strconv.Atoi(word); err == nil {
			nDestinationColor++
		}
	}
	destinationColors := make([]string, nDestinationColor)
	padding := 4
	for i := 0; i < nDestinationColor; i++ {
		destinationColor := split[(i+1)*padding+1] + " " + split[(i+1)*padding+2]
		destinationColors[i] = destinationColor
	}
	return destinationColors
}

func addElements(nodesToVisit []*Node, nodes []*Node) []*Node {
	for _, node := range nodes {
		nodesToVisit = append(nodesToVisit, node)
	}
	return nodesToVisit
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	existingNodes := make(map[string]*Node)
	graphIsIn := &Graph{}
	for _, s := range split {
		parseRuleIsIn(s, existingNodes, graph)
	}
	fmt.Println(len(graph.getAccessibleNodes(getNode("shiny gold", existingNodes))))
}
