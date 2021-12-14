package main

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
)

type polymer struct {
	start, end         string
	pairCounts         map[string]int
	pairInsertionRules map[string]string
}

func (p *polymer) pairInsertionStep() {
	newPairs := make(map[string]int)
	for pair, insertionChar := range p.pairInsertionRules {
		splitPair := strings.Split(pair, "")
		newPairs[splitPair[0]+insertionChar] += p.pairCounts[pair]
		newPairs[insertionChar+splitPair[1]] += p.pairCounts[pair]
	}
	p.pairCounts = newPairs
}

func (p *polymer) countOfEachElement() map[string]int {
	res := make(map[string]int)
	for pair, count := range p.pairCounts {
		splitPair := strings.Split(pair, "")
		res[splitPair[0]] += count
		res[splitPair[1]] += count
	}
	for char, count := range res {
		res[char] = count / 2
	}
	res[p.end]++
	res[p.start]++
	return res
}

func parseInput() *polymer {
	data, err := ioutil.ReadFile("../input/day14.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := &polymer{
		pairCounts:         make(map[string]int),
		pairInsertionRules: make(map[string]string),
	}
	for i, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if i == 0 {
			chars := strings.Split(line, "")
			for j := 0; j < len(chars)-1; j++ {
				if j == 0 {
					res.start = chars[j]
				}
				if j == len(chars)-2 {
					res.end = chars[j+1]
				}
				res.pairCounts[chars[j]+chars[j+1]]++
			}
			continue
		}
		if line == "" {
			continue
		}

		insertionPair := strings.Split(line, " -> ")
		res.pairInsertionRules[insertionPair[0]] = insertionPair[1]
	}
	return res
}

func main() {
	poly := parseInput()
	for i := 0; i < 10; i++ {
		poly.pairInsertionStep()
	}
	largest, smallest := 0, math.MaxInt64
	for _, count := range poly.countOfEachElement() {
		if count > largest {
			largest = count
		}
		if count < smallest && count != 0 {
			smallest = count
		}
	}
	log.Print(largest - smallest)

	poly = parseInput()
	for i := 0; i < 40; i++ {
		poly.pairInsertionStep()
	}
	largest, smallest = 0, math.MaxInt64
	for _, count := range poly.countOfEachElement() {
		if count > largest {
			largest = count
		}
		if count < smallest && count != 0 {
			smallest = count
		}
	}
	log.Print(largest - smallest)
}
