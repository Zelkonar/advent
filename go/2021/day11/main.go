package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type octopus struct {
	flashed bool
	v       int
}

type octopusGrid struct {
	octopi        []octopus
	width, height int
}

func (o octopusGrid) octopusAt(row, col int) *octopus {
	return &o.octopi[row*o.height+col]
}

func (o octopusGrid) String() string {
	s := ""
	for row := 0; row < o.height; row++ {
		for col := 0; col < o.width; col++ {
			s += strconv.Itoa(o.octopusAt(row, col).v)
		}
		s += "\n"
	}
	return s
}

func (o octopusGrid) flash(row, col int) {
	if row >= o.height || row < 0 || col >= o.width || col < 0 {
		return
	}
	thisOcto := o.octopusAt(row, col)
	if thisOcto.flashed {
		return
	}
	thisOcto.flashed = true
	// increase those around by 1, if > 9 then also try flash...
	for r := row-1; r <= row+1; r++ {
		for c := col-1; c <= col+1; c++ {
			if r >= o.height || r < 0 || c >= o.width || c < 0 {
				continue
			}
			if r == row && c == col {
				continue // this octopus
			}
			surroundingOctopus := o.octopusAt(r, c)
			surroundingOctopus.v++
			if surroundingOctopus.v > 9 {
				o.flash(r, c)
			}
		}
	}
}

func (o octopusGrid) step() int {
	for i := range o.octopi {
		o.octopi[i].v++ // lol
	}

	for row := 0; row < o.height; row++ {
		for col := 0; col < o.width; col++ {
			thisOcto := o.octopusAt(row, col)
			if thisOcto.v > 9 && !thisOcto.flashed {
				o.flash(row, col)
			}
		}
	}

	flashes := 0
	for i := range o.octopi {
		if o.octopi[i].v > 9 {
			o.octopi[i].v = 0
			flashes++
		}
		o.octopi[i].flashed = false
	}
	return flashes
}

func parseInput() octopusGrid {
	var res octopusGrid
	data, err := ioutil.ReadFile("../input/day11.txt")
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
		res.octopi = append(res.octopi, octopus{v: v})
	}

	return res
}

func main() {
	octoGrid := parseInput()
	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += octoGrid.step()
	}
	log.Print(flashes)

	octoGrid = parseInput()
	synchronizationStep := 0
	flashes = 0
	for flashes != 100{
		flashes = octoGrid.step()
		synchronizationStep++
	}
	log.Print(synchronizationStep)
}
