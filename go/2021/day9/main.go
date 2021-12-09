package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type grid struct {
	values        []int
	width, height int
}

func (g *grid) rowColIndex(column, row int) int { return g.width*row + column }

func (g *grid) valAt(column, row int) int {
	if row < 0 || column < 0 || row >= g.height || column >= g.width {
		return 9
	}
	return g.values[g.rowColIndex(column, row)]
}

func (g *grid) isLocalMinimum(column, row int) bool {
	v := g.valAt(column, row)
	return v < g.valAt(column+1, row) &&
		v < g.valAt(column, row+1) &&
		v < g.valAt(column-1, row) &&
		v < g.valAt(column, row-1)
}

func (g *grid) basinSizes() []int {
	visited := make(map[int]bool)
	var basinSizes []int
	for row := 0; row < g.height; row++ {
		for column := 0; column < g.width; column++ {
			if visited[g.rowColIndex(column, row)] || g.valAt(column, row) == 9 {
				continue
			}
			basinSizes = append(basinSizes, g.traverseBasinSize(column, row, visited))
		}
	}
	return basinSizes
}

// traverse each spot around, if 9, stop
func (g *grid) traverseBasinSize(column, row int, visited map[int]bool) int {
	if visited[g.rowColIndex(column, row)] || g.valAt(column, row) == 9 {
		return 0
	}
	visited[g.rowColIndex(column, row)] = true
	basinSize := 1 // self
	basinSize += g.traverseBasinSize(column+1, row, visited)
	basinSize += g.traverseBasinSize(column, row+1, visited)
	basinSize += g.traverseBasinSize(column-1, row, visited)
	basinSize += g.traverseBasinSize(column, row-1, visited)
	return basinSize
}

func parseInput() *grid {
	res := &grid{}
	data, err := ioutil.ReadFile("../input/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := string(data)
	res.width = len(strings.Split(s, "\n")[0])
	res.height = strings.Count(s, "\n")
	for _, val := range strings.Split(strings.ReplaceAll(s, "\n", ""), "") {
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		res.values = append(res.values, v)
	}

	return res
}

func main() {
	heightMap := parseInput()

	// Part 1
	riskLevelSum := 0
	for y := 0; y < heightMap.height; y++ {
		for x := 0; x < heightMap.width; x++ {
			if heightMap.isLocalMinimum(x, y) {
				riskLevelSum += heightMap.valAt(x, y) + 1
			}
		}
	}
	log.Print(riskLevelSum)

	// Part 2
	basins := heightMap.basinSizes()
	sort.Ints(basins)
	total := 1
	for i := len(basins) - 3; i < len(basins); i++ {
		total *= basins[i]
	}
	log.Print(total)
}
