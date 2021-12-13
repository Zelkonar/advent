package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type caveNode struct {
	neighbors map[string]*caveNode
	label     string
	isBig     bool
}

func newCaveNode(label string, neighbors map[string]*caveNode) *caveNode {
	isBig, err := regexp.MatchString("[A-Z]+", label)
	if err != nil {
		panic(err)
	}
	if neighbors == nil {
		neighbors = make(map[string]*caveNode)
	}
	return &caveNode{
		neighbors: neighbors,
		label:     label,
		isBig:     isBig,
	}
}

func parseInput() *caveNode {
	data, err := ioutil.ReadFile("../input/day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	var res *caveNode
	allCaveNodes := make(map[string]*caveNode)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		vals := strings.Split(line, "-")
		if _, ok := allCaveNodes[vals[1]]; !ok {
			allCaveNodes[vals[1]] = newCaveNode(vals[1], nil)
		}
		if node, ok := allCaveNodes[vals[0]]; ok {
			node.neighbors[vals[1]] = allCaveNodes[vals[1]]
		} else {
			allCaveNodes[vals[0]] = newCaveNode(vals[0], map[string]*caveNode{vals[1]: allCaveNodes[vals[1]]})
		}
		allCaveNodes[vals[1]].neighbors[vals[0]] = allCaveNodes[vals[0]]
		if vals[0] == "start" {
			res = allCaveNodes[vals[0]]
		}
	}
	return res
}

type path struct {
	v []string
}

func findAllPaths(node *caveNode, haveVisitedSmallCaveMax bool, visited map[string]bool) []*path {
	if visited == nil {
		visited = make(map[string]bool)
	}
	if visited[node.label] && !node.isBig {
		haveVisitedSmallCaveMax = true
	}
	visited[node.label] = true

	var paths []*path
	for k := range node.neighbors {
		if (!node.neighbors[k].isBig && visited[k] && haveVisitedSmallCaveMax) || k == "start"{
			continue // don't want these paths
		}

		if k == "end" {
			paths = append(paths, []*path{{v: []string{node.label, k}}}...)
			continue
		}

		// copy map for each path so they don't 'share' visits
		copyVisited := make(map[string]bool)
		for k, v := range visited {
			copyVisited[k] = v
		}

		// for every path my neighbor has, create a new path with myself as the root
		for _, p := range findAllPaths(node.neighbors[k], haveVisitedSmallCaveMax, copyVisited) {
			paths = append(paths, &path{v: append([]string{node.label}, p.v...)})
		}
	}
	return paths
}

func main() {
	startNode := parseInput()
	paths := findAllPaths(startNode, true, nil)
	log.Print(len(paths))

	paths = findAllPaths(startNode, false, nil)
	log.Print(len(paths))
}
