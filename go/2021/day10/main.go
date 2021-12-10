package main

import (
	"errors"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var errNoCorruption = errors.New("no corruption on this line")

var openClose = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var corruptionScore = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var completionScore = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

type stack struct {
	data []string
}

// look at value at top of stack without pop
func (s *stack) peek() string {
	if s.data == nil || len(s.data) == 0 {
		return "" // normally error, but this isn't production code
	}
	return s.data[len(s.data)-1]
}
func (s *stack) push(v string) { s.data = append(s.data, v) }

func (s *stack) pop() string {
	v := s.peek()
	s.data = s.data[:len(s.data)-1]
	return v
}

func isOpenChar(s string) bool {
	return openClose[s] != ""
}

func parseInput() []string {
	data, err := ioutil.ReadFile("../input/day10.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func findCorruption(line string) (string, error) {
	closeStack := &stack{}
	for _, char := range strings.Split(line, "") {
		if isOpenChar(char) {
			closeStack.push(openClose[char])
			continue
		}
		if char == closeStack.pop() {
			continue
		}
		return char, nil
	}
	return "", errNoCorruption
}

// precondition: must have no corruption
// we could write a method to check, but I don't want to right now :)
func lineCompletionScore(line string) int {
	closeStack := &stack{}
	for _, char := range strings.Split(line, "") {
		if isOpenChar(char) {
			closeStack.push(openClose[char])
			continue
		}
		if char == closeStack.pop() {
			continue
		}
	}
	score := 0
	for len(closeStack.data) != 0 {
		v := closeStack.pop()
		score = score*5 + completionScore[v]
	}

	return score
}

func main() {
	lines := parseInput()
	var incompleteLines []string
	var corruptions []string
	for _, line := range lines {
		char, err := findCorruption(line)
		if err == errNoCorruption {
			incompleteLines = append(incompleteLines, line)
			continue
		}
		corruptions = append(corruptions, char)
	}
	sum := 0
	for _, corruption := range corruptions {
		sum += corruptionScore[corruption]
	}
	log.Print(sum)

	var scores []int
	for _, line := range incompleteLines {
		scores = append(scores, lineCompletionScore(line))
	}
	sort.Ints(scores)
	log.Print(scores[(len(scores)/2)])
}
