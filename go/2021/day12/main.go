package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	//"regexp"
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
	data, err := ioutil.ReadFile("../input/dumb.txt")
	if err != nil {
		log.Fatal(err)
	}

	var res *caveNode
	allCaveNodes := make(map[string]*caveNode)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		// vals[1] needs vals[0] as a neighbor too!
		vals := strings.Split(line, "-")
		if _, ok := allCaveNodes[vals[1]]; !ok {
			allCaveNodes[vals[1]] = newCaveNode(vals[1], nil)
		}
		if node, ok := allCaveNodes[vals[0]]; ok {
			node.neighbors[vals[1]] = allCaveNodes[vals[1]]
		} else {
			allCaveNodes[vals[0]] = newCaveNode(vals[0], map[string]*caveNode{vals[1]: allCaveNodes[vals[1]]})
		}
		if vals[0] == "start" {
			res = allCaveNodes[vals[0]]
		}
	}
	return res
}

func main() {
	nodes := parseInput()
	fmt.Println(nodes)
}
