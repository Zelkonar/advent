package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("../../../input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	depth, horiz := 0, 0
	p2Depth, p2Horiz, p2Aim := 0, 0, 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		dir := splitLine[0]
		amount, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case "up":
			depth -= amount
			p2Aim -= amount
		case "down":
			depth += amount
			p2Aim += amount
		case "forward":
			horiz += amount

			p2Horiz += amount
			p2Depth += p2Aim * amount
		}
	}
	fmt.Println(depth * horiz)
	fmt.Println(p2Depth * p2Horiz)
}
