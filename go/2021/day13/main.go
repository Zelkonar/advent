package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	dirUp   = "up"
	dirLeft = "left"
)

type point struct {
	row, column int
}

type fold struct {
	val int
	dir string
}

type paper struct {
	points        map[point]bool
	width, height int // 11 x 15 for example
}

func (p *paper) fold(f fold) *paper {
	res := &paper{points: make(map[point]bool)}
	for paperPoint, v := range p.points {
		if !v { // shouldn't happen...
			continue
		}
		if f.dir == dirUp && paperPoint.row > f.val {
			distanceFromFold := paperPoint.row - f.val
			newPoint := point{
				row:    f.val - distanceFromFold,
				column: paperPoint.column,
			}
			res.points[newPoint] = true
			continue
		}
		if f.dir == dirLeft && paperPoint.column > f.val {
			distanceFromFold := paperPoint.column - f.val
			newPoint := point{
				row:    paperPoint.row,
				column: f.val - distanceFromFold,
			}
			res.points[newPoint] = true
			continue
		}
		res.points[paperPoint] = true
	}

	if f.dir == dirUp {
		res.height = f.val - 1
		res.width = p.width
	} else { // dirLeft
		res.width = f.val - 1
		res.height = p.height
	}
	return res
}

func (p *paper) String() string {
	res := ""
	for row := 0; row <= p.height; row++ {
		for col := 0; col <= p.width; col++ {
			if p.points[point{row, col}] {
				res += "#"
				continue
			}
			res += "."
		}
		res += "\n"
	}
	return res
}

func parseInput() (*paper, []fold) {
	data, err := ioutil.ReadFile("../input/day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	resPaper := &paper{points: make(map[point]bool)}
	var folds []fold
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if line == "" {
			continue
		}
		if strings.Contains(line, "fold along") {
			line = strings.ReplaceAll(line, "fold along ", "")
			strs := strings.Split(line, "=")
			f := fold{}
			if strs[0] == "y" {
				f.dir = dirUp
			} else {
				f.dir = dirLeft
			}
			f.val, err = strconv.Atoi(strs[1])
			if err != nil {
				panic(err)
			}

			folds = append(folds, f)
			continue
		}
		nums := make([]int, 2)
		strs := strings.Split(line, ",")
		nums[0], err = strconv.Atoi(strs[0])
		if err != nil {
			panic(err)
		}
		nums[1], err = strconv.Atoi(strs[1])
		if err != nil {
			panic(err)
		}
		if nums[0] > resPaper.width {
			resPaper.width = nums[0]
		}
		if nums[1] > resPaper.height {
			resPaper.height = nums[1]
		}
		resPaper.points[point{nums[1], nums[0]}] = true
	}
	return resPaper, folds
}

func main() {
	paper, folds := parseInput()

	p1Paper := paper.fold(folds[0])
	log.Println(len(p1Paper.points))

	for _, f := range folds {
		paper = paper.fold(f)
	}
	fmt.Print(paper)
}
