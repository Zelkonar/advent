package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type inputLine struct {
	signalPatterns, outputValue []string
}

type displayKey struct {
	top, topLeft, topRight, middle, bottom, bottomLeft, bottomRight string
}

// in retrospect... the display key keeping a map of (sorted) strings -> number would have been easier
// oh well
func numberFromWires(wires string, key displayKey) string {
	// sorting makes things easier
	blah := sortString(wires)
	switch (blah) {
	case sortString(key.top + key.topLeft + key.topRight + key.bottomLeft + key.bottomRight + key.bottom):
		return "0"
	case sortString(key.topRight + key.bottomRight):
		return "1"
	case sortString(key.top + key.topRight + key.middle + key.bottomLeft + key.bottom):
		return "2"
	case sortString(key.top + key.topRight + key.middle + key.bottomRight + key.bottom):
		return "3"
	case sortString(key.topLeft + key.topRight + key.middle + key.bottomRight):
		return "4"
	case sortString(key.top + key.topLeft + key.middle + key.bottomRight + key.bottom):
		return "5"
	case sortString(key.top + key.topLeft + key.middle + key.bottomLeft + key.bottomRight + key.bottom):
		return "6"
	case sortString(key.top + key.topRight + key.bottomRight):
		return "7"
	case sortString(key.top + key.topLeft + key.topRight + key.middle + key.bottomLeft + key.bottomRight + key.bottom):
		return "8"
	case sortString(key.top + key.topLeft + key.topRight + key.middle + key.bottomRight + key.bottom):
		return "9"
	default:
		panic("something has gone horribly horribly wrong")
	}
}

func sortString(v string) string {
	s := strings.Split(v, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func parseInput() []inputLine {
	data, err := ioutil.ReadFile("../input/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var input []inputLine
	for _, l := range lines {
		newInput := inputLine{}
		line := strings.Split(l, " | ")
		newInput.signalPatterns = strings.Split(line[0], " ")
		newInput.outputValue = strings.Split(line[1], " ")
		input = append(input, newInput)
	}
	return input
}

// this is obviously not the ideal way, but it works :)
func decryptDisplay(values []string) displayKey {
	key := displayKey{}

	known := make(map[int]string)
	for _, v := range values {
		if len(v) == 2 {
			known[1] = v
		}
		if len(v) == 3 {
			known[7] = v
		}
		if len(v) == 4 {
			known[4] = v
		}
		if len(v) == 7 {
			known[8] = v
		}
	}

	// top middle is easy, just the diff between 7 & 1
	for _, char := range strings.Split(known[7], "") {
		if strings.Contains(known[1], char) {
			continue
		}
		key.top = char
		break
	}

	// of the 3 values that have a len of 6, the number 6 is the only one
	// that doesn't have both '1' segments, ergo of the three numbers with
	// len 6, the one value missing that 1 includes is the top right
	for _, v := range values {
		if len(v) != 6 {
			continue
		}

		for _, char := range strings.Split(known[1], "") {
			if strings.Contains(v, char) {
				continue
			}
			key.topRight = char
			break
		}
	}

	// now bottom right is ezpz (the one from 1 that isn't topright
	for _, char := range strings.Split(known[1], "") {
		if char == key.topRight {
			continue
		}
		key.bottomRight = char
		break
	}

	// now we can do middle and bottom using "3", as it's the only len 5
	// display with topright, and bottomright. after determining the 2
	// candidates, we can use 4 to determine middle, then bottom, and topleft
	for _, v := range values {
		if len(v) != 5 {
			continue
		}
		if strings.Contains(v, key.topRight) && strings.Contains(v, key.bottomRight) {
			known[3] = v
			break
		}
	}

	// so middle is the one that both 3 and 4 has, that aren't topright or bottom right
	for _, char := range strings.Split(known[4], "") {
		if char == key.topRight || char == key.bottomRight || !strings.Contains(known[3], char) {
			continue
		}
		key.middle = char
		break
	}

	// now bottom is just the odd man out from 3
	for _, char := range strings.Split(known[3], "") {
		if char == key.topRight || char == key.top || char == key.middle || char == key.bottomRight {
			continue
		}
		key.bottom = char
		break
	}

	// we can go back to 4 for top left
	for _, char := range strings.Split(known[4], "") {
		if char == key.topRight || char == key.middle || char == key.bottomRight {
			continue
		}
		key.topLeft = char
		break
	}

	// finally, 8 for bottom left
	for _, char := range strings.Split(known[8], "") {
		if char == key.top || char == key.topRight || char == key.topLeft ||
			char == key.middle || char == key.bottomRight || char == key.bottom {
			continue
		}
		key.bottomLeft = char
		break
	}
	return key
}

func main() {
	inputs := parseInput()

	// part1
	counter := 0
	for _, input := range inputs {
		for _, outputValue := range input.outputValue {
			v := len(outputValue)
			if v == 2 || v == 3 || v == 4 || v == 7 {
				counter++
			}
		}
	}
	log.Print(counter)

	// part2
	total := 0
	for _, input := range inputs {
		key := decryptDisplay(input.signalPatterns)
		output := ""
		for _, outputValue := range input.outputValue {
			output += numberFromWires(outputValue, key)
		}
		num, err := strconv.Atoi(output)
		if err != nil {
			panic(err)
		}
		total += num
	}
	log.Print(total)
}
