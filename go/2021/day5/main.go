package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type line struct {
	start, end point
}

func parseInput() []line {
	var res []line
	data, err := ioutil.ReadFile("../input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, l := range lines {
		strs := strings.Split(strings.ReplaceAll(l, " -> ", ","), ",")
		var nums []int
		for i := range strs {
			num, err := strconv.Atoi(strs[i])
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		if nums == nil || len(nums) != 4 {
			panic("nums must be 4 ints")
		}
		line := line{}
		line.start.x = nums[0]
		line.start.y = nums[1]
		line.end.x = nums[2]
		line.end.y = nums[3]
		res = append(res, line)
	}
	return res
}

func (l line) gridPointBetween(includeDiagonals bool) []point {
	var res []point
	if l.start.x == l.end.x {
		x := l.start.x
		var start, end int
		if l.start.y < l.end.y {
			start = l.start.y
			end = l.end.y
		} else {
			start = l.end.y
			end = l.start.y
		}
		for y := start; y <= end; y++ {
			res = append(res, point{x, y})
		}
	} else if l.start.y == l.end.y {
		y := l.start.y
		var start, end int
		if l.start.x < l.end.x {
			start = l.start.x
			end = l.end.x
		} else {
			start = l.end.x
			end = l.start.x
		}
		for x := start; x <= end; x++ {
			res = append(res, point{x, y})
		}
	} else if includeDiagonals {
		res = append(l.diagonalPoints())
	}
	return res
}

func (l line) diagonalPoints() []point {
	var res []point
	if l.start.x - l.end.x == l.start.y - l.end.y {
		// x = y graph /
		var start, end point
		if l.start.x < l.end.x {
			start = l.start
			end = l.end
		} else {
			start = l.end
			end = l.start
		}
		for x, y := start.x, start.y; x <= end.x; x, y = x+1, y+1 {
			res = append(res, point{x, y})
		}
	} else {
		// x = -y graph \
		var start, end point
		if l.start.x < l.end.x {
			start = l.start
			end = l.end
		} else {
			start = l.end
			end = l.start
		}
		for x, y := start.x, start.y; x <= end.x; x, y = x+1, y-1 {
			res = append(res, point{x, y})
		}
	}
	return res
}

func main() {
	lines := parseInput()
	coveredPoints := make(map[point]int)
	coveredPointsP2 := make(map[point]int)
	for _, line := range lines {
		pointsBetween := line.gridPointBetween(false)
		for _, point := range pointsBetween {
			coveredPoints[point]++
		}
		pointsBetween = line.gridPointBetween(true)
		for _, point := range pointsBetween {
			coveredPointsP2[point]++
		}
	}
	intersections := 0
	for _, intersection := range coveredPoints {
		if intersection > 1 {
			intersections++
		}
	}
	log.Print(intersections)

	intersections = 0
	for _, intersection := range coveredPointsP2 {
		if intersection > 1 {
			intersections++
		}
	}

	log.Print(intersections)
}
